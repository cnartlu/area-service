package service

import (
	"github.com/cnartlu/area-service/internal/service/area"
	"github.com/cnartlu/area-service/internal/service/area/release"
	"github.com/cnartlu/area-service/internal/service/area/release/asset"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	asset.NewService,
	release.NewService,
	area.NewService,
)
