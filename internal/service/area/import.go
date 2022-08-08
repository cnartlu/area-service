package area

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type Importer interface {
	Import(ctx context.Context, asset *ent.AreaReleaseAsset) error
}

func (s *Service) Import(ctx context.Context, asset *ent.AreaReleaseAsset) error {
	// 导入的文件可能是csv的文件格式，也可能是gz格式

	return nil
}
