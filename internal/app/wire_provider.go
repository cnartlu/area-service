//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/cnartlu/area-service/internal/app/config"
	"github.com/cnartlu/area-service/internal/app/cron"
	"github.com/cnartlu/area-service/internal/app/repository"
	"github.com/cnartlu/area-service/internal/app/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	cron.ProviderSet,
	transport.ProviderSet,
	repository.ProviderSet,
	service.ProviderSet,
)
