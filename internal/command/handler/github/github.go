package github

import (
	"context"

	bizgithub "github.com/cnartlu/area-service/internal/biz/github"
)

type Handler interface {
	Load(context.Context) error
}

type handler struct {
	github *bizgithub.GithubUsecase
}

func (h *handler) Load(ctx context.Context) error {
	r, err := h.github.GetLatestRelease(ctx)
	if err != nil {
		return err
	}
	err = h.github.Download(ctx, r)
	if err != nil {
		return err
	}
	// 打开下载得文件并拉取数据
	return nil
}

func NewHandler(
	github *bizgithub.GithubUsecase,
) Handler {
	return &handler{
		github: github,
	}
}
