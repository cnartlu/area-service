package asset

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/go-redis/redis/v8"
)

type AssetRepository interface{}

type AssetRepo struct {
	ent *ent.Client
	rds *redis.Client
}

func (r *AssetRepo) FindList(ctx context.Context, options ...option) ([]*ent.AreaReleaseAsset, error) {
	query := r.ent.AreaReleaseAsset.Query()
	for _, option := range options {
		option(query)
	}
	return query.All(ctx)
}

func (r *AssetRepo) FindOne(ctx context.Context, options ...option) (*ent.AreaReleaseAsset, error) {
	query := r.ent.AreaReleaseAsset.Query()
	for _, option := range options {
		option(query)
	}
	return query.First(ctx)
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
