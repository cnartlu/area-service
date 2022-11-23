package area

import (
	"strings"

	bizArea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/area"
)

var _ bizArea.OptionInterface = (*option)(nil)

type option struct {
	query *ent.AreaQuery
}

func (o *option) Offset(offset int) {
	o.query.Offset(offset)
}

func (o *option) Limit(limit int) {
	o.query.Limit(limit)
}

func (o *option) Order(order string) {
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

func (o *option) IDEQ(id uint64) {
	o.query.Where(area.IDEQ(id))
}

func (o *option) IDIn(ids ...uint64) {
	o.query.Where(area.IDIn(ids...))
}

func (o *option) ParentIDEQ(pid uint64) {
	o.query.Where(area.ParentIDEQ(pid))
}

func (o *option) RegionIDEQ(regionID string) {
	o.query.Where(area.RegionIDEQ(regionID))
}

func (o *option) LevelEQ(level int) {
	o.query.Where(area.LevelEQ(uint8(level)))
}

func (o *option) TitleContains(keyword string) {
	o.query.Where(area.TitleContains(keyword))
}

func newOption(query *ent.AreaQuery) *option {
	return &option{query: query}
}
