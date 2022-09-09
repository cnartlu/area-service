package repository

import (
	"github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	area.NewManagerUsecase,
	release.NewManaement,
)
