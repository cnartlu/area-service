package jobs

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/cnartlu/area-service/internal/component/ent"
	releaseservice "github.com/cnartlu/area-service/internal/service/area/release"
	serviceReleaseAsset "github.com/cnartlu/area-service/internal/service/area/release/asset"

	"github.com/cnartlu/area-service/pkg/component/log"
)

type SyncArea struct {
	logger              *log.Logger
	releaseService      releaseservice.ServiceInterface
	releaseAssetService serviceReleaseAsset.ServiceInterface
}

func (s *SyncArea) Run() {
	s.logger.Info("load github area")
	ctx := context.Background()
	areaRelease, err := s.releaseService.LoadLatestRelease(ctx)
	if err != nil {
		s.logger.Error("拉取加载最后发布的区域资源失败", zap.Error(err))
		return
	}
	releaseAssets, err := s.releaseAssetService.FindListByReleaseID(ctx, areaRelease.ID)
	if err != nil {
		s.logger.Error("查找资源失败", zap.Uint64("areaRelease", areaRelease.ID), zap.Error(err))
		return
	}
	var waitGroup = &sync.WaitGroup{}
	for _, releaseAsset := range releaseAssets {
		waitGroup.Add(1)
		go func(areaReleaseAsset *ent.AreaReleaseAsset) {
			defer waitGroup.Done()
			err = s.releaseAssetService.Import(ctx, areaReleaseAsset)
			if err != nil {
				s.logger.Error(
					"导入资源文件失败",
					zap.Uint64("areaRelease", areaRelease.ID),
					zap.Uint64("areaReleaseAssetID", areaReleaseAsset.ID),
					zap.Error(err),
				)
			}
		}(releaseAsset)
	}
	waitGroup.Wait()
}

func NewSyncArea(
	logger *log.Logger,
	releaseService releaseservice.ServiceInterface,
	releaseAssetService serviceReleaseAsset.ServiceInterface,
) *SyncArea {
	return &SyncArea{
		logger:              logger,
		releaseService:      releaseService,
		releaseAssetService: releaseAssetService,
	}
}
