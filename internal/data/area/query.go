package area

import (
	"strings"

	bizArea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/area"
)

var _ bizArea.Inquirer = (*queryOption)(nil)

type queryOption struct {
	query *ent.AreaQuery
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
	o.query.Where(area.IDEQ(id))
}

func (o *queryOption) IDIn(ids ...uint64) {
	o.query.Where(area.IDIn(ids...))
}

func (o *queryOption) ParentIDEQ(pid uint64) {
	o.query.Where(area.ParentIDEQ(pid))
}

func (o *queryOption) RegionIDEQ(regionID string) {
	o.query.Where(area.RegionIDEQ(regionID))
}

func (o *queryOption) LevelEQ(level int) {
	o.query.Where(area.LevelEQ(uint8(level)))
}

func (o *queryOption) TitleContains(keyword string) {
	o.query.Where(area.TitleContains(keyword))
}

func NewQuery(query *ent.AreaQuery) queryOption {
	return queryOption{query: query}
}
