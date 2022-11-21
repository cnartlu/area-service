package github

import (
	"context"

	bizAreaRelease "github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/data/ent"

	"github.com/google/go-github/v45/github"
)

var (
	_ bizAreaRelease.XiangyuecnRepository = (*XiangyuecnRepo)(nil)
)

type XiangyuecnRepo struct {
	ent   *ent.Client
	g     *github.Client
	owner string
	repo  string
}

func (r *XiangyuecnRepo) GetLatestRelease(ctx context.Context) (*github.RepositoryRelease, error) {
	var owner = r.owner
	var repo = r.repo
	rep, _, err := r.g.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	return rep, nil
}

func NewXiangyuecnRepo(ent *ent.Client, g *github.Client) *XiangyuecnRepo {
	if g == nil {
		g = github.NewClient(nil)
	}
	return &XiangyuecnRepo{
		ent:   ent,
		g:     g,
		owner: "xiangyuecn",
		repo:  "AreaCity-JsSpider-StatsGov",
	}
}
