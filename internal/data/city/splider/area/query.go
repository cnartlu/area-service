package area

import (
	bizarea "github.com/cnartlu/area-service/internal/biz/city/splider/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/cityspliderarea"
)

var _ bizarea.Inquirer = new(queryOption)

type queryOption struct {
	query *ent.CitySpliderAreaQuery
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
	o.query.Where(cityspliderarea.IDEQ(id))
}

func (o queryOption) IDIn(ids ...int) {
	o.query.Where(cityspliderarea.IDIn(ids...))
}

func (o queryOption) RegionIDEQ(regionID string) {
	o.query.Where(cityspliderarea.RegionIDEQ(regionID))
}

func (o queryOption) LevelEQ(level int) {
	o.query.Where(cityspliderarea.LevelEQ(uint8(level)))
}

func NewQuery(query *ent.CitySpliderAreaQuery) queryOption {
	return queryOption{query: query}
}
