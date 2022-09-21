package biz

import (
	"github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/data"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	data.ProviderSet,
	area.NewManagerUsecase,
	release.NewManaement,
)
