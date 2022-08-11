//go:build wireinject
// +build wireinject

package tests

import (
	"github.com/cnartlu/area-service/internal/component"
	"github.com/cnartlu/area-service/internal/config"
	pCompant "github.com/cnartlu/area-service/pkg/component"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	NewConfig,
	log.NewDefault,
	config.ProviderSet,
	pCompant.ProviderSet,
	component.ProviderSet,
)

func Init() (*Tests, func(), error) {
	panic(wire.Build(
		providerSet,
		New,
	))
}
