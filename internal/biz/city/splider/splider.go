package splider

import (
	"context"

	"github.com/cnartlu/area-service/errors"
)

type SpliderRepo interface {
	Count(ctx context.Context, options ...Query) int
	FindList(ctx context.Context, options ...Query) ([]*Splider, error)
	FindOne(ctx context.Context, options ...Query) (*Splider, error)
	Save(ctx context.Context, data *Splider) (*Splider, error)
	Remove(ctx context.Context, options ...Query) error
}

type SpliderUsecase struct {
	repo SpliderRepo
}

func (r *SpliderUsecase) List(ctx context.Context) ([]*Splider, error) {
	return r.repo.FindList(ctx)
}

func (r *SpliderUsecase) FindOne(ctx context.Context, id int) (*Splider, error) {
	return r.repo.FindOne(ctx, IDEQ(id))
}

func (r *SpliderUsecase) FindOneWithInstance(ctx context.Context, queries ...Query) (*Splider, error) {
	return r.repo.FindOne(ctx, queries...)
}

func (r *SpliderUsecase) Save(ctx context.Context, data *Splider) (*Splider, error) {
	if data == nil {
		return nil, errors.ErrorParamMissing("save")
	}
	return r.repo.Save(ctx, data)
}

func (r *SpliderUsecase) Remove(ctx context.Context) error {
	return r.repo.Remove(ctx)
}

func NewSpliderUsecase(repo SpliderRepo) *SpliderUsecase {
	return &SpliderUsecase{
		repo: repo,
	}
}
