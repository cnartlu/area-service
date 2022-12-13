package area

import (
	pkgsort "github.com/cnartlu/area-service/pkg/data/sort"
)

type Sort = pkgsort.Sort

type SortField int

const (
	SortFieldDefault  SortField = 0
	SortFieldRegionID SortField = 1
)

type Inquirer interface {
	Limit(limit int)
	Offset(offset int)
	// Order(orders ...pkgsort.Sort)
	IDEQ(id uint64)
	IDIn(ids ...uint64)
	ParentIDEQ(parentID uint64)
	RegionIDEQ(regionID string)
	LevelEQ(level int)
	TitleContains(keyword string)
}

type Query func(Inquirer)

func Offset(offset int) Query {
	return func(r Inquirer) {
		r.Offset(offset)
	}
}

func Limit(limit int) Query {
	return func(r Inquirer) {
		r.Limit(limit)
	}
}

func Order(orders ...pkgsort.Sort) Query {
	return func(r Inquirer) {
		// r.Order(orders...)
	}
}

func IDEQ(id uint64) Query {
	return func(r Inquirer) {
		r.IDEQ(id)
	}
}

func IDIn(ids ...uint64) Query {
	return func(r Inquirer) {
		r.IDIn(ids...)
	}
}

func ParentIDEQ(parentID uint64) Query {
	return func(r Inquirer) {
		r.ParentIDEQ(parentID)
	}
}

func RegionIDEQ(regionID string) Query {
	return func(r Inquirer) {
		r.RegionIDEQ(regionID)
	}
}

func LevelEQ(level int) Query {
	return func(r Inquirer) {
		r.LevelEQ(level)
	}
}

func TitleContains(keyword string) Query {
	return func(r Inquirer) {
		r.TitleContains(keyword)
	}
}
