// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/cnartlu/area-service/internal/component"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/cron"
	"github.com/cnartlu/area-service/internal/repository"
	"github.com/cnartlu/area-service/internal/service"
	"github.com/cnartlu/area-service/internal/transport"
	pCompant "github.com/cnartlu/area-service/pkg/component"
	"github.com/cnartlu/area-service/pkg/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	pCompant.ProviderSet,
	component.ProviderSet,
	cron.ProviderSet,
	transport.ProviderSet,
	repository.ProviderSet,
	service.ProviderSet,
)
