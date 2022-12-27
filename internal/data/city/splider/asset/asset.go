package asset

import (
	"context"
	"time"

	"ariga.io/entcache"
	"github.com/cnartlu/area-service/errors"
	bizasset "github.com/cnartlu/area-service/internal/biz/city/splider/asset"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/cityspliderasset"
)

var _ bizasset.AssetRepo = (*AssetRepo)(nil)

func convertAsset(model *ent.CitySpliderAsset) bizasset.Asset {
	return bizasset.Asset{
		ID:            model.ID,
		CitySpliderID: model.CitySpliderID,
		SourceID:      model.SourceID,
		FileTitle:     model.FileTitle,
		FilePath:      model.FilePath,
		FileSize:      model.FileSize,
		Status:        bizasset.Status(model.Status),
		CreatedAt:     time.Unix(int64(model.CreateTime), 0),
		UpdatedAt:     time.Unix(int64(model.UpdateTime), 0),
	}
}

type AssetRepo struct {
	data *data.Data
}

func (r *AssetRepo) Count(ctx context.Context, queries ...bizasset.Query) int {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderAsset.Query()
	search := NewQuery(query)
	for _, fn := range queries {
		fn(search)
	}
	i, _ := query.Count(r.data.WithCacheContext(ctx, search.ttl))
	return i
}

func (r *AssetRepo) FindList(ctx context.Context, queries ...bizasset.Query) ([]*bizasset.Asset, error) {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderAsset.Query()
	search := NewQuery(query)
	for _, fn := range queries {
		fn(search)
	}
	models, err := query.All(r.data.WithCacheContext(ctx, search.ttl))
	if err != nil {
		return nil, err
	}
	items := []*bizasset.Asset{}
	for _, model := range models {
		model := model
		data := convertAsset(model)
		items = append(items, &data)
	}
	return items, nil
}

func (r *AssetRepo) FindOne(ctx context.Context, queries ...bizasset.Query) (*bizasset.Asset, error) {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderAsset.Query()
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
	result := convertAsset(model)
	return &result, nil
}

func (r *AssetRepo) Save(ctx context.Context, data *bizasset.Asset) (*bizasset.Asset, error) {
	client := r.data.GetClient(ctx)
	var (
		model    *ent.CitySpliderAsset
		err      error
		isUpdate bool
	)
	if data.ID > 0 {
		isUpdate = true
		model, err = client.CitySpliderAsset.Query().
			Where(cityspliderasset.IDEQ(data.ID)).
			First(entcache.Evict(ctx))
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			isUpdate = false
		}
	}
	if data.CreatedAt.IsZero() {
		data.CreatedAt = time.Now()
	}
	if data.UpdatedAt.IsZero() {
		data.UpdatedAt = time.Now()
	}
	if isUpdate {
		model, err = model.Update().
			SetCitySpliderID(data.CitySpliderID).
			SetSourceID(data.SourceID).
			SetFileTitle(data.FileTitle).
			SetFilePath(data.FilePath).
			SetFileSize(data.FileSize).
			SetStatus(uint8(data.Status)).
			SetUpdateTime(uint64(data.UpdatedAt.Unix())).
			Save(ctx)
	} else {
		model, err = client.CitySpliderAsset.Create().
			SetCitySpliderID(data.CitySpliderID).
			SetSourceID(data.SourceID).
			SetFileTitle(data.FileTitle).
			SetFilePath(data.FilePath).
			SetFileSize(data.FileSize).
			SetStatus(uint8(data.Status)).
			SetCreateTime(uint64(data.CreatedAt.Unix())).
			SetUpdateTime(uint64(data.UpdatedAt.Unix())).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	result := convertAsset(model)
	return &result, nil
}

func (r *AssetRepo) Remove(ctx context.Context, queries ...bizasset.Query) error {
	client := r.data.GetClient(ctx)
	query := client.CitySpliderAsset.Query()
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
			if err1 := client.CitySpliderAsset.DeleteOne(result).Exec(ctx); err1 != nil {
				err = err1
				break
			}
		}
		return err
	})
	return err
}

func NewAssetRepo(
	d *data.Data,
) *AssetRepo {
	return &AssetRepo{
		data: d,
	}
}
