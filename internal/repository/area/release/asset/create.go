package asset

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type Creator interface {
	// Create 创建资源记录
	Create(ctx context.Context, data *ent.AreaReleaseAsset) (*ent.AreaReleaseAsset, error)
}

// Create 创建记录
func (r *Repository) Create(ctx context.Context, data *ent.AreaReleaseAsset) (*ent.AreaReleaseAsset, error) {
	return r.ent.AreaReleaseAsset.Create().
		Save(ctx)
}
