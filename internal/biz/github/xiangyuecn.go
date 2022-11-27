package github

import (
	"context"
	"net/url"
	"path/filepath"
	"sync"

	"github.com/cnartlu/area-service/api"
	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/biz/transaction"
)

type XiangyuecnRepository interface {
	GetLatestRelease(ctx context.Context) (*GithubRepositoryRelease, error)
}

type GithubUsecase struct {
	app         *app.App
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

func (g *GithubUsecase) Download(ctx context.Context, release *release.Release) error {
	assets, err := g.asset.FindList(ctx, asset.AreaReleaseIDEQ(release.ID), asset.StatusEQ(asset.StatusWaitSync))
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	for _, result := range assets {
		result := result
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = result
		}()
	}
	wg.Wait()
	return nil
}

func NewGithubUsecase(
	app *app.App,
	transaction transaction.Transaction,
	xiangyuecn XiangyuecnRepository,
	release release.ManageRepo,
	asset asset.ManageRepo,
) *GithubUsecase {
	return &GithubUsecase{
		app:         app,
		xiangyuecn:  xiangyuecn,
		release:     release,
		asset:       asset,
		transaction: transaction,
	}
}
