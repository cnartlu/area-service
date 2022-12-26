package area

import (
	"context"

	"github.com/cnartlu/area-service/errors"
	bizarea "github.com/cnartlu/area-service/internal/biz/city/splider/area"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/cityspliderarea"
)

var _ bizarea.AreaRepo = (*AreaRepo)(nil)

func convertArea(model *ent.CitySpliderArea) bizarea.Area {
	return bizarea.Area{
		ID:         model.ID,
		ParentID:   model.ParentID,
		RegionID:   model.RegionID,
		ParentList: model.ParentList,
		Title:      model.Title,
		Level:      int(model.Level),
	}
}

type AreaRepo struct {
	data *data.Data
}

func (r *AreaRepo) Count(ctx context.Context, options ...bizarea.Query) int {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderArea.Query()
	o := NewQuery(query)
	for _, option := range options {
		option := option
		option(o)
	}
	i, _ := query.Count(ctx)
	return i
}

func (r *AreaRepo) FindList(ctx context.Context, options ...bizarea.Query) ([]*bizarea.Area, error) {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderArea.Query()
	o := NewQuery(query)
	for _, option := range options {
		option(o)
	}
	models, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	items := []*bizarea.Area{}
	for _, model := range models {
		model := model
		data := convertArea(model)
		items = append(items, &data)
	}
	return items, nil
}

func (r *AreaRepo) FindOne(ctx context.Context, options ...bizarea.Query) (*bizarea.Area, error) {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderArea.Query()
	o := NewQuery(query)
	for _, option := range options {
		option(o)
	}
	model, err := query.First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrorDataNotFound(err.Error())
		}
		return nil, err
	}
	result := convertArea(model)
	return &result, nil
}

func (r *AreaRepo) Save(ctx context.Context, data *bizarea.Area) (*bizarea.Area, error) {
	client := r.data.GetClient(ctx)
	var (
		model    *ent.CitySpliderArea
		err      error
		isUpdate bool
	)
	if data.ID > 0 {
		isUpdate = true
		model, err = client.CitySpliderArea.Query().Where(cityspliderarea.IDEQ(data.ID)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			isUpdate = false
		}
	}
	if isUpdate {
		model, err = model.Update().
			SetParentID(data.ParentID).
			SetRegionID(data.RegionID).
			SetParentList(data.ParentList).
			SetTitle(data.Title).
			SetLevel(uint8(data.Level)).
			Save(ctx)
	} else {
		model, err = client.CitySpliderArea.Create().
			SetParentID(data.ParentID).
			SetRegionID(data.RegionID).
			SetParentList(data.ParentList).
			SetTitle(data.Title).
			SetLevel(uint8(data.Level)).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	result := convertArea(model)
	return &result, nil
}

func (r *AreaRepo) Remove(ctx context.Context, options ...bizarea.Query) error {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderArea.Query()
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
		if err1 := client.CitySpliderArea.DeleteOne(result).Exec(ctx); err1 != nil {
			err = err1
			break
		}
	}
	return err
}

func NewAreaRepo(
	d *data.Data,
) *AreaRepo {
	return &AreaRepo{
		data: d,
	}
}
