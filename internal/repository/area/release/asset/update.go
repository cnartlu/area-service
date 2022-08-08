package asset

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

type Updater interface {
	// Update 更新资源记录
	Update(ctx context.Context, model *ent.AreaReleaseAsset) (*ent.AreaReleaseAsset, error)
	// UpdateOrCreate 更新或创建数据
	UpdateOrCreate(ctx context.Context, model *ent.AreaReleaseAsset) (*ent.AreaReleaseAsset, error)
}

// Update 创建记录
func (r *Repository) Update(ctx context.Context, model *ent.AreaReleaseAsset) (*ent.AreaReleaseAsset, error) {
	return model.Update().Save(ctx)
}

func (r *Repository) UpdateOrCreate(ctx context.Context, data *ent.AreaReleaseAsset) (*ent.AreaReleaseAsset, error) {
	var mm interface {
		Save(context.Context) (*ent.AreaReleaseAsset, error)
	}
	model, err := r.FindOneByID(ctx, data.ID)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		mm = r.ent.AreaReleaseAsset.Create().
			SetAreaReleaseID(data.AreaReleaseID).
			SetAssetID(data.AssetID).
			SetAssetName(data.AssetName).
			SetAssetLabel(data.AssetLabel).
			SetAssetState(data.AssetState).
			SetFilePath(data.FilePath).
			SetFileSize(data.FileSize).
			SetDownloadURL(data.DownloadURL).
			SetStatus(data.Status)
	} else {
		mm = model.Update().
			SetAreaReleaseID(data.AreaReleaseID).
			SetAssetID(data.AssetID).
			SetAssetName(data.AssetName).
			SetAssetLabel(data.AssetLabel).
			SetAssetState(data.AssetState).
			SetFilePath(data.FilePath).
			SetFileSize(data.FileSize).
			SetDownloadURL(data.DownloadURL).
			SetStatus(data.Status)
	}
	return mm.Save(ctx)
}
