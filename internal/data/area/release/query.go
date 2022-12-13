package release

import (
	"strings"

	bizrelease "github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/arearelease"
)

var _ bizrelease.Inquirer = new(queryOption)

type queryOption struct {
	query *ent.AreaReleaseQuery
}

func (o queryOption) Offset(offset int) {
	o.query.Offset(offset)
}

func (o queryOption) Limit(limit int) {
	o.query.Limit(limit)
}

func (o queryOption) Order(order string) {
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

func (o queryOption) IDEQ(id uint64) {
	o.query.Where(arearelease.IDEQ(id))
}

func (o queryOption) IDIn(ids ...uint64) {
	o.query.Where(arearelease.IDIn(ids...))
}

func (o queryOption) ReleaseIDEQ(releaseID uint64) {
	o.query.Where(arearelease.ReleaseIDEQ(releaseID))
}

func NewQuery(query *ent.AreaReleaseQuery) queryOption {
	return queryOption{query: query}
}
