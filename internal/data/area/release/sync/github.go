package sync

import (
	"context"
	"net/url"

	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/arearelease"
	"github.com/cnartlu/area-service/internal/data/ent/areareleaseasset"
	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
)

type GithubRepository interface {
	// LoadRemoteLatestRelease 远端拉取最后的资源
	LoadRemoteLatestRelease(ctx context.Context, owner, repo string) (*ent.AreaRelease, error)
}

var (
	_            GithubRepository = (*GithubRepo)(nil)
	defaultOwner                  = "xiangyuecn"
	defaultRepo                   = "AreaCity-JsSpider-StatsGov"
)

type GithubRepo struct {
	ent *ent.Client
	g   *github.Client
}

func (r *GithubRepo) LoadRemoteLatestRelease(ctx context.Context, owner, repo string) (*ent.AreaRelease, error) {
	if owner == "" {
		owner = defaultOwner
		repo = defaultRepo
	}
	if repo == "" {
		repo = defaultRepo
	}
	rep, _, err := r.g.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to getLatestRelease by github client")
	}
	areaRelease, err := r.ent.AreaRelease.Query().
		Where(arearelease.ReleaseIDEQ(uint64(rep.GetID()))).
		Order(ent.Desc(arearelease.FieldID)).
		First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, errors.Wrap(err, "failed to FindByReleaseID by release repository")
		}
		areaRelease, err = r.ent.AreaRelease.Create().
			SetOwner(owner).
			SetRepo(repo).
			SetReleaseID(uint64(rep.GetID())).
			SetReleaseName(rep.GetName()).
			SetReleaseNodeID(rep.GetNodeID()).
			SetReleasePublishedAt(uint64(rep.GetPublishedAt().Unix())).
			SetReleaseContent(rep.GetBody()).
			SetStatus(arearelease.StatusWaitLoaded).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create area.release repository")
		}
	}
	if areaRelease.Status == arearelease.StatusWaitLoaded {
		createBulks := []*ent.AreaReleaseAssetCreate{}
		for _, asset := range rep.Assets {
			_, err := r.ent.AreaReleaseAsset.Query().
				Where(areareleaseasset.AreaReleaseIDEQ(areaRelease.ID), areareleaseasset.AssetIDEQ(uint64(asset.GetID()))).
				First(ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return nil, errors.Wrap(err, "failed to FindOneByAssetID by area.release.asset repository")
				}
				downloadUrl := asset.GetBrowserDownloadURL()
				_, err := url.Parse(downloadUrl)
				if err != nil {
					return nil, errors.Wrap(err, "failed to parse url")
				}
				createBulks = append(createBulks, r.ent.AreaReleaseAsset.Create().
					SetAreaReleaseID(areaRelease.ID).
					SetAssetID(uint64(asset.GetID())).
					SetAssetName(asset.GetName()).
					SetAssetLabel(asset.GetLabel()).
					SetAssetState(asset.GetState()).
					SetDownloadURL(downloadUrl).
					// SetFilePath(utils.RelativePath(filepath.Join(utils.RuntimePath(), areaRelease.ReleaseNodeID, filepath.Base(uri.Path)))).
					SetFileSize(uint(asset.GetSize())).
					SetStatus(areareleaseasset.StatusWaitLoaded))
			}
		}
		if len(createBulks) > 0 {
			_, err = r.ent.AreaReleaseAsset.CreateBulk(createBulks...).Save(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "failed to create area.release.asset repository")
			}
		}
	}

	return areaRelease, nil
}

func NewGithubRepo(ent *ent.Client) *GithubRepo {
	return &GithubRepo{
		ent: ent,
		g:   github.NewClient(nil),
	}
}
