package release

import (
	"context"
	"sync"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/areareleaseasset"
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Importer interface {
	Import(ctx context.Context, areaRelease *ent.AreaRelease) error
}

// Import 导入资源文件
func (s *Service) Import(ctx context.Context, areaRelease *ent.AreaRelease) error {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	releaseAssets, err := s.assetRepo.FindList(ctx, asset.NewFindListParam().SetAreaReleaseID(areaRelease.ID))
	if err != nil {
		return errors.Wrap(err, "area.release.asset repository")
	}
	var waitGroup = &sync.WaitGroup{}
	for _, releaseAsset := range releaseAssets {
		waitGroup.Add(1)
		go func(areaReleaseAsset *ent.AreaReleaseAsset) {
			defer waitGroup.Done()
			err = s.importerService.Import(ctx, areaReleaseAsset, nil)
			if err != nil {
				s.logger.Error(
					"导入资源文件失败",
					zap.Uint64("areaRelease", areaRelease.ID),
					zap.Uint64("areaReleaseAssetID", areaReleaseAsset.ID),
					zap.Error(err),
				)
				return
			}
			areaReleaseAsset.Status = areareleaseasset.StatusFinishedLoaded
			_, err = s.assetRepo.Update(ctx, areaReleaseAsset)
			if err != nil {
				s.logger.Error(
					"更新导入状态失败",
					zap.Uint64("areaRelease", areaRelease.ID),
					zap.Uint64("areaReleaseAssetID", areaReleaseAsset.ID),
					zap.Error(err),
				)
			}
		}(releaseAsset)
	}
	waitGroup.Wait()
	// 判断是否完成完成
	return nil
}
