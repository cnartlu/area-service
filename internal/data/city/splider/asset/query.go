package asset

import (
	"strings"

	bizasset "github.com/cnartlu/area-service/internal/biz/city/splider/asset"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/cityspliderasset"
)

var _ bizasset.Inquirer = new(queryOption)

type queryOption struct {
	query *ent.CitySpliderAssetQuery
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
	o.query.Where(cityspliderasset.IDEQ(id))
}

func (o queryOption) IDIn(ids ...int) {
	o.query.Where(cityspliderasset.IDIn(ids...))
}

func (o queryOption) CitySpliderIDEQ(id int) {
	o.query.Where(cityspliderasset.CitySpliderIDEQ(id))
}

func (o queryOption) SourceIDEQ(sourceID uint64) {
	o.query.Where(cityspliderasset.SourceIDEQ(sourceID))
}

func (o queryOption) SourceIDIn(sourceIDs []uint64) {
	o.query.Where(cityspliderasset.SourceIDIn(sourceIDs...))
}

func (o queryOption) StatusEQ(status bizasset.Status) {
	o.query.Where(cityspliderasset.StatusEQ(uint8(status)))
}

func NewQuery(query *ent.CitySpliderAssetQuery) queryOption {
	return queryOption{query: query}
}
