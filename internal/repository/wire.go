package repository

import (
	"github.com/cnartlu/area-service/internal/repository/area"
	"github.com/cnartlu/area-service/internal/repository/area/release"
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	asset.NewRepository,
	wire.Bind(new(asset.RepositoryInterface), new(*asset.Repository)),
	release.NewRepository,
	wire.Bind(new(release.RepositoryInterface), new(*release.Repository)),
	area.NewRepository,
	wire.Bind(new(area.RepositoryInterface), new(*area.Repository)),
)
