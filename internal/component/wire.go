//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package component

import (
	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.NewEnt,
)
