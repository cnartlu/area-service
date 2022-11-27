package github

import (
	"context"

	bizgithub "github.com/cnartlu/area-service/internal/biz/github"

	"github.com/google/go-github/v45/github"
)

var (
	_ bizgithub.XiangyuecnRepository = (*XiangyuecnRepo)(nil)
)

type XiangyuecnRepo struct {
	g     *github.Client
	owner string
	repo  string
}

func (r *XiangyuecnRepo) GetLatestRelease(ctx context.Context) (*bizgithub.GithubRepositoryRelease, error) {
	var owner = r.owner
	var repo = r.repo
	rep, _, err := r.g.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	result := bizgithub.GithubRepositoryRelease{
		Owner:             owner,
		Repo:              repo,
		RepositoryRelease: rep,
	}
	return &result, nil
}

func NewXiangyuecnRepo(g *github.Client) *XiangyuecnRepo {
	if g == nil {
		g = github.NewClient(nil)
	}
	return &XiangyuecnRepo{
		g:     g,
		owner: "xiangyuecn",
		repo:  "AreaCity-JsSpider-StatsGov",
	}
}
