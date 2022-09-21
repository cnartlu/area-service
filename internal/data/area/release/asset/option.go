package asset

import (
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/areareleaseasset"
)

type option func(*ent.AreaReleaseAssetQuery)

func WithIDEQ(id uint64) option {
	return func(query *ent.AreaReleaseAssetQuery) {
		query.Where(areareleaseasset.IDEQ(id))
	}
}

func WithT() option {
	return func(query *ent.AreaReleaseAssetQuery) {

	}
}
