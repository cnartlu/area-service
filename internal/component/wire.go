package component

import (
	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.NewEnt,
)
