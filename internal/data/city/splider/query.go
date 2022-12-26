package asset

import (
	bizsplider "github.com/cnartlu/area-service/internal/biz/city/splider"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/citysplider"
)

var _ bizsplider.Inquirer = new(queryOption)

type queryOption struct {
	query *ent.CitySpliderQuery
}

func (o queryOption) Offset(offset int) {
	o.query.Offset(offset)
}

func (o queryOption) Limit(limit int) {
	o.query.Limit(limit)
}

func (o queryOption) Order(orders ...string) {
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
	o.query.Where(citysplider.IDEQ(id))
}

func (o queryOption) IDIn(ids ...int) {
	o.query.Where(citysplider.IDIn(ids...))
}

func (o queryOption) SourceEQ(source string) {
	o.query.Where(citysplider.SourceEQ(source))
}

func (o queryOption) OwnerEQ(owner string) {
	o.query.Where(citysplider.OwnerEQ(owner))
}

func (o queryOption) RepoEQ(repo string) {
	o.query.Where(citysplider.RepoEQ(repo))
}

func (o queryOption) SourceIDEQ(sourceID uint64) {
	o.query.Where(citysplider.SourceIDEQ(sourceID))
}

func (o queryOption) StatusEQ(status bizsplider.Status) {
	o.query.Where(citysplider.StatusEQ(uint8(status)))
}

func NewQuery(query *ent.CitySpliderQuery) queryOption {
	return queryOption{query: query}
}
