package release

import (
	"github.com/cnartlu/area-service/internal/repository/area/release"
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
	"github.com/cnartlu/area-service/internal/service/area/release/asset/importer"
	"github.com/cnartlu/area-service/pkg/component/log"
)

type Servicer interface {
	GithubInterface
	Querier
	Importer
}

type Service struct {
	logger          *log.Logger
	releaseRepo     release.RepositoryManager
	assetRepo       asset.RepositoryManager
	importerService importer.Servicer
}

var _ Servicer = (*Service)(nil)

func NewService(
	logger *log.Logger,
	releaseRepo release.RepositoryManager,
	assetRepo asset.RepositoryManager,
	importerService importer.Servicer,
) *Service {
	return &Service{
		logger:          logger,
		releaseRepo:     releaseRepo,
		assetRepo:       assetRepo,
		importerService: importerService,
	}
}
