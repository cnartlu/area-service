package release

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type Querier interface {
	FindOneWithLastAdded(ctx context.Context) (*ent.AreaRelease, error)
}

func (s *Service) FindOneWithLastAdded(ctx context.Context) (*ent.AreaRelease, error) {
	return nil, nil
}
