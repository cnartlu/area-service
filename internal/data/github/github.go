package github

import (
	"context"

	"github.com/google/go-github/v45/github"
)

type GithubRepo struct {
	g *github.Client
}

func (r *GithubRepo) GetLatestRelease(ctx context.Context, owner string, repo string) (*github.RepositoryRelease, error) {
	rep, _, err := r.g.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	return rep, nil
}

func NewGithubRepo(g *github.Client) *GithubRepo {
	if g == nil {
		g = github.NewClient(nil)
	}
	return &GithubRepo{
		g: g,
	}
}
