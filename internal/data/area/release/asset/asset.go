package asset

import (
	"context"

	"github.com/cnartlu/area-service/api"
	bizAsset "github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/areareleaseasset"

	"github.com/go-redis/redis/v8"
)

var _ bizAsset.ManageRepo = (*AssetRepo)(nil)

type AssetRepo struct {
	ent *ent.Client
	rds *redis.Client
}

func (r *AssetRepo) Count(ctx context.Context, options ...bizAsset.Option) int {
	query := r.ent.AreaReleaseAsset.Query()
	o := newOption(query)
	for _, option := range options {
		option := option
		option(o)
	}
	i, _ := query.Count(ctx)
	return i
}

func (r *AssetRepo) FindList(ctx context.Context, options ...bizAsset.Option) ([]*bizAsset.Asset, error) {
	query := r.ent.AreaReleaseAsset.Query()
	o := newOption(query)
	for _, option := range options {
		option(o)
	}
	models, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	items := []*bizAsset.Asset{}
	for _, model := range models {
		model := model
		items = append(items, &bizAsset.Asset{
			ID:            model.ID,
			CreateTime:    model.CreateTime,
			UpdateTime:    model.UpdateTime,
			AreaReleaseID: model.AreaReleaseID,
			AssetID:       model.AssetID,
			AssetName:     model.AssetName,
			AssetLabel:    model.AssetLabel,
			AssetState:    model.AssetState,
			FilePath:      model.FilePath,
			FileSize:      model.FileSize,
			DownloadURL:   model.DownloadURL,
			// Status: model.Status,
		})
	}
	return items, nil
}

func (r *AssetRepo) FindOne(ctx context.Context, options ...bizAsset.Option) (*bizAsset.Asset, error) {
	query := r.ent.AreaReleaseAsset.Query()
	o := newOption(query)
	for _, option := range options {
		option(o)
	}
	model, err := query.First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, api.ErrorDataNotFound(err.Error())
		}
		return nil, err
	}
	p := &bizAsset.Asset{
		ID:            model.ID,
		CreateTime:    model.CreateTime,
		UpdateTime:    model.UpdateTime,
		AreaReleaseID: model.AreaReleaseID,
		AssetID:       model.AssetID,
		AssetName:     model.AssetName,
		AssetLabel:    model.AssetLabel,
		AssetState:    model.AssetState,
		FilePath:      model.FilePath,
		FileSize:      model.FileSize,
		DownloadURL:   model.DownloadURL,
		// Status: model.Status,
	}
	return p, nil
}

func (r *AssetRepo) Save(ctx context.Context, data *bizAsset.Asset) (*bizAsset.Asset, error) {
	var (
		model    *ent.AreaReleaseAsset
		err      error
		isUpdate bool
	)
	if data.ID > 0 {
		isUpdate = true
		model, err = r.ent.AreaReleaseAsset.Query().Where(areareleaseasset.IDEQ(data.ID)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			isUpdate = false
		}
	}
	if isUpdate {
		model, err = model.Update().
			SetAreaReleaseID(data.AreaReleaseID).
			SetAssetID(data.AssetID).
			SetAssetLabel(data.AssetLabel).
			SetAssetName(data.AssetName).
			SetAssetState(data.AssetState).
			SetFilePath(data.FilePath).
			SetFileSize(data.FileSize).
			SetDownloadURL(data.DownloadURL).
			// SetStatus(data.Status).
			Save(ctx)
	} else {
		model, err = r.ent.AreaReleaseAsset.Create().
			SetAreaReleaseID(data.AreaReleaseID).
			SetAssetID(data.AssetID).
			SetAssetLabel(data.AssetLabel).
			SetAssetName(data.AssetName).
			SetAssetState(data.AssetState).
			SetFilePath(data.FilePath).
			SetFileSize(data.FileSize).
			SetDownloadURL(data.DownloadURL).
			// SetStatus(data.Status).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}

	return &bizAsset.Asset{
		ID:            model.ID,
		CreateTime:    model.CreateTime,
		UpdateTime:    model.UpdateTime,
		AreaReleaseID: model.AreaReleaseID,
		AssetID:       model.AssetID,
		AssetName:     model.AssetName,
		AssetLabel:    model.AssetLabel,
		AssetState:    model.AssetState,
		FilePath:      model.FilePath,
		FileSize:      model.FileSize,
		DownloadURL:   model.DownloadURL,
		// Status: model.Status,
	}, nil
}

func (r *AssetRepo) Remove(ctx context.Context, options ...bizAsset.Option) error {
	query := r.ent.AreaReleaseAsset.Query()
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
		r.ent.AreaReleaseAsset.DeleteOne(result).Exec(ctx)
	}
	return err
}

func NewAssetRepo(
	ent *ent.Client,
	rds *redis.Client,
) *AssetRepo {
	return &AssetRepo{
		ent: ent,
		rds: rds,
	}
}
