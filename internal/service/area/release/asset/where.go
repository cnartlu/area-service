package asset

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
)

type Querier interface {
	FindListByReleaseID(ctx context.Context, areaReleaseID uint64) ([]*ent.AreaReleaseAsset, error)
}

func (s *Service) FindListByReleaseID(ctx context.Context, areaReleaseID uint64) ([]*ent.AreaReleaseAsset, error) {
	var params = asset.FindListParam{
		AreaReleaseID: areaReleaseID,
	}
	return s.repo.FindList(ctx, params, []string{})
}
