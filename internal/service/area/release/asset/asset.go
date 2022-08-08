package asset

import (
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
	"github.com/cnartlu/area-service/pkg/component/log"
)

type ServiceInterface interface {
	Exporter
	Importer
}

type Service struct {
	logger *log.Logger
	repo   asset.RepositoryInterface
}

var _ ServiceInterface = (*Service)(nil)

func NewService(logger *log.Logger, repo asset.RepositoryInterface) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}
