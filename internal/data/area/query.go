package area

import (
	"strings"

	bizarea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/area"
)

var _ bizarea.Inquirer = new(queryOption)

type queryOption struct {
	query *ent.AreaQuery
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
	o.query.Where(area.IDEQ(id))
}

func (o queryOption) IDIn(ids ...int) {
	o.query.Where(area.IDIn(ids...))
}

func (o queryOption) ParentIDEQ(pid int) {
	o.query.Where(area.ParentIDEQ(pid))
}

func (o queryOption) RegionIDEQ(regionID string) {
	o.query.Where(area.RegionIDEQ(regionID))
}

func (o queryOption) LevelEQ(level int) {
	o.query.Where(area.LevelEQ(uint8(level)))
}

func (o queryOption) TitleContains(keyword string) {
	o.query.Where(area.TitleContains(keyword))
}

func NewQuery(query *ent.AreaQuery) queryOption {
	return queryOption{query: query}
}
