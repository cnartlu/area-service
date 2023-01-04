package polygon

import (
	"context"

	"ariga.io/entcache"
	"github.com/cnartlu/area-service/errors"
	bizareapolygon "github.com/cnartlu/area-service/internal/biz/city/splider/area/polygon"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/cityspliderareapolygon"
)

var _ bizareapolygon.PolygonRepo = (*PolygonRepo)(nil)

func convertPolygon(model *ent.CitySpliderAreaPolygon) bizareapolygon.Polygon {
	return bizareapolygon.Polygon{
		ID:       model.ID,
		RegionID: model.RegionID,
		Lng:      model.Lng,
		Lat:      model.Lat,
	}
}

type PolygonRepo struct {
	data *data.Data
}

func (r *PolygonRepo) Count(ctx context.Context, queries ...bizareapolygon.Query) int {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderAreaPolygon.Query()
	search := NewQuery(query)
	for _, fn := range queries {
		fn(search)
	}
	i, _ := query.Count(r.data.WithCacheContext(ctx, search.ttl))
	return i
}

func (r *PolygonRepo) FindList(ctx context.Context, queries ...bizareapolygon.Query) ([]*bizareapolygon.Polygon, error) {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderAreaPolygon.Query()
	search := NewQuery(query)
	for _, fn := range queries {
		fn(search)
	}
	models, err := query.All(r.data.WithCacheContext(ctx, search.ttl))
	if err != nil {
		return nil, err
	}
	items := []*bizareapolygon.Polygon{}
	for _, model := range models {
		model := model
		data := convertPolygon(model)
		items = append(items, &data)
	}
	return items, nil
}

func (r *PolygonRepo) FindOne(ctx context.Context, queries ...bizareapolygon.Query) (*bizareapolygon.Polygon, error) {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderAreaPolygon.Query()
	search := NewQuery(query)
	for _, fn := range queries {
		fn(search)
	}
	model, err := query.First(r.data.WithCacheContext(ctx, search.ttl))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrorDataNotFound(err.Error())
		}
		return nil, err
	}
	result := convertPolygon(model)
	return &result, nil
}

func (r *PolygonRepo) Save(ctx context.Context, data *bizareapolygon.Polygon) (*bizareapolygon.Polygon, error) {
	client := r.data.GetClient(ctx)
	var (
		model    *ent.CitySpliderAreaPolygon
		err      error
		isUpdate bool
	)
	if data.ID > 0 {
		isUpdate = true
		model, err = client.CitySpliderAreaPolygon.Query().
			Where(cityspliderareapolygon.IDEQ(data.ID)).
			First(entcache.Evict(ctx))
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			isUpdate = false
		}
	}

	if isUpdate {
		model, err = model.Update().
			SetRegionID(data.RegionID).
			SetLng(data.Lng).
			SetLat(data.Lat).
			Save(ctx)
	} else {
		model, err = client.CitySpliderAreaPolygon.Create().
			SetRegionID(data.RegionID).
			SetLng(data.Lng).
			SetLat(data.Lat).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	result := convertPolygon(model)
	return &result, nil
}

func (r *PolygonRepo) Remove(ctx context.Context, queries ...bizareapolygon.Query) error {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderAreaPolygon.Query()
	if len(queries) > 0 {
		search := NewQuery(query)
		for _, fn := range queries {
			fn(search)
		}
	}
	err := r.data.Transaction(entcache.Evict(ctx), func(ctx context.Context) error {
		results, err := query.All(entcache.Evict(ctx))
		if err != nil {
			return err
		}
		for _, result := range results {
			if err1 := client.CitySpliderAreaPolygon.DeleteOne(result).Exec(ctx); err1 != nil {
				err = err1
				break
			}
		}
		return err
	})
	return err
}

func NewPolygonRepo(
	d *data.Data,
) *PolygonRepo {
	return &PolygonRepo{
		data: d,
	}
}
