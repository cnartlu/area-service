package area_release

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type RepositoryWhere interface {
	// FindList 查找列表
	FindList(ctx context.Context, params FindListParam, columns []string) (*ent.AreaRelease, error)
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
