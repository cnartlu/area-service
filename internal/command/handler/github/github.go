package github

import (
	"context"
	"os"

	bizgithub "github.com/cnartlu/area-service/internal/biz/city/github"
	"go.uber.org/zap"
)

type Handler interface {
	Latest(context.Context) error
	Load(ctx context.Context, ids ...int) error
}

type handler struct {
	github *bizgithub.GithubUsecase
	logger *zap.Logger
}

func (h *handler) Latest(ctx context.Context) error {
	_, err := h.github.LatestRelease(ctx)
	if err != nil {
		h.logger.Error("latest error", zap.Error(err), zap.StackSkip("stack", 1))
		os.Exit(2)
	}
	return nil
}

func (h *handler) Load(ctx context.Context, ids ...int) error {
	err := h.github.LoadLatestRelease(ctx)
	return err
}

func NewHandler(
	github *bizgithub.GithubUsecase,
	logger *zap.Logger,
) Handler {
	return &handler{
		github: github,
		logger: logger,
	}
}
