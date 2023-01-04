package polygon

import (
	"context"

	"github.com/cnartlu/area-service/errors"
)

type PolygonRepo interface {
	Count(ctx context.Context, options ...Query) int
	FindList(ctx context.Context, options ...Query) ([]*Polygon, error)
	FindOne(ctx context.Context, options ...Query) (*Polygon, error)
	Save(ctx context.Context, data *Polygon) (*Polygon, error)
	Remove(ctx context.Context, options ...Query) error
}

type PolygonUsecase struct {
	repo PolygonRepo
}

func (a *PolygonUsecase) List(ctx context.Context) ([]*Polygon, error) {
	return a.repo.FindList(ctx)
}

func (a *PolygonUsecase) FindOne(ctx context.Context, id int) (*Polygon, error) {
	return a.repo.FindOne(ctx, IDEQ(id))
}

func (a *PolygonUsecase) FindOneWithInstance(ctx context.Context, queries ...Query) (*Polygon, error) {
	return a.repo.FindOne(ctx, queries...)
}

func (a *PolygonUsecase) Save(ctx context.Context, data *Polygon) (*Polygon, error) {
	if data == nil {
		return nil, errors.ErrorParamMissing("save")
	}
	return a.repo.Save(ctx, data)
}

func (a *PolygonUsecase) Create(ctx context.Context, data *Polygon) error {
	return nil
}

func (a *PolygonUsecase) Update(ctx context.Context, data *Polygon) error {
	return nil
}

func (a *PolygonUsecase) Remove(ctx context.Context) error {
	return a.repo.Remove(ctx)
}

func NewPolygonUsecase(repo PolygonRepo) *PolygonUsecase {
	return &PolygonUsecase{
		repo: repo,
	}
}
