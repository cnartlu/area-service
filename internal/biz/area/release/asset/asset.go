package asset

import (
	"context"
)

type AssetRepo interface {
	Count(ctx context.Context, options ...Query) int
	FindList(ctx context.Context, options ...Query) ([]*Asset, error)
	FindOne(ctx context.Context, options ...Query) (*Asset, error)
	Save(ctx context.Context, data *Asset) (*Asset, error)
	Remove(ctx context.Context, options ...Query) error
}

type AssetUsecase struct {
	repo AssetRepo
}

func (m *AssetUsecase) List(ctx context.Context) ([]*Asset, error) {
	return m.repo.FindList(ctx)
}

func (m *AssetUsecase) FindOne(ctx context.Context) (*Asset, error) {
	return m.repo.FindOne(ctx)
}

func NewAssetUsecase(
	repo AssetRepo,
) *AssetUsecase {
	return &AssetUsecase{
		repo: repo,
	}
}
