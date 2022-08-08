package area

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type Deleter interface {
	// Delete 删除数据
	Delete(ctx context.Context, area *ent.Area) error
}

func (r *Repository) Delete(ctx context.Context, area *ent.Area) error {
	return r.ent.Area.DeleteOne(area).Exec(ctx)
}
