package jobs

import (
	"context"

	releaseservice "github.com/cnartlu/area-service/internal/service/area/release"

	"github.com/cnartlu/area-service/pkg/component/log"
)

type SyncArea struct {
	logger         *log.Logger
	releaseService releaseservice.Servicer
}

func (s *SyncArea) Run() {
	s.logger.Info("load github area")
	ctx := context.Background()
	s.releaseService.Import(ctx, nil)
}

func NewSyncArea(
	logger *log.Logger,
	releaseService releaseservice.Servicer,
) *SyncArea {
	return &SyncArea{
		logger:         logger,
		releaseService: releaseService,
	}
}
