package release

import (
	"context"

	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/areareleaseasset"
	"github.com/go-redis/redis/v8"
)

type ImportRepository interface {
	Import(ctx context.Context, areaRelease *ent.AreaRelease) error
	ImportAsset(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset) error
}

var _ ImportRepository = (*ImportRepo)(nil)

type ImportRepo struct {
	ent *ent.Client
	rdb *redis.Client
}

func (r *ImportRepo) Import(ctx context.Context, areaRelease *ent.AreaRelease) error {
	assets, err := r.ent.AreaReleaseAsset.Query().
		Where(areareleaseasset.AreaReleaseIDEQ(areaRelease.ID)).
		All(ctx)
	if err != nil {
		return err
	}
	for _, asset := range assets {
		err := r.ImportAsset(ctx, asset)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ImportRepo) ImportAsset(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset) error {
	return nil
}

func NewImportRepo(ent *ent.Client, rdb *redis.Client) *ImportRepo {
	return &ImportRepo{
		ent: ent,
		rdb: rdb,
	}
}
