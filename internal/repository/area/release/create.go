package release

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type Creator interface {
	// Create 创建资源记录
	Create(ctx context.Context, data *ent.AreaRelease) (*ent.AreaRelease, error)
}

// Create 创建记录
func (r *Repository) Create(ctx context.Context, data *ent.AreaRelease) (*ent.AreaRelease, error) {
	return r.ent.AreaRelease.Create().
		SetOwner(data.Owner).
		SetRepo(data.Repo).
		SetReleaseID(data.ReleaseID).
		SetReleaseName(data.ReleaseName).
		SetReleaseNodeID(data.ReleaseNodeID).
		SetReleasePublishedAt(data.ReleasePublishedAt).
		SetReleaseContent(data.ReleaseContent).
		SetStatus(data.Status).
		Save(ctx)
}
