package asset

import (
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
	"github.com/cnartlu/area-service/internal/service/area/release/asset/importer"
	"github.com/cnartlu/area-service/pkg/component/log"
)

type Servicer interface {
	Querier
	Exporter
	Importer
}

type Service struct {
	logger          *log.Logger
	repo            asset.RepositoryManager
	importerService importer.Service
}

var _ Servicer = (*Service)(nil)

func NewService(
	logger *log.Logger,
	repo asset.RepositoryManager,
	importerService importer.Service,
) *Service {
	return &Service{
		logger:          logger,
		repo:            repo,
		importerService: importerService,
	}
}
