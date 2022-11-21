package asset

import (
	"context"

	bizAsset "github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/ent"

	"github.com/go-redis/redis/v8"
)

var _ bizAsset.ManageRepo = (*AssetRepo)(nil)

type AssetRepo struct {
	ent *ent.Client
	rds *redis.Client
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
	for _, v := range models {
		v := v
		items = append(items, &bizAsset.Asset{
			ID: v.ID,
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
		return nil, err
	}
	p := bizAsset.Asset{
		ID: model.ID,
	}
	return &p, nil
}

func (r *AssetRepo) Create(ctx context.Context, data *ent.AreaReleaseAsset) (*ent.AreaReleaseAsset, error) {
	r.ent.Tx(ctx)
	model, err := r.ent.AreaReleaseAsset.Create().
		SetAreaReleaseID(data.AreaReleaseID).
		SetAssetID(data.AssetID).
		SetAssetLabel(data.AssetLabel).
		SetAssetName(data.AssetName).
		SetAssetState(data.AssetState).
		SetFilePath(data.FilePath).
		SetFileSize(data.FileSize).
		SetDownloadURL(data.DownloadURL).
		SetStatus(data.Status).
		Save(ctx)
	return model, err
}

func (r *AssetRepo) Update(ctx context.Context, data *ent.AreaReleaseAsset) (*ent.AreaReleaseAsset, error) {
	model, err := r.ent.AreaReleaseAsset.UpdateOne(data).
		SetAreaReleaseID(data.AreaReleaseID).
		SetAssetID(data.AssetID).
		AddAreaReleaseID(-1).
		SetAssetLabel(data.AssetLabel).
		SetAssetName(data.AssetName).
		SetAssetState(data.AssetState).
		SetFilePath(data.FilePath).
		SetFileSize(data.FileSize).
		SetDownloadURL(data.DownloadURL).
		SetStatus(data.Status).
		Save(ctx)
	return model, err
}

func (r *AssetRepo) Delete(ctx context.Context, id uint64) error {
	return r.ent.AreaReleaseAsset.DeleteOneID(id).Exec(ctx)
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
