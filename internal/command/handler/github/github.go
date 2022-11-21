package github

import (
	"context"

	bizarearealease "github.com/cnartlu/area-service/internal/biz/area/release"
)

type Handler interface {
	Load(context.Context) error
}

type handler struct {
	github *bizarearealease.GithubUsecase
}

func (h *handler) Load(ctx context.Context) error {
	return h.github.Load(ctx)
}

func NewHandler(
	github *bizarearealease.GithubUsecase,
) Handler {
	return &handler{
		github: github,
	}
}
