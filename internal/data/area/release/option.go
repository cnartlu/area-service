package release

import (
	"strings"

	bizrelease "github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/arearelease"
)

var _ bizrelease.OptionInterface = (*option)(nil)

type option struct {
	*ent.AreaReleaseQuery
}

func (o *option) Offset(offset int) {
	o.AreaReleaseQuery.Offset(offset)
}

func (o *option) Limit(limit int) {
	o.AreaReleaseQuery.Limit(limit)
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
			o.AreaReleaseQuery.Order(ent.Desc(v[1:]))
		} else {
			o.AreaReleaseQuery.Order(ent.Asc(v))
		}
	}
}

func (o *option) IDEQ(id uint64) {
	o.AreaReleaseQuery.Where(arearelease.IDEQ(id))
}

func (o *option) IDIn(ids ...uint64) {
	o.AreaReleaseQuery.Where(arearelease.IDIn(ids...))
}

func newOption(query *ent.AreaReleaseQuery) *option {
	return &option{query}
}
