package asset

import (
	"strings"

	bizasset "github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/areareleaseasset"
)

var _ bizasset.OptionInterface = (*option)(nil)

type option struct {
	query *ent.AreaReleaseAssetQuery
}

func (o *option) Offset(offset int) {
	o.query.Offset(offset)
}

func (o *option) Limit(limit int) {
	o.query.Limit(limit)
}

func (o *option) Order(order string) {
	if order == "" {
		return
	}
	orders := strings.Split(order, ",")
	for _, v := range orders {
		v := strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if strings.HasPrefix(v, "-") {
			o.query.Order(ent.Desc(v[1:]))
		} else {
			o.query.Order(ent.Asc(v))
		}
	}
}

func (o *option) IDEQ(id uint64) {
	o.query.Where(areareleaseasset.IDEQ(id))
}

func (o *option) IDIn(ids ...uint64) {
	o.query.Where(areareleaseasset.IDIn(ids...))
}

func (o *option) AreaReleaseIDEQ(id uint64) {
	o.query.Where(areareleaseasset.AreaReleaseIDEQ(id))
}

func (o *option) StatusEQ(status bizasset.Status) {
	o.query.Where(areareleaseasset.StatusEQ(uint8(status)))
}

func newOption(query *ent.AreaReleaseAssetQuery) *option {
	return &option{query: query}
}
