package service

import (
	"github.com/cnartlu/area-service/internal/service/area"
	"github.com/cnartlu/area-service/internal/service/area/release"
	"github.com/cnartlu/area-service/internal/service/area/release/asset"
	"github.com/cnartlu/area-service/internal/service/area/release/asset/importer"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	importer.NewService,
	wire.Bind(new(importer.Servicer), new(*importer.Service)),
	asset.NewService,
	wire.Bind(new(asset.Servicer), new(*asset.Service)),
	release.NewService,
	wire.Bind(new(release.Servicer), new(*release.Service)),
	area.NewService,
	wire.Bind(new(area.Servicer), new(*area.Service)),
)
