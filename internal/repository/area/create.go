package area

import (
	"context"
	"strconv"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/area"
)

type Creator interface {
	// Create 创建数据
	Create(ctx context.Context, area *ent.Area) (*ent.Area, error)
}

// Create 创建数据
func (r *Repository) Create(ctx context.Context, data *ent.Area) (*ent.Area, error) {
	data = FormatEntArea(data)
	client := r.ent
	var parentList string = "0"
	data.Level = 1
	if data.ParentID > 0 {
		parentModel, err := client.Area.Query().
			Select(area.FieldParentList, area.FieldLevel).
			Where(area.IDEQ(data.ParentID)).
			First(ctx)
		if err != nil {
			return nil, err
		}
		parentList = parentModel.ParentList
		data.Level = parentModel.Level
	}
	create := client.Area.Create().
		SetParentID(data.ParentID).
		SetRegionID(data.RegionID).
		SetTitle(data.Title).
		SetPinyin(data.Pinyin).
		SetUcfirst(data.Ucfirst).
		SetLng(data.Lng).
		SetLat(data.Lat).
		SetGeohash(data.Geohash).
		SetGeoGs2(data.GeoGs2).
		SetGeoGs2Face(data.GeoGs2Face).
		SetGeoGs2ID(data.GeoGs2ID).
		SetGeoGs2Level(data.GeoGs2Level).
		SetCityCode(data.CityCode).
		SetZipCode(data.ZipCode).
		SetLevel(data.Level)
	model, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	model, err = model.Update().
		SetParentList(parentList + "," + strconv.FormatUint(model.ID, 10)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	if model.ParentID > 0 {
		_, err = client.Area.Update().
			Where(area.IDEQ(model.ParentID)).
			AddChildrenNumber(1).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	return model, nil
}

// CreateOrUpdate 创建或更新数据
func (r *Repository) CreateOrUpdate(ctx context.Context, data *ent.Area) (*ent.Area, error) {
	var (
		model *ent.Area
		err   error
	)
	if data.ID > 0 {
		model, err = r.FindOneById(ctx, data.ID, []string{})
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			model = nil
		}
	}
	if model == nil {
		model, err = r.FindOneByRegionID(ctx, data.RegionID, data.Level)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			model = nil
		}
	}
	if model == nil {
		model, err = r.Create(ctx, data)
	} else {
		model, err = r.Update(ctx, model, data)
	}
	if err != nil {
		return nil, err
	}
	return model, nil
}
