package area

import (
	bizArea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/area"
)

type option struct {
	*ent.AreaQuery
}

var _ bizArea.OptionInterface = (*option)(nil)

// Offset 偏移量
func (o *option) Offset(offset int) {
	o.AreaQuery.Offset(offset)
}

// Limit 限制查询数量
func (o *option) Limit(limit int) {
	o.AreaQuery.Limit(limit)

}

func (o *option) WithID(id uint64) {
	o.AreaQuery.Where(area.IDEQ(id))
}

func (o *option) WithParentID(pid uint64) {
	o.AreaQuery.Where(area.ParentIDEQ(pid))
}

func (o *option) WithRegionID(regionID string) {
	o.AreaQuery.Where(area.RegionIDEQ(regionID))
}

func (o *option) WithLevel(level int) {
	o.AreaQuery.Where(area.LevelEQ(uint8(level)))
}

func (o *option) WithKeywordContains(keyword string) {
	o.AreaQuery.Where(area.TitleContains(keyword))
}

func newOption(query *ent.AreaQuery) *option {
	return &option{query}
}
