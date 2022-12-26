package github

import (
	"context"

	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/component/filesystem"
	"github.com/cnartlu/area-service/errors"
	"github.com/cnartlu/area-service/internal/biz/city/splider"
	"github.com/cnartlu/area-service/internal/biz/city/splider/area"
	"github.com/cnartlu/area-service/internal/biz/city/splider/asset"
	"github.com/cnartlu/area-service/internal/biz/transaction"
	"github.com/google/go-github/v45/github"
)

type GithubRepo interface {
	GetLatestRelease(ctx context.Context, owner string, repo string) (*github.RepositoryRelease, error)
}

type GithubUsecase struct {
	account        Account
	repo           GithubRepo
	app            *app.App
	filesystem     *filesystem.FileSystem
	transaction    transaction.Transaction
	spliderUsecase *splider.SpliderUsecase
	assetUsecase   *asset.AssetUsecase
	areaUsecase    *area.AreaUsecase
}

func (g *GithubUsecase) LoadLatestRelease(ctx context.Context) (*Github, error) {
	var result Github
	release, err := g.repo.GetLatestRelease(ctx, g.account.User, g.account.Repo)
	if err != nil {
		return nil, err
	}
	spliderResponse, err := g.spliderUsecase.FindOneWithInstance(ctx,
		splider.SourceIDEQ(uint64(release.GetID())),
		splider.SourceEQ(Source),
		splider.OwnerEQ(g.account.User),
		splider.RepoEQ(g.account.Repo),
	)
	if err != nil {
		if !errors.IsDataNotFound(err) {
			return nil, err
		}
		spliderResponse, err = g.spliderUsecase.Save(ctx, &splider.Splider{
			Source:      Source,
			Owner:       g.account.User,
			Repo:        g.account.Repo,
			SourceID:    uint64(release.GetID()),
			Title:       release.GetName(),
			Draft:       release.GetDraft(),
			PreRelease:  release.GetPrerelease(),
			PublishedAt: release.GetPublishedAt().Time,
			Status:      splider.StatusWaitSync,
			CreatedAt:   release.GetCreatedAt().Time,
		})
		if err != nil {
			return nil, err
		}
	}
	result.Splider = spliderResponse
	if len(release.Assets) > 0 {
		var assets = make([]*asset.Asset, len(release.Assets))
		for k, data := range release.Assets {
			data := data
			assetData, err := g.assetUsecase.FindOneWithInstance(ctx, asset.CitySpliderIDEQ(spliderResponse.ID), asset.SourceIDEQ(uint64(data.GetID())))
			if err != nil {
				if !errors.IsDataNotFound(err) {
					return nil, err
				}
				assetData = &asset.Asset{
					CitySpliderID: spliderResponse.ID,
					SourceID:      uint64(data.GetID()),
					FileTitle:     data.GetName(),
					FilePath:      data.GetBrowserDownloadURL(),
					FileSize:      uint(data.GetSize()),
					Status:        asset.StatusWaitSync,
				}
				err = g.assetUsecase.Create(ctx, assetData)
				if err != nil {
					return nil, err
				}
			}
			assets[k] = assetData
		}
		result.Assets = assets
	}
	return &result, nil
}

func NewGithubRepoUsecase(
	repo GithubRepo,
	app *app.App,
	filesystem *filesystem.FileSystem,
	transaction transaction.Transaction,
	spliderUsecase *splider.SpliderUsecase,
	assetUsecase *asset.AssetUsecase,
	areaUsecase *area.AreaUsecase,
) *GithubUsecase {
	var account = Account{
		User: "xiangyuecn",
		Repo: "AreaCity-JsSpider-StatsGov",
	}
	return &GithubUsecase{
		account:        account,
		repo:           repo,
		app:            app,
		filesystem:     filesystem,
		transaction:    transaction,
		spliderUsecase: spliderUsecase,
		assetUsecase:   assetUsecase,
		areaUsecase:    areaUsecase,
	}
}
