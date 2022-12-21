package area

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cnartlu/area-service/errors"
	bizarea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/area"
)

var _ bizarea.AreaRepo = (*AreaRepo)(nil)

type AreaRepo struct {
	data *data.Data
}

func (r *AreaRepo) convertEntry(result *ent.Area) bizarea.Area {
	return bizarea.Area{
		ID:             result.ID,
		ParentID:       result.ParentID,
		RegionID:       result.RegionID,
		ParentList:     result.ParentList,
		Title:          result.Title,
		Pinyin:         result.Pinyin,
		Ucfirst:        result.Ucfirst,
		Lng:            result.Lng,
		Lat:            result.Lat,
		CityCode:       result.CityCode,
		ZipCode:        result.ZipCode,
		Level:          int(result.Level),
		ChildrenNumber: int(result.ChildrenNumber),
		DeletedAt:      time.Unix(int64(result.DeleteTime), 0),
		CreateAt:       time.Unix(int64(result.CreateTime), 0),
		UpddateAt:      time.Unix(int64(result.UpdateTime), 0),
	}
}

// Count 数量
func (r *AreaRepo) Count(ctx context.Context, options ...bizarea.Query) int {
	client := r.data.GetClient(ctx)
	query := client.Area.Query()
	if len(options) > 0 {
		q := NewQuery(query)
		for _, option := range options {
			option(q)
		}
	}
	i, _ := query.Count(ctx)
	return i
}

// FindList 查找数据
func (r *AreaRepo) FindList(ctx context.Context, options ...bizarea.Query) (list []*bizarea.Area, err error) {
	client := r.data.GetClient(ctx)
	query := client.Area.Query()
	if len(options) > 0 {
		q := NewQuery(query)
		for _, option := range options {
			option(q)
		}
	}
	results, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	list = make([]*bizarea.Area, len(results))
	for k, result := range results {
		entry := r.convertEntry(result)
		list[k] = &entry
	}
	return list, nil
}

// FindList 查找数据
func (r *AreaRepo) FindOne(ctx context.Context, options ...bizarea.Query) (*bizarea.Area, error) {
	client := r.data.GetClient(ctx)
	query := client.Area.Query()
	if len(options) > 0 {
		q := NewQuery(query)
		for _, option := range options {
			option(q)
		}
	}
	result, err := query.First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			err = errors.ErrorDataNotFound(err.Error())
		}
		return nil, err
	}
	data := r.convertEntry(result)
	return &data, nil
}

// FindList 查找数据
func (r *AreaRepo) Save(ctx context.Context, data *bizarea.Area) (*bizarea.Area, error) {
	client := r.data.GetClient(ctx)
	var (
		model    *ent.Area
		err      error
		isUpdate bool
	)
	if data.ID > 0 {
		isUpdate = true
		model, err = client.Area.Query().Where(area.IDEQ(data.ID)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			isUpdate = false
		}
	}
	if isUpdate {
		model, err = model.Update().
			SetParentID(data.ID).
			SetRegionID(data.RegionID).
			SetParentList(data.ParentList).
			SetTitle(data.Title).
			SetPinyin(data.Pinyin).
			SetUcfirst(data.Ucfirst).
			SetLevel(uint8(data.Level)).
			SetChildrenNumber(uint32(data.ChildrenNumber)).
			SetDeleteTime(uint64(data.DeletedAt.Unix())).
			SetUpdateTime(uint64(data.UpddateAt.Unix())).
			Save(ctx)
	} else {
		model, err = client.Area.Create().
			SetParentID(data.ID).
			SetRegionID(data.RegionID).
			SetParentList(data.ParentList).
			SetTitle(data.Title).
			SetPinyin(data.Pinyin).
			SetUcfirst(data.Ucfirst).
			SetLevel(uint8(data.Level)).
			SetChildrenNumber(uint32(data.ChildrenNumber)).
			SetDeleteTime(uint64(data.DeletedAt.Unix())).
			SetUpdateTime(uint64(data.UpddateAt.Unix())).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	entry := r.convertEntry(model)
	return &entry, err
}

// Remove 移除数据
func (r *AreaRepo) Remove(ctx context.Context, options ...bizarea.Query) error {
	client := r.data.GetClient(ctx)
	query := client.Area.Query()
	if len(options) > 0 {
		q := NewQuery(query)
		for _, option := range options {
			option(q)
		}
	}
	results, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, result := range results {
		if err1 := client.Area.DeleteOne(result).Exec(ctx); err1 != nil {
			err = err1
			break
		}
	}

	return err
}

func (r *AreaRepo) ReplaceParentListPrefix(ctx context.Context, oldPrefix, newPrefix string) (int, error) {
	client := r.data.GetClient(ctx)
	return client.Area.Update().
		Modify(func(u *sql.UpdateBuilder) {
			u.Set(area.FieldParentList, sql.Expr(fmt.Sprintf("REPLACE(`%s`, ?, ?)", area.FieldParentList), oldPrefix, newPrefix))
		}).
		Where(area.ParentListHasPrefix(oldPrefix), area.DeleteTimeGT(0)).
		Save(ctx)
}

func NewAreaRepo(d *data.Data) *AreaRepo {
	return &AreaRepo{
		data: d,
	}
}
