package service

import (
	"github.com/cnartlu/area-service/internal/service/area"
	"github.com/cnartlu/area-service/internal/service/area/release"
	"github.com/cnartlu/area-service/internal/service/area/release/asset"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	asset.NewService,
	wire.Bind(new(asset.ServiceInterface), new(*asset.Service)),
	release.NewService,
	wire.Bind(new(release.ServiceInterface), new(*release.Service)),
	area.NewService,
	wire.Bind(new(area.ServiceInterface), new(*area.Service)),
)
