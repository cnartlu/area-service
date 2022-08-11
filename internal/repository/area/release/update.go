package release

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type Updater interface {
	// Update 更新资源记录
	Update(ctx context.Context, data *ent.AreaRelease) (*ent.AreaRelease, error)
}

var _ Updater = (*Repository)(nil)

// Update 创建记录
func (r *Repository) Update(ctx context.Context, data *ent.AreaRelease) (*ent.AreaRelease, error) {
	return data.Update().
		SetOwner(data.Owner).
		SetRepo(data.Repository).
		SetReleaseID(data.ReleaseID).
		SetReleaseName(data.ReleaseName).
		SetReleaseNodeID(data.ReleaseNodeID).
		SetReleasePublishedAt(data.ReleasePublishedAt).
		SetReleaseContent(data.ReleaseContent).
		SetStatus(data.Status).
		Save(ctx)
}
