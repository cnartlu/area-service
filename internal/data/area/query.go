package area

import (
	bizarea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/area"
)

var _ bizarea.Inquirer = new(queryOption)

type queryOption struct {
	query *ent.AreaQuery
}

func (o queryOption) Offset(offset int) {
	o.query.Offset(offset)
}

func (o queryOption) Limit(limit int) {
	o.query.Limit(limit)
}

func (o queryOption) Order(orders ...bizarea.Sort) {
	for _, order := range orders {
		if order.Desc {
			o.query.Order(ent.Desc(order.Field))
		} else {
			o.query.Order(ent.Asc(order.Field))
		}
	}
}

func (o queryOption) IDEQ(id uint64) {
	o.query.Where(area.IDEQ(id))
}

func (o queryOption) IDIn(ids ...uint64) {
	o.query.Where(area.IDIn(ids...))
}

func (o queryOption) ParentIDEQ(pid uint64) {
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
