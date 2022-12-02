package github

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	"github.com/cnartlu/area-service/api"
	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/component/compress/zip7"
	"github.com/cnartlu/area-service/component/filesystem"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/biz/transaction"
	"go.uber.org/zap"
)

type XiangyuecnRepository interface {
	GetLatestRelease(ctx context.Context) (*GithubRepositoryRelease, error)
}

type GithubUsecase struct {
	app         *app.App
	logger      *log.Logger
	filesystem  *filesystem.FileSystem
	transaction transaction.Transaction
	xiangyuecn  XiangyuecnRepository
	release     release.ManageRepo
	asset       asset.ManageRepo
}

func (g *GithubUsecase) GetLatestRelease(ctx context.Context) (*release.Release, error) {
	githubReponseRelease, err := g.xiangyuecn.GetLatestRelease(ctx)
	if err != nil {
		return nil, err
	}
	rep := githubReponseRelease.RepositoryRelease
	result, err := g.release.FindOne(ctx, release.ReleaseIDEQ(uint64(rep.GetID())))
	if err != nil {
		if !api.IsDataNotFound(err) {
			return nil, err
		}
		err = g.transaction.Transaction(ctx, func(ctx context.Context) error {
			if err != nil {
				if !api.IsDataNotFound(err) {
					return err
				}
				result, err = g.release.Save(ctx, &release.Release{
					Owner:              githubReponseRelease.Owner,
					Repo:               githubReponseRelease.Repo,
					ReleaseID:          uint64(rep.GetID()),
					ReleaseName:        rep.GetName(),
					ReleaseNodeID:      rep.GetNodeID(),
					ReleasePublishedAt: uint64(rep.GetPublishedAt().Unix()),
					ReleaseContent:     rep.GetBody(),
					Status:             release.StatusWaitSync,
				})
				if err != nil {
					return err
				}
				for _, repAsset := range rep.Assets {
					repAsset := repAsset
					downloadUrl := repAsset.GetBrowserDownloadURL()
					uri, err := url.Parse(downloadUrl)
					if err != nil {
						return err
					}
					var filepath string = filepath.Join(g.app.GetShortRuntimePath(), result.Owner, result.ReleaseName, filepath.Base(uri.Path))
					filepath = filepath[1:]
					_, err = g.asset.Save(ctx, &asset.Asset{
						AreaReleaseID: result.ID,
						AssetName:     repAsset.GetName(),
						AssetID:       uint64(repAsset.GetID()),
						AssetLabel:    repAsset.GetLabel(),
						AssetState:    repAsset.GetState(),
						FileSize:      uint(repAsset.GetSize()),
						FilePath:      filepath,
						DownloadURL:   downloadUrl,
						Status:        asset.StatusWaitSync,
					})
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (g *GithubUsecase) Download(ctx context.Context, data *release.Release) error {
	assets, err := g.asset.FindList(ctx, asset.AreaReleaseIDEQ(data.ID), asset.StatusEQ(asset.StatusWaitSync), asset.Limit(10), asset.Order("-id"))
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	for _, result := range assets {
		result := result
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := g.filesystem.Download(ctx, result.DownloadURL, result.FilePath); err != nil {
				g.logger.Error("download asset file failed", zap.Error(err), zap.String("filename", result.FilePath), zap.Uint64("areaReleaseAreaID", result.ID))
				return
			}
			result.Status = asset.StatusWaitLoaded
			if _, err1 := g.asset.Save(ctx, result); err1 != nil {
				g.logger.Error("save asset download status failed", zap.Error(err1), zap.Uint64("areaReleaseAreaID", result.ID))
				return
			}
		}()
	}
	wg.Wait()
	return nil
}

func (g *GithubUsecase) Loaded(ctx context.Context, data *release.Release) error {
	assets, err := g.asset.FindList(ctx, asset.StatusEQ(asset.StatusWaitLoaded), asset.Limit(10), asset.Order("-id"))
	if err != nil {
		return err
	}
	for _, result := range assets {
		result := result
		f, err := os.Open(result.FilePath)
		if err != nil {
			return err
		}
		head := make([]byte, 4)
		f.Read(head)
		headHex := 0
		for _, b := range head {
			headHex = headHex << 8
			headHex = headHex | int(b)
		}
		if err := f.Close(); err != nil {
			return err
		}
		switch headHex {
		case 930790575:
			// 解压文件
			absPath, _ := filepath.Abs(result.FilePath)
			if err := zip7.Extract(absPath, "-o"+filepath.Base(absPath)); err != nil {
				return err
			}
		case 4022058857:
			// 转到导入
			if err := g.LoadedFile(ctx, result.FilePath); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *GithubUsecase) LoadedFile(ctx context.Context, filename string) error {
	return nil
}

func NewGithubUsecase(
	app *app.App,
	logger *log.Logger,
	transaction transaction.Transaction,
	filesystem *filesystem.FileSystem,
	xiangyuecn XiangyuecnRepository,
	release release.ManageRepo,
	asset asset.ManageRepo,
) *GithubUsecase {
	return &GithubUsecase{
		app:         app,
		logger:      logger,
		filesystem:  filesystem,
		xiangyuecn:  xiangyuecn,
		release:     release,
		asset:       asset,
		transaction: transaction,
	}
}
