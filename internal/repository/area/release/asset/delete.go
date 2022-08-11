package asset

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type Deleter interface {
	// DeleteOne 通过数据模型删除数据
	DeleteOne(ctx context.Context, model *ent.AreaReleaseAsset) error
	// DeleteOneByID 通过主键ID删除一项数据
	DeleteOneByID(ctx context.Context, id uint64) error
}

var _ Deleter = (*Repository)(nil)

func (r *Repository) DeleteOne(ctx context.Context, model *ent.AreaReleaseAsset) error {
	return r.ent.AreaReleaseAsset.DeleteOne(model).Exec(ctx)
}

func (r *Repository) DeleteOneByID(ctx context.Context, id uint64) error {
	return r.ent.AreaReleaseAsset.DeleteOneID(id).Exec(ctx)
}
