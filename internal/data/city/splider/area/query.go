package area

import (
	"strings"

	bizarea "github.com/cnartlu/area-service/internal/biz/city/splider/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/cityspliderarea"
)

var _ bizarea.Inquirer = new(queryOption)

type queryOption struct {
	query *ent.CitySpliderAreaQuery
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
