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

func (a *AssetUsecase) FindList(ctx context.Context) ([]*Asset, error) {
	return a.repo.FindList(ctx)
}

func (a *AssetUsecase) FindOne(ctx context.Context, id uint64) (*Asset, error) {
	return a.repo.FindOne(ctx)
}

func (a *AssetUsecase) FindOneWithInstance(ctx context.Context, queries ...Query) (*Asset, error) {
	return a.repo.FindOne(ctx, queries...)
}

func (a *AssetUsecase) Create(ctx context.Context, data *Asset) error {
	var err error
	data, err = a.repo.Save(ctx, data)
	return err
}

func NewAssetUsecase(
	repo AssetRepo,
) *AssetUsecase {
	return &AssetUsecase{
		repo: repo,
	}
}
