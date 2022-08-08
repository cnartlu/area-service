package area

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/area"
)

type Querier interface {
	// FindList 列表查询
	FindList(ctx context.Context, param FindListParam, columns []string, order string) ([]*ent.Area, error)
	// FindOneById 根据 id 查询详情
	FindOneById(ctx context.Context, id uint64, columns []string) (*ent.Area, error)
}

type FindListParam struct {
	Kw       string
	ParentID uint64
}

// FindList 列表查询
func (r *Repository) FindList(ctx context.Context, param FindListParam, columns []string, order string) ([]*ent.Area, error) {
	client := r.ent
	query := client.Area.Query().
		Where(area.ParentIDEQ(param.ParentID), area.DeleteTimeEQ(0))
	if param.Kw != "" {
		query.Where(area.TitleContains(param.Kw))
	}
	if len(columns) > 0 {
		query.Select(columns...)
	}
	switch order {
	case area.FieldPinyin:
		query.Order(ent.Desc(area.FieldPinyin))
	default:
		query.Order(ent.Desc(area.FieldID))
	}
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
