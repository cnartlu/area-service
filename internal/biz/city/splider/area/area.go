package area

import (
	"context"

	"github.com/cnartlu/area-service/errors"
)

type AreaRepo interface {
	Count(ctx context.Context, options ...Query) int
	FindList(ctx context.Context, options ...Query) ([]*Area, error)
	FindOne(ctx context.Context, options ...Query) (*Area, error)
	Save(ctx context.Context, data *Area) (*Area, error)
	Remove(ctx context.Context, options ...Query) error
}

type AreaUsecase struct {
	repo AreaRepo
}

func (a *AreaUsecase) List(ctx context.Context) ([]*Area, error) {
	return a.repo.FindList(ctx)
}

func (a *AreaUsecase) FindOne(ctx context.Context, id int) (*Area, error) {
	return a.repo.FindOne(ctx, IDEQ(id))
}

func (a *AreaUsecase) FindOneWithInstance(ctx context.Context, queries ...Query) (*Area, error) {
	return a.repo.FindOne(ctx, queries...)
}

func (a *AreaUsecase) Save(ctx context.Context, data *Area) (*Area, error) {
	if data == nil {
		return nil, errors.ErrorParamMissing("save")
	}
	return a.repo.Save(ctx, data)
}

func (a *AreaUsecase) Create(ctx context.Context, data *Area) error {
	return nil
}

func (a *AreaUsecase) Update(ctx context.Context, data *Area) error {
	return nil
}

func (a *AreaUsecase) Remove(ctx context.Context) error {
	return a.repo.Remove(ctx)
}

func NewAreaUsecase(repo AreaRepo) *AreaUsecase {
	return &AreaUsecase{
		repo: repo,
	}
}
