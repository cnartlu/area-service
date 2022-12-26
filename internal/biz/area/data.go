package area

import (
	"time"
)

const (
	// 区域地址长度为10
	AREA_REGIONID_LENGTH int = 10
	// 区域最小等级
	AREA_MIN_LEVEL int = 1
	// 区域最大等级
	AREA_MAX_LEVEL int = 4
)

// 区域数据实体
type Area struct {
	ID             int
	ParentID       int
	RegionID       string
	ParentList     string
	Title          string
	Pinyin         string
	Ucfirst        string
	Lng            float64
	Lat            float64
	Geohash        string
	GeoGs2         string
	GeoGs2ID       uint64
	GeoGs2Level    uint32
	GeoGs2Face     uint32
	CityCode       string
	ZipCode        string
	Level          int
	ChildrenNumber int
	DeletedAt      time.Time
	CreateAt       time.Time
	UpddateAt      time.Time
}

// 级联区域实体
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

// 查找参数
type FindListParam struct {
	// ParentID 父级ID
	ParentID int
	// RegionID 父级区域ID
	RegionID string
	// Level 区域级别
	Level int
	// Keyword 搜索关键字
	Keyword string
	// Order 排序
	Order string
}

type CreateParam struct {
	ParentID int
	RegionID string
	Title    string
	Lat      float64
	Lng      float64
	CityCode string
	ZipCode  string
}

type UpdateParam struct {
	ID int
	CreateParam
}
