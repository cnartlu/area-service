package release

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/arearelease"
)

type Querier interface {
	// FindList 查找列表
	FindList(ctx context.Context, params *FindListParam) ([]*ent.AreaRelease, error)
	FindByReleaseID(ctx context.Context, releaseID uint64) (*ent.AreaRelease, error)
	FindOneWithLastRecord(ctx context.Context) (*ent.AreaRelease, error)
}

var _ Querier = (*Repository)(nil)

// FindList 查找数据列表
func (r *Repository) FindList(ctx context.Context, params *FindListParam) ([]*ent.AreaRelease, error) {
	query := r.ent.AreaRelease.Query()
	if params != nil {
		if params.Keyword != "" {
			query.Where(arearelease.ReleaseNameContains(params.Keyword))
		}
		if params.Owner != "" {
			query.Where(arearelease.OwnerContains(params.Owner))
		}
		if params.Repository != "" {
			query.Where(arearelease.RepoContains(params.Repository))
		}
		if params.Pagination {
			query.Offset(params.Offset).Limit(params.Limit)
		}
	}
	return query.All(ctx)
}

// FindByReleaseID 通过releaseID查找记录
func (r *Repository) FindByReleaseID(ctx context.Context, releaseID uint64) (*ent.AreaRelease, error) {
	return r.ent.AreaRelease.Query().
		Where(arearelease.ReleaseIDEQ(releaseID)).
		Order(ent.Desc(arearelease.FieldID)).
		First(ctx)
}

// FindOneWithLastRecord 查找最后一条记录
func (r *Repository) FindOneWithLastRecord(ctx context.Context) (*ent.AreaRelease, error) {
	return r.ent.AreaRelease.Query().
		Order(ent.Desc(arearelease.FieldID)).
		First(ctx)
}
