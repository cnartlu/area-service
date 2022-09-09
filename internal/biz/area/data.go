package area

import "time"

type Area struct {
	ID        uint64
	RegionID  string
	Title     string
	Pinyin    string
	Ucfirst   string
	CityCode  string
	ZipCode   string
	Level     int
	CreateAt  time.Time
	UpddateAt time.Time
}

type CascadeArea struct {
	RegionID string
}
