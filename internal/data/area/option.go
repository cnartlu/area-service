package area

import (
	"strings"

	bizArea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/area"
)

var _ bizArea.OptionInterface = (*option)(nil)

type option struct {
	*ent.AreaQuery
}

func (o *option) Offset(offset int) {
	o.AreaQuery.Offset(offset)
}

func (o *option) Limit(limit int) {
	o.AreaQuery.Limit(limit)
}

func (o *option) Order(order string) {
	if order == "" {
		order = "-id"
	}
	orders := strings.Split(order, ",")
	for _, v := range orders {
		v := strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if strings.HasSuffix(v, "-") {
			o.AreaQuery.Order(ent.Desc(v[1:]))
		} else {
			o.AreaQuery.Order(ent.Asc(v))
		}
	}
}

func (o *option) IDEQ(id uint64) {
	o.AreaQuery.Where(area.IDEQ(id))
}

func (o *option) IDIn(ids ...uint64) {
	o.AreaQuery.Where(area.IDIn(ids...))
}

func (o *option) ParentIDEQ(pid uint64) {
	o.AreaQuery.Where(area.ParentIDEQ(pid))
}

func (o *option) RegionIDEQ(regionID string) {
	o.AreaQuery.Where(area.RegionIDEQ(regionID))
}

func (o *option) LevelEQ(level int) {
	o.AreaQuery.Where(area.LevelEQ(uint8(level)))
}

func (o *option) TitleContains(keyword string) {
	o.AreaQuery.Where(area.TitleContains(keyword))
}

func newOption(query *ent.AreaQuery) *option {
	return &option{query}
}
