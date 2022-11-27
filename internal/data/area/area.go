package area

import (
	"context"
	"time"

	"github.com/cnartlu/area-service/api"
	bizArea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/area"
)

var _ bizArea.Manager = (*AreaRepo)(nil)

type AreaRepo struct {
	data *data.Data
}

func (r *AreaRepo) toAreaData(result *ent.Area) bizArea.Area {
	return bizArea.Area{
		ID:             result.ID,
		RegionID:       result.RegionID,
		Title:          result.Title,
		Pinyin:         result.Pinyin,
		Ucfirst:        result.Ucfirst,
		CityCode:       result.CityCode,
		ZipCode:        result.ZipCode,
		Level:          int(result.Level),
		ChildrenNumber: int(result.ChildrenNumber),
		CreateAt:       time.Unix(int64(result.CreateTime), 0),
		UpddateAt:      time.Unix(int64(result.UpdateTime), 0),
	}
}

// Count 数量
func (r *AreaRepo) Count(ctx context.Context, options ...bizArea.Option) int {
	client := r.data.GetClient(ctx)
	query := client.Area.Query()
	if len(options) > 0 {
		q := newOption(query)
		for _, option := range options {
			option(q)
		}
	}
	i, _ := query.Count(ctx)
	return i
}

// FindList 查找数据
func (r *AreaRepo) FindList(ctx context.Context, options ...bizArea.Option) (list []*bizArea.Area, err error) {
	client := r.data.GetClient(ctx)
	query := client.Area.Query()
	if len(options) > 0 {
		q := newOption(query)
		for _, option := range options {
			option(q)
		}
	}
	results, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		d := r.toAreaData(result)
		list = append(list, &d)
	}
	return list, nil
}

// FindList 查找数据
func (r *AreaRepo) FindOne(ctx context.Context, options ...bizArea.Option) (*bizArea.Area, error) {
	client := r.data.GetClient(ctx)
	query := client.Area.Query()
	if len(options) > 0 {
		q := newOption(query)
		for _, option := range options {
			option(q)
		}
	}
	result, err := query.First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			err = api.ErrorDataNotFound(err.Error())
		}
		return nil, err
	}
	data := r.toAreaData(result)
	return &data, nil
}

// FindList 查找数据
func (r *AreaRepo) Save(ctx context.Context, x *bizArea.Area) (*bizArea.Area, error) {
	client := r.data.GetClient(ctx)
	var (
		model    *ent.Area
		err      error
		isUpdate bool
	)
	if x.ID > 0 {
		isUpdate = true
		model, err = client.Area.Query().Where(area.IDEQ(x.ID)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			isUpdate = false
		}
	}
	if isUpdate {
		model, err = model.Update().Save(ctx)
	} else {
		model, err = client.Area.Create().Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	data := r.toAreaData(model)
	return &data, err
}

// Remove 移除数据
func (r *AreaRepo) Remove(ctx context.Context, options ...bizArea.Option) error {
	client := r.data.GetClient(ctx)
	query := client.Area.Query()
	if len(options) > 0 {
		q := newOption(query)
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

func NewAreaRepo(d *data.Data) *AreaRepo {
	return &AreaRepo{
		data: d,
	}
}
