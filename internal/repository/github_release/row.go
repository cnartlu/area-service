package github_release

import (
	"context"
	"strconv"
	"strings"

	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/area"
	"github.com/cnartlu/area-service/internal/component/ent/areapolygon"
	"github.com/golang/geo/s2"
	"github.com/mmcloughlin/geohash"
)

type LevelRow struct {
	ID      string
	Pid     string
	Level   int
	Name    string
	Pinyin  string
	ExtId   string
	ExtName string
}

// Ucfirst 首字母
func (r LevelRow) Ucfirst() string {
	ucfirst := ""
	if len(r.Pinyin) > 0 {
		ucfirst = strings.ToUpper(r.Pinyin[0:1])
	}
	return ucfirst
}

// Writer 写入数据
func (r LevelRow) Writer(ctx context.Context, db *db.DB) error {
	tx := db.DB()
	ok, err := tx.Area.Query().
		Where(area.RegionIDEQ(r.ID), area.LevelEQ(uint8(r.Level))).
		Exist(ctx)
	if err != nil {
		return err
	}
	if ok {
		_, err = tx.Area.Update().
			Where(area.RegionIDEQ(r.ID), area.LevelEQ(uint8(r.Level))).
			SetRegionID(r.ID).
			SetTitle(r.Name).
			SetPinyin(r.Pinyin).
			SetUcfirst(r.Ucfirst()).
			SetLevel(uint8(r.Level)).
			Save(ctx)
	} else {
		_, err = tx.Area.Create().
			SetRegionID(r.ID).
			SetTitle(r.Name).
			SetPinyin(r.Pinyin).
			SetUcfirst(r.Ucfirst()).
			SetLevel(uint8(r.Level)).
			Save(ctx)
	}
	if err != nil {
		return err
	}
	return nil
}

type geo struct {
	Lat float64
	Lng float64
}

// GetGeoHash 获取geohash值
func (g *geo) GetGeoHash() string {
	return geohash.Encode(g.Lat, g.Lng)
}

// GetS2Cell 获取s2单元格
func (g *geo) GetS2Cell() s2.Cell {
	s2ll := s2.LatLngFromDegrees(g.Lat, g.Lng)
	return s2.CellFromLatLng(s2ll)
}

type GeoRow struct {
	ID       string
	Pid      string
	Level    int
	Name     string
	ExtPath  string
	Geo      *geo
	Polygons []*geo
}

// Writer 写入数据
func (r GeoRow) Writer(ctx context.Context, db *db.DB) error {
	tx := db.DB()
	areaId, err := tx.Area.Query().
		Where(area.RegionIDEQ(r.ID)).
		Order(ent.Asc(area.FieldLevel)).
		FirstID(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return err
	}
	// 删除全部的
	_, err = tx.AreaPolygon.Delete().
		Where(areapolygon.RegionIDEQ(r.ID)).
		Exec(ctx)
	if err != nil {
		return err
	}
	// 父级的lng和lat
	if areaId > 0 {
		geo := *r.Geo
		s2Cell := geo.GetS2Cell()
		_, err := tx.Area.Update().
			Where(area.RegionIDEQ(r.ID)).
			SetLat(geo.Lat).
			SetLng(geo.Lng).
			SetGeohash(geo.GetGeoHash()).
			SetGeoGs2(s2Cell.ID().String()).
			SetGeoGs2Face(uint32(s2Cell.Face())).
			SetGeoGs2ID(s2Cell.ID().Pos()).
			SetGeoGs2Level(uint32(s2Cell.Level())).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	for _, polygon := range r.Polygons {
		if polygon == nil {
			continue
		}
		if err := r.writerPolygon(ctx, tx.AreaPolygon, areaId, r.ID, polygon); err != nil {
			return err
		}
	}
	return nil
}

func (r GeoRow) writerPolygon(ctx context.Context, client *ent.AreaPolygonClient, areaID uint64, regionID string, g *geo) error {
	ok, err := client.Query().
		Where(areapolygon.RegionIDEQ(regionID)).
		Where(areapolygon.LatEQ(g.Lat), areapolygon.LngEQ(g.Lng)).
		Exist(ctx)
	if err != nil {
		return err
	}
	if !ok {
		return client.Create().
			SetAreaID(areaID).
			SetRegionID(regionID).
			SetLat(g.Lat).
			SetLng(g.Lng).
			Exec(ctx)
	}
	return nil
}

// toGeoWithLatLng 将str转为geo结构
func toGeo(geoStr string) *geo {
	var geodata = geo{}
	geoStrings := strings.SplitN(geoStr, ",", 2)
	if len(geoStrings) == 2 {
		geodata.Lng, _ = strconv.ParseFloat(geoStrings[0], 64)
		geodata.Lat, _ = strconv.ParseFloat(geoStrings[1], 64)
	}
	return &geodata
}

// toPolygons 字符串转区域
func toPolygons(str string) []*geo {
	polygons := []*geo{}
	strs := strings.SplitN(str, " ", -1)
	for _, geoStr := range strs {
		geoStr = strings.TrimSpace(geoStr)
		if geoStr == "" {
			continue
		}
		if g := toGeo(geoStr); g != nil {
			polygons = append(polygons, g)
		}
	}
	return polygons
}
