package asset

import (
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
	"github.com/cnartlu/area-service/pkg/component/log"
)

type ServiceInterface interface {
	Querier
	Exporter
	Importer
}

type Service struct {
	logger      *log.Logger
	application *config.Application
	repo        asset.RepositoryInterface
}

var _ ServiceInterface = (*Service)(nil)

func NewService(
	logger *log.Logger,
	application *config.Application,
	repo asset.RepositoryInterface,
) *Service {
	return &Service{
		logger:      logger,
		application: application,
		repo:        repo,
	}
}
