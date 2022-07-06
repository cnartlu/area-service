//go:build wireinject
// +build wireinject

package tests

import (
	"github.com/cnartlu/area-service/internal/component"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/cron"
	pCompant "github.com/cnartlu/area-service/pkg/component"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.NewDefault,
	config.ProviderSet,
	pCompant.ProviderSet,
	component.ProviderSet,
	cron.ProviderSet,
	NewCronJob,
)

func Init() (*Tests, func(), error) {
	panic(wire.Build(
		providerSet,
		New,
	))
}
