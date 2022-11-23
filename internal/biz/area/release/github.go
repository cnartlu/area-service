package release

import (
	"context"
	"net/url"
	"path/filepath"

	"github.com/cnartlu/area-service/api"
	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/google/go-github/v45/github"
)

type XiangyuecnRepository interface {
	GetLatestRelease(ctx context.Context) (*github.RepositoryRelease, error)
}

type GithubUsecase struct {
	app        *app.App
	xiangyuecn XiangyuecnRepository
	release    ManageRepo
	asset      asset.ManageRepo
}

func (g *GithubUsecase) Load(ctx context.Context) error {
	rep, err := g.xiangyuecn.GetLatestRelease(ctx)
	if err != nil {
		return err
	}
	// 这里需要开启事务
	release, err := g.release.FindOne(ctx, ReleaseIDEQ(uint64(rep.GetID())))
	if err != nil {
		if !api.IsDataNotFound(err) {
			return err
		}
		release, err = g.release.Save(ctx, &Release{
			Owner:              "",
			Repo:               "",
			ReleaseID:          uint64(rep.GetID()),
			ReleaseName:        rep.GetName(),
			ReleaseNodeID:      rep.GetNodeID(),
			ReleasePublishedAt: uint64(rep.GetPublishedAt().Unix()),
			ReleaseContent:     rep.GetBody(),
			Status:             StatusWaitSync,
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
			var filepath string = filepath.Join(g.app.GetShortRuntimePath(), release.Owner, release.ReleaseName, filepath.Base(uri.Path))
			_, err = g.asset.Save(ctx, &asset.Asset{
				AreaReleaseID: release.ID,
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
}

func (g *GithubUsecase) Download(ctx context.Context, release *Release) error {
	assets, err := g.asset.FindList(ctx, asset.AreaReleaseIDEQ(release.ID), asset.StatusEQ(asset.StatusWaitSync))
	if err != nil {
		return err
	}
	for _, result := range assets {
		result := result
		go func() {
			_ = result
		}()
	}
	return nil
}

func NewGithubUsecase(
	app *app.App,
	xiangyuecn XiangyuecnRepository,
	release ManageRepo,
	asset asset.ManageRepo,
) *GithubUsecase {
	return &GithubUsecase{
		app:        app,
		xiangyuecn: xiangyuecn,
		release:    release,
		asset:      asset,
	}
}
