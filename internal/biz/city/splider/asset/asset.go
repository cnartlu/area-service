package asset

import (
	"context"

	"github.com/cnartlu/area-service/errors"
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

func (a *AssetUsecase) FindList(ctx context.Context, param FindListParams) ([]*Asset, error) {
	queries := []Query{}
	if param.CitySpliderID != 0 {
		queries = append(queries, CitySpliderIDEQ(param.CitySpliderID))
	}
	if param.SourceID > 0 {
		queries = append(queries, SourceIDEQ(param.SourceID))
	}
	if param.Status != 0 {
		queries = append(queries, StatusEQ(Status(param.Status)))
	}
	if param.Page > 0 {
		queries = append(queries, Offset((param.Page-1)*param.PageSize))
	}
	if param.PageSize > 0 {
		queries = append(queries, Limit(param.PageSize))
	}
	return a.repo.FindList(ctx, queries...)
}

func (a *AssetUsecase) FindOne(ctx context.Context, id int) (*Asset, error) {
	return a.repo.FindOne(ctx, IDEQ(id))
}

func (a *AssetUsecase) FindOneWithInstance(ctx context.Context, queries ...Query) (*Asset, error) {
	return a.repo.FindOne(ctx, queries...)
}

func (a *AssetUsecase) Create(ctx context.Context, data *Asset) (*Asset, error) {
	data.ID = 0
	var err error
	data, err = a.repo.Save(ctx, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AssetUsecase) Update(ctx context.Context, data *Asset) (*Asset, error) {
	if data.ID < 1 {
		return nil, errors.ErrorParamMissing("missing unique identifier")
	}
	var err error
	data, err = a.repo.Save(ctx, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewAssetUsecase(
	repo AssetRepo,
) *AssetUsecase {
	return &AssetUsecase{
		repo: repo,
	}
}
