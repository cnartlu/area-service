package asset

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/arearelease"
	"github.com/cnartlu/area-service/internal/component/ent/areareleaseasset"
)

type Querier interface {
	// FindList 查找列表
	FindList(ctx context.Context, params *FindListParam) ([]*ent.AreaReleaseAsset, error)
	FindOneByID(ctx context.Context, id uint64) (*ent.AreaReleaseAsset, error)
	FindOneByAssetID(ctx context.Context, assetID uint64) (*ent.AreaReleaseAsset, error)
}

var _ Querier = (*Repository)(nil)

func (r *Repository) FindList(ctx context.Context, params *FindListParam) ([]*ent.AreaReleaseAsset, error) {
	query := r.ent.AreaReleaseAsset.Query()
	if params != nil {
		if params.AreaReleaseID > 0 {
			query.Where(areareleaseasset.AreaReleaseIDEQ(params.AreaReleaseID))
		}
		if params.Status != "" {
			query.Where(areareleaseasset.StatusEQ(areareleaseasset.Status(params.Status)))
		}
		if params.Keyword == "" {
			query.Where(areareleaseasset.AssetNameContains(params.Keyword))
		}
		if params.Pagination {
			query.Offset(params.Offset).Limit(params.Limit)
		}
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
