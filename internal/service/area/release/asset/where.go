package asset

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
)

type Querier interface {
}

type FindListParam struct {
}

func (s *Service) FindList(ctx context.Context, params FindListParam) ([]*ent.AreaReleaseAsset, error) {
	return s.repo.FindList(ctx, asset.FindListParam{}, []string{})
}
 