package polygon

import (
	"strings"

	bizareapolygon "github.com/cnartlu/area-service/internal/biz/city/splider/area/polygon"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/cityspliderareapolygon"
)

var _ bizareapolygon.Inquirer = new(queryOption)

type queryOption struct {
	query *ent.CitySpliderAreaPolygonQuery
	ttl   *int
}

func (o queryOption) Cache(ttl int) {
	o.ttl = &ttl
}

func (o queryOption) Offset(offset int) {
	o.query.Offset(offset)
}

func (o queryOption) Limit(limit int) {
	o.query.Limit(limit)
}

func (o queryOption) Order(orders ...string) {
	if len(orders) > 0 {
		var lastOrderDesc = false
		var fields = []string{}
		for _, str := range orders {
			if strings.HasPrefix(str, "-") {
				if !lastOrderDesc {
					lastOrderDesc = true
					o.query.Order(ent.Asc(fields...))
					fields = []string{}
				}
				fields = append(fields, str[1:])
			} else {
				if lastOrderDesc {
					lastOrderDesc = false
					o.query.Order(ent.Desc(fields...))
					fields = []string{}
				}
				fields = append(fields, str)
			}
		}
	}
}

func (o queryOption) IDEQ(id int) {
	o.query.Where(cityspliderareapolygon.IDEQ(id))
}

func (o queryOption) IDIn(ids ...int) {
	o.query.Where(cityspliderareapolygon.IDIn(ids...))
}

func (o queryOption) RegionIDEQ(regionID string) {
	o.query.Where(cityspliderareapolygon.RegionIDEQ(regionID))
}

func NewQuery(query *ent.CitySpliderAreaPolygonQuery) queryOption {
	return queryOption{query: query}
}
