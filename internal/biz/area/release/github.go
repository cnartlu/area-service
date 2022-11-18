package release

import (
	"context"
	"net/url"

	"github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/data/github"
)

var _ GithubRepository = (*GithubUsecase)(nil)

type GithubRepository interface {
	Load(ctx context.Context) (*Release, error)
}

type GithubUsecase struct {
	repo      Releasor
	assetRepo asset.Asseter
	github    github.XiangyuecnRepository
	config    *config.Config
}

func (r *GithubUsecase) Load(ctx context.Context) (*Release, error) {
	rep, err := r.github.GetLatestRelease(ctx)
	if err != nil {
		return nil, err
	}
	release, err := r.repo.Save(ctx, &Release{
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
		return nil, err
	}
	var assets []*asset.Asset
	for _, repAsset := range rep.Assets {
		downloadUrl := repAsset.GetBrowserDownloadURL()
		uri, err := url.Parse(downloadUrl)
		if err != nil {
			return nil, err
		}
		var filepath string = uri.String()
		// filepath = pkgfilepath.RelativePath(filepath.Join(pkgfilepath.RuntimePath(), release.ReleaseNodeID, filepath.Base(uri.Path)))
		asset, err := r.assetRepo.Save(ctx, &asset.Asset{
			AreaReleaseId: release.ID,
			AssetName:     repAsset.GetName(),
			AssetId:       uint64(repAsset.GetID()),
			AssetLabel:    repAsset.GetLabel(),
			AssetState:    repAsset.GetState(),
			FileSize:      uint32(repAsset.GetSize()),
			FilePath:      filepath,
			DownloadUrl:   downloadUrl,
			Status:        nil,
		})
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return release, nil
}

func NewgithubUsecase(
	repo Releasor,
	assetRepo asset.Asseter,
	github github.XiangyuecnRepository,
) *GithubUsecase {
	return &GithubUsecase{
		repo:      repo,
		assetRepo: assetRepo,
		github:    github,
	}
}
