package release

import (
	"context"
	"net/url"

	"github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/google/go-github/v45/github"
)

type XiangyuecnRepository interface {
	GetLatestRelease(ctx context.Context) (*github.RepositoryRelease, error)
}

type GithubUsecase struct {
	xiangyuecn XiangyuecnRepository
	release    ManageRepo
	asset      asset.ManageRepo
}

func (g *GithubUsecase) Load(ctx context.Context) error {
	rep, err := g.xiangyuecn.GetLatestRelease(ctx)
	if err != nil {
		return err
	}
	_, err = g.release.Save(ctx, &Release{
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
	// var assets []*asset.Asset
	for _, repAsset := range rep.Assets {
		downloadUrl := repAsset.GetBrowserDownloadURL()
		_, err := url.Parse(downloadUrl)
		if err != nil {
			return err
		}
		// var filepath string = uri.String()
		// filepath = pkgfilepath.RelativePath(filepath.Join(pkgfilepath.RuntimePath(), release.ReleaseNodeID, filepath.Base(uri.Path)))
		// asset, err := g.asset.Save(ctx, &asset.Asset{
		// AreaReleaseId: release.ID,
		// AssetName:     repAsset.GetName(),
		// AssetId:       uint64(repAsset.GetID()),
		// AssetLabel:    repAsset.GetLabel(),
		// AssetState:    repAsset.GetState(),
		// FileSize:      uint32(repAsset.GetSize()),
		// FilePath:      filepath,
		// DownloadUrl:   downloadUrl,
		// Status:        nil,
		// })
		if err != nil {
			return err
		}
		// assets = append(assets, asset)
	}
	return nil
}

func NewGithubUsecase(
	xiangyuecn XiangyuecnRepository,
	release ManageRepo,
	asset asset.ManageRepo,
) *GithubUsecase {
	return &GithubUsecase{
		xiangyuecn: xiangyuecn,
		release:    release,
		asset:      asset,
	}
}
