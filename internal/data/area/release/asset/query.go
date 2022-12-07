package asset

import (
	"strings"

	bizasset "github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/areareleaseasset"
)

var _ bizasset.Inquirer = (*queryOption)(nil)

type queryOption struct {
	query *ent.AreaReleaseAssetQuery
}

func (o *queryOption) Offset(offset int) {
	o.query.Offset(offset)
}

func (o *queryOption) Limit(limit int) {
	o.query.Limit(limit)
}

func (o *queryOption) Order(order string) {
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

func (o *queryOption) IDEQ(id uint64) {
	o.query.Where(areareleaseasset.IDEQ(id))
}

func (o *queryOption) IDIn(ids ...uint64) {
	o.query.Where(areareleaseasset.IDIn(ids...))
}

func (o *queryOption) AreaReleaseIDEQ(id uint64) {
	o.query.Where(areareleaseasset.AreaReleaseIDEQ(id))
}

func (o *queryOption) StatusEQ(status bizasset.Status) {
	o.query.Where(areareleaseasset.StatusEQ(uint8(status)))
}

func NewQuery(query *ent.AreaReleaseAssetQuery) queryOption {
	return queryOption{query: query}
}
