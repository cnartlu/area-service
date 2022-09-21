package data

import (
	bizarea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/area"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	area.NewAreaRepo,
	wire.Bind(new(bizarea.Manager), new(*area.AreaRepo)),
)
