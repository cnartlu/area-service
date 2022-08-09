package release

import (
	"github.com/cnartlu/area-service/internal/repository/area/release"
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
	"github.com/cnartlu/area-service/pkg/component/log"
)

type ServiceInterface interface {
	GithubInterface
	Querier
}

type Service struct {
	logger      *log.Logger
	releaseRepo release.RepositoryInterface
	assetRepo   asset.RepositoryInterface
}

var _ ServiceInterface = (*Service)(nil)

func NewService(
	logger *log.Logger,
	releaseRepo release.RepositoryInterface,
	assetRepo asset.RepositoryInterface,
) *Service {
	return &Service{
		logger:      logger,
		releaseRepo: releaseRepo,
		assetRepo:   assetRepo,
	}
}
