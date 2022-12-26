package asset

import (
	bizasset "github.com/cnartlu/area-service/internal/biz/city/splider/asset"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/cityspliderasset"
)

var _ bizasset.Inquirer = new(queryOption)

type queryOption struct {
	query *ent.CitySpliderAssetQuery
}

func (o queryOption) Offset(offset int) {
	o.query.Offset(offset)
}

func (o queryOption) Limit(limit int) {
	o.query.Limit(limit)
}

func (o queryOption) Order(order ...string) {
	return
	// if order == "" {
	// 	return
	// }
	// orders := strings.Split(order, ",")
	// for _, v := range orders {
	// 	v := strings.TrimSpace(v)
	// 	if v == "" {
	// 		continue
	// 	}
	// 	if strings.HasPrefix(v, "-") {
	// 		o.query.Order(ent.Desc(v[1:]))
	// 	} else {
	// 		o.query.Order(ent.Asc(v))
	// 	}
	// }
}

func (o queryOption) IDEQ(id int) {
	o.query.Where(cityspliderasset.IDEQ(id))
}

func (o queryOption) IDIn(ids ...int) {
	o.query.Where(cityspliderasset.IDIn(ids...))
}

func (o queryOption) CitySpliderIDEQ(id int) {
	o.query.Where(cityspliderasset.CitySpliderIDEQ(id))
}

func (o queryOption) SourceIDEQ(sourceID uint64) {
	o.query.Where(cityspliderasset.SourceIDEQ(sourceID))
}

func (o queryOption) SourceIDIn(sourceIDs []uint64) {
	o.query.Where(cityspliderasset.SourceIDIn(sourceIDs...))
}

func (o queryOption) StatusEQ(status bizasset.Status) {
	o.query.Where(cityspliderasset.StatusEQ(uint8(status)))
}

func NewQuery(query *ent.CitySpliderAssetQuery) queryOption {
	return queryOption{query: query}
}
