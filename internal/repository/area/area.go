package area

import (
	"fmt"
	"strings"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/area"
	"github.com/go-redis/redis/v8"
	"github.com/golang/geo/s2"
	"github.com/mmcloughlin/geohash"
	"github.com/mozillazg/go-pinyin"
)

var (
	// ErrorIDCanEqualParentID 主键ID不能与parentID一致
	ErrorIDCanEqualParentID = fmt.Errorf("primary key %s cannot be the same as %s", area.FieldID, area.FieldParentID)
	// ErrorParentIDSubordinateRecord 更改后的 parentID 不能是其从属记录
	ErrorParentIDSubordinateRecord = fmt.Errorf("the changed %s cannot be its subordinate record", area.FieldParentID)
)

type RepositoryManager interface {
	Querier
	Creator
	Updater
	Deleter
}

type Repository struct {
	ent *ent.Client
	rdb *redis.Client
}

var _ RepositoryManager = (*Repository)(nil)

// NewRepository 实例化存储数据
func NewRepository(ent *ent.Client, rdb *redis.Client) *Repository {
	return &Repository{
		ent: ent,
		rdb: rdb,
	}
}

func NewEntArea(regionID, title, pinyin string, lat, lng float64, zipCode, cityCode string, level uint8) *ent.Area {
	data := &ent.Area{
		RegionID: regionID,
		Title:    title,
		Pinyin:   pinyin,
		Lat:      lat,
		Lng:      lng,
		CityCode: cityCode,
		ZipCode:  zipCode,
		Level:    level,
	}
	return FormatEntArea(data)
}

// FormatEntArea 格式化area数据
func FormatEntArea(data *ent.Area) *ent.Area {
	geohashEncode := geohash.Encode(data.Lat, data.Lng)
	s2ll := s2.LatLngFromDegrees(data.Lat, data.Lng)
	s2Cell := s2.CellFromLatLng(s2ll)
	data.Geohash = geohashEncode
	data.GeoGs2 = s2Cell.ID().String()
	data.GeoGs2Face = uint32(s2Cell.Face())
	data.GeoGs2ID = s2Cell.ID().Pos()
	data.GeoGs2Level = uint32(s2Cell.Level())
	py := pinyin.NewArgs()
	data.Pinyin = strings.Join(pinyin.LazyPinyin(data.Title, py), " ")
	data.Ucfirst = ""
	if data.Pinyin != "" {
		data.Ucfirst = strings.ToUpper(data.Pinyin)[:1]
	}
	return data
}
