package asset

import (
	"context"
)

type ManageRepo interface {
	Count(ctx context.Context, options ...Option) int
	FindList(ctx context.Context, options ...Option) ([]*Asset, error)
	FindOne(ctx context.Context, options ...Option) (*Asset, error)
	Save(ctx context.Context, data *Asset) (*Asset, error)
	Remove(ctx context.Context, options ...Option) error
}

type ManageUsecase struct {
	repo ManageRepo
}

func (m *ManageUsecase) List(ctx context.Context) ([]*Asset, error) {
	return m.repo.FindList(ctx)
}

func (m *ManageUsecase) FindOne(ctx context.Context) (*Asset, error) {
	return m.repo.FindOne(ctx)
}

func NewManageUsecase(
	repo ManageRepo,
) *ManageUsecase {
	return &ManageUsecase{
		repo: repo,
	}
}
