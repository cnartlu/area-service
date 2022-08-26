// The build tag makes sure the stub is not built in the final build.
package internal

import (
	"github.com/cnartlu/area-service/internal/component"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/cron"
	"github.com/cnartlu/area-service/internal/server"
	"github.com/cnartlu/area-service/internal/service"
	pCompant "github.com/cnartlu/area-service/pkg/component"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	pCompant.ProviderSet,
	component.ProviderSet,
	cron.ProviderSet,
	service.ProviderSet,
	wire.NewSet(
		server.NewCronServer,
		server.NewGRPCServer,
		server.NewHTTPServer,
	),
)
