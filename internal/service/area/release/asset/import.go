package asset

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type Importer interface {
	Import(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset) error
}

func (s *Service) Import(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset) error {
	return s.importerService.Import(ctx, areaReleaseAsset, nil)
}
