package asset

import (
	"strings"

	bizAsset "github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/areareleaseasset"
)

var _ bizAsset.OptionInterface = (*option)(nil)

type option struct {
	*ent.AreaReleaseAssetQuery
}

func (o *option) Offset(offset int) {
	o.AreaReleaseAssetQuery.Offset(offset)
}

func (o *option) Limit(limit int) {
	o.AreaReleaseAssetQuery.Limit(limit)
}

func (o *option) Order(order string) {
	if order == "" {
		order = "-id"
	}
	orders := strings.Split(order, ",")
	for _, v := range orders {
		v := strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if strings.HasSuffix(v, "-") {
			o.AreaReleaseAssetQuery.Order(ent.Desc(v[1:]))
		} else {
			o.AreaReleaseAssetQuery.Order(ent.Asc(v))
		}
	}
}

func (o *option) IDEQ(id uint64) {
	o.AreaReleaseAssetQuery.Where(areareleaseasset.IDEQ(id))
}

func (o *option) IDIn(ids ...uint64) {
	o.AreaReleaseAssetQuery.Where(areareleaseasset.IDIn(ids...))
}

func newOption(query *ent.AreaReleaseAssetQuery) *option {
	return &option{query}
}
