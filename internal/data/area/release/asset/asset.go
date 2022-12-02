package asset

import (
	"context"

	"github.com/cnartlu/area-service/api"
	bizasset "github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/areareleaseasset"
)

var _ bizasset.ManageRepo = (*AssetRepo)(nil)

type AssetRepo struct {
	data *data.Data
}

func (r *AssetRepo) Count(ctx context.Context, options ...bizasset.Option) int {
	client := r.data.GetClient(ctx)
	query := client.AreaReleaseAsset.Query()
	o := newOption(query)
	for _, option := range options {
		option := option
		option(o)
	}
	i, _ := query.Count(ctx)
	return i
}

func (r *AssetRepo) FindList(ctx context.Context, options ...bizasset.Option) ([]*bizasset.Asset, error) {
	client := r.data.GetClient(ctx)
	query := client.AreaReleaseAsset.Query()
	o := newOption(query)
	for _, option := range options {
		option(o)
	}
	models, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	items := []*bizasset.Asset{}
	for _, model := range models {
		model := model
		items = append(items, &bizasset.Asset{
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
			Status:        bizasset.Status(model.Status),
		})
	}
	return items, nil
}

func (r *AssetRepo) FindOne(ctx context.Context, options ...bizasset.Option) (*bizasset.Asset, error) {
	client := r.data.GetClient(ctx)
	query := client.AreaReleaseAsset.Query()
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
	p := &bizasset.Asset{
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
		Status:        bizasset.Status(model.Status),
	}
	return p, nil
}

func (r *AssetRepo) Save(ctx context.Context, data *bizasset.Asset) (*bizasset.Asset, error) {
	client := r.data.GetClient(ctx)
	var (
		model    *ent.AreaReleaseAsset
		err      error
		isUpdate bool
	)
	if data.ID > 0 {
		isUpdate = true
		model, err = client.AreaReleaseAsset.Query().Where(areareleaseasset.IDEQ(data.ID)).First(ctx)
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
			SetStatus(uint8(data.Status)).
			Save(ctx)
	} else {
		model, err = client.AreaReleaseAsset.Create().
			SetAreaReleaseID(data.AreaReleaseID).
			SetAssetID(data.AssetID).
			SetAssetLabel(data.AssetLabel).
			SetAssetName(data.AssetName).
			SetAssetState(data.AssetState).
			SetFilePath(data.FilePath).
			SetFileSize(data.FileSize).
			SetDownloadURL(data.DownloadURL).
			SetStatus(uint8(data.Status)).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}

	return &bizasset.Asset{
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
		Status:        bizasset.Status(model.Status),
	}, nil
}

func (r *AssetRepo) Remove(ctx context.Context, options ...bizasset.Option) error {
	client := r.data.GetClient(ctx)
	query := client.AreaReleaseAsset.Query()
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
		if err1 := client.AreaReleaseAsset.DeleteOne(result).Exec(ctx); err1 != nil {
			err = err1
			break
		}
	}
	return err
}

func NewAssetRepo(
	d *data.Data,
) *AssetRepo {
	return &AssetRepo{
		data: d,
	}
}
