package area

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/area"
	"github.com/cnartlu/area-service/pkg/utils"
)

// update 更新区域
func update(ctx context.Context, client *ent.Client, update *ent.AreaUpdateOne, data *ent.Area) (model *ent.Area, err error) {
	var (
		parentList       string = "0"
		id, _                   = update.Mutation().ID()
		idStr                   = strconv.FormatUint(id, 10)
		oldParentId, _          = update.Mutation().OldParentID(ctx)
		oldParentList, _        = update.Mutation().OldParentList(ctx)
		oldLevel, _             = update.Mutation().OldLevel(ctx)
	)
	data = FormatEntArea(data)
	data.Level = 1
	if oldParentId != data.ParentID {
		// 当前ID不能等于parentID的值
		if id == data.ParentID {
			return nil, ErrorIDCanEqualParentID
		}
		parentModel, err := client.Area.Query().
			Select(area.FieldParentList, area.FieldLevel).
			Where(area.IDEQ(data.ParentID)).
			First(ctx)
		if err != nil {
			return nil, err
		}
		parentList = parentModel.ParentList
		parentListArr := strings.Split(parentList, ",")
		// 不允许改变父级ID为其子集数据
		if utils.InArray(idStr, parentListArr) {
			return nil, ErrorParentIDSubordinateRecord
		}
		data.Level = parentModel.Level + 1
	}
	// 设置为事务执行
	err = db.WithTx(ctx, client, func(tx *ent.Tx) error {
		update.SetParentID(data.ParentID).
			SetParentList(parentList + "," + idStr).
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
		model, err = update.Save(ctx)
		if err != nil {
			return err
		}
		if oldParentId != data.ParentID {
			// 当前父级的数量+1
			if data.ParentID > 0 {
				_, err = client.Area.Update().
					Where(area.IDEQ(data.ParentID)).
					AddChildrenNumber(1).
					Save(ctx)
				if err != nil {
					return err
				}
			}
			// 原来的子集数量-1
			if oldParentId > 0 {
				_, err = client.Area.Update().
					Where(area.IDEQ(data.ParentID)).
					AddChildrenNumber((-1)).
					Save(ctx)
				if err != nil {
					return err
				}
			}
			// 切换子集数量
			sql := fmt.Sprintf(
				"UPDATE `%s` SET `%s`=`%s`-%d+%d, `%s`=REPLACE(`%s`, \"%s\", \"%s\") WHERE `%s` LIKE '%s,%%' AND `%s` <> %d",
				area.Table,
				area.FieldLevel,
				area.FieldLevel,
				oldLevel,
				model.Level,
				area.FieldParentList,
				area.FieldParentList,
				oldParentList,
				model.ParentList,
				area.FieldParentList,
				oldParentList,
				area.FieldID,
				model.ID,
			)
			_, err := update.ExecContext(ctx, sql)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return model, err
}

// Update 更新用户数据
func (r *Repository) Update(ctx context.Context, old *ent.Area, new *ent.Area) (*ent.Area, error) {
	return update(ctx, r.ent, r.ent.Area.UpdateOne(old), new)
}

// UpdateOneID 通过ID，更新数据
func (r *Repository) UpdateOneID(ctx context.Context, id uint64, data *ent.Area) (*ent.Area, error) {
	return update(ctx, r.ent, r.ent.Area.UpdateOneID(id), data)
}
