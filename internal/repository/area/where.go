package area

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/area"
)

type Querier interface {
	// FindList 列表查询
	FindList(ctx context.Context, params *FindListParam) ([]*ent.Area, error)
	// FindOneById 根据 id 查询详情
	FindOneById(ctx context.Context, id uint64, columns []string) (*ent.Area, error)
}

// FindList 列表查询
func (r *Repository) FindList(ctx context.Context, params *FindListParam) ([]*ent.Area, error) {
	client := r.ent
	query := client.Area.Query().
		Where(area.DeleteTimeEQ(0))
	if params == nil {
		query.Where(area.ParentIDEQ(0))
	} else {
		query.Where(area.ParentIDEQ(params.ParentID))
		if params.Keyword != "" {
			query.Where(area.TitleContains(params.Keyword))
		}
	}
	query.Order(ent.Desc(area.FieldID))
	return query.All(ctx)
}

// FindOneById 根据 id 查询详情
func (r *Repository) FindOneById(ctx context.Context, id uint64, columns []string) (*ent.Area, error) {
	client := r.ent
	query := client.Area.
		Query().
		Where(area.IDEQ(id))
	if len(columns) > 0 {
		query.Select(columns...)
	}
	return query.First(ctx)
}

// FindOneByRegionID 根据RegionID和level查找详情
func (r *Repository) FindOneByRegionID(ctx context.Context, regionID string, level uint8) (*ent.Area, error) {
	client := r.ent
	return client.Area.
		Query().
		Where(area.RegionIDEQ(regionID), area.LevelEQ(level)).
		First(ctx)
}
