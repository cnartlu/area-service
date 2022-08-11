package repository

import (
	"github.com/cnartlu/area-service/internal/repository/area"
	"github.com/cnartlu/area-service/internal/repository/area/release"
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	asset.NewRepository,
	wire.Bind(new(asset.RepositoryManager), new(*asset.Repository)),
	release.NewRepository,
	wire.Bind(new(release.RepositoryManager), new(*release.Repository)),
	area.NewRepository,
	wire.Bind(new(area.RepositoryManager), new(*area.Repository)),
)
