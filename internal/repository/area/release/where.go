package release

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/arearelease"
)

type Querier interface {
	// FindList 查找列表
	FindList(ctx context.Context, params FindListParam, columns []string) ([]*ent.AreaRelease, error)
	FindByReleaseID(ctx context.Context, releaseID uint64) (*ent.AreaRelease, error)
}

type FindListParam struct {
	Kw       string
	ParentID uint64
	// 偏移
	Offset *int64
	// 限量
	Limit int64
}

func (r *Repository) FindList(ctx context.Context, params FindListParam, columns []string) ([]*ent.AreaRelease, error) {
	// query := r.ent.AreaRelease.Query()

	if params.Kw == "" {
		// query.Where(arearelease.T)
	}
	return nil, nil
}

// FindByReleaseID 通过releaseID查找记录
func (r *Repository) FindByReleaseID(ctx context.Context, releaseID uint64) (*ent.AreaRelease, error) {
	return r.ent.AreaRelease.Query().
		Where(arearelease.ReleaseIDEQ(releaseID)).
		Order(ent.Desc(arearelease.FieldID)).
		First(ctx)
}
