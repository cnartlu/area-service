package release

import (
	"strings"

	bizrelease "github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/arearelease"
)

var _ bizrelease.OptionInterface = (*option)(nil)

type option struct {
	query *ent.AreaReleaseQuery
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
	o.query.Where(arearelease.IDEQ(id))
}

func (o *option) IDIn(ids ...uint64) {
	o.query.Where(arearelease.IDIn(ids...))
}

func (o *option) ReleaseIDEQ(releaseID uint64) {
	o.query.Where(arearelease.ReleaseIDEQ(releaseID))
}

func newOption(query *ent.AreaReleaseQuery) *option {
	return &option{query: query}
}
