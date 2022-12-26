package github

import (
	"context"

	bizgithub "github.com/cnartlu/area-service/internal/biz/city/github"
)

type Handler interface {
	Load(context.Context) error
}

type handler struct {
	github *bizgithub.GithubUsecase
}

func (h *handler) Load(ctx context.Context) error {
	latestRelease, err := h.github.LoadLatestRelease(ctx)
	if err != nil {
		return err
	}
	err = h.github.WriteByGithub(ctx, latestRelease)
	if err != nil {
		return err
	}
	return nil
}

func NewHandler(
	github *bizgithub.GithubUsecase,
) Handler {
	return &handler{
		github: github,
	}
}
