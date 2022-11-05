package release

import (
	"context"
	"net/url"
	"path/filepath"

	"github.com/cnartlu/area-service/internal/biz/area/release/asset"
	pkgfilepath "github.com/cnartlu/area-service/pkg/path"
	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
)

type ManagerUsecase struct {
	repo         Releasor
	assetRepo    asset.Asseter
	githubClient *github.Client
}

func (m *ManagerUsecase) List(ctx context.Context) ([]*Release, error) {
	return m.repo.FindList(ctx)
}

func (m *ManagerUsecase) PullTheLatestFromGithubRelease(ctx context.Context) (*ReleaseWithAssets, error) {
	owner := GITHUB_OWNER
	repo := GITHUB_REPOSITORY
	rep, _, err := m.githubClient.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to getLatestRelease by github client")
	}

	release, err := m.repo.Save(ctx, &Release{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to find latest release")
	}
	var assets []*asset.Asset
	for _, repAsset := range rep.Assets {
		downloadUrl := repAsset.GetBrowserDownloadURL()
		uri, err := url.Parse(downloadUrl)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse url")
		}
		asset, err := m.assetRepo.Save(ctx, &asset.Asset{
			AreaReleaseId: release.GetId(),
			AssetName:     repAsset.GetName(),
			AssetId:       uint64(repAsset.GetID()),
			AssetLabel:    repAsset.GetLabel(),
			AssetState:    repAsset.GetState(),
			FileSize:      uint32(repAsset.GetSize()),
			FilePath:      pkgfilepath.RelativePath(filepath.Join(pkgfilepath.RuntimePath(), release.GetNodeId(), filepath.Base(uri.Path))),
			DownloadUrl:   downloadUrl,
			Status:        nil,
		})
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	result := ReleaseWithAssets{
		Release: release,
		Assets:  assets,
	}
	return &result, nil
}

func NewManaement(repo Releasor, assetRepo asset.Asseter) *ManagerUsecase {
	return &ManagerUsecase{
		repo:      repo,
		assetRepo: assetRepo,
	}
}
