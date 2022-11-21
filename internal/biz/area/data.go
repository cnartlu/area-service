package area

import "time"

type Area struct {
	ID             uint64
	RegionID       string
	Title          string
	Pinyin         string
	Ucfirst        string
	CityCode       string
	ZipCode        string
	Level          int
	ChildrenNumber int
	CreateAt       time.Time
	UpddateAt      time.Time
}

type CascadeArea struct {
	ID             uint64
	RegionID       string
	Title          string
	Pinyin         string
	Ucfirst        string
	Level          int
	ChildrenNumber int
	Items          []*CascadeArea
}
