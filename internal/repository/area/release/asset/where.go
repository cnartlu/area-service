package asset

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/arearelease"
	"github.com/cnartlu/area-service/internal/component/ent/areareleaseasset"
)

type FindListParam struct {
	AreaReleaseID uint64
	Kw            string
	Status        string
	Offset        *int
	Limit         int
}

type Querier interface {
	// FindList 查找列表
	FindList(ctx context.Context, params FindListParam, columns []string) ([]*ent.AreaReleaseAsset, error)
	FindOneByID(ctx context.Context, id uint64) (*ent.AreaReleaseAsset, error)
	FindOneByAssetID(ctx context.Context, assetID uint64) (*ent.AreaReleaseAsset, error)
}

func (r *Repository) FindList(ctx context.Context, params FindListParam, columns []string) ([]*ent.AreaReleaseAsset, error) {
	query := r.ent.AreaReleaseAsset.Query()
	if params.AreaReleaseID > 0 {
		query.Where(areareleaseasset.AreaReleaseIDEQ(params.AreaReleaseID))
	}
	if params.Status != "" {
		query.Where(areareleaseasset.StatusEQ(areareleaseasset.Status(params.Status)))
	}
	if params.Kw == "" {
		query.Where(areareleaseasset.AssetNameContains(params.Kw))
	}
	if params.Offset != nil {
		query.Offset(*params.Offset)
		if params.Limit > 0 {
			query.Limit(params.Limit)
		}
	}
	if len(columns) > 0 {
		query.Select(columns...)
	}
	return query.All(ctx)
}

// FindByID 通过ID查询
func (r *Repository) FindOneByID(ctx context.Context, id uint64) (*ent.AreaReleaseAsset, error) {
	return r.ent.AreaReleaseAsset.Query().
		Where(areareleaseasset.IDEQ(id)).
		Order(ent.Desc(arearelease.FieldID)).
		First(ctx)
}

// FindByAssetID 通过assetID查找记录
func (r *Repository) FindOneByAssetID(ctx context.Context, assetID uint64) (*ent.AreaReleaseAsset, error) {
	return r.ent.AreaReleaseAsset.Query().
		Where(areareleaseasset.AssetIDEQ(assetID)).
		Order(ent.Desc(arearelease.FieldID)).
		First(ctx)
}
