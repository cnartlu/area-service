package jobs

import (
	"context"

	"go.uber.org/zap"

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
	err := s.releaseService.LoadLatestRelease(ctx)
	if err != nil {
		s.logger.Error("", zap.Error(err))
		return
	}
	err = s.releaseAssetService.Import(ctx, nil)
	if err != nil {
		s.logger.Error("", zap.Error(err))
		return
	}
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
