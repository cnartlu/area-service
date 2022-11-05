//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package internal

import (
	pkgcompant "github.com/cnartlu/area-service/component"
	"github.com/cnartlu/area-service/internal/biz"
	"github.com/cnartlu/area-service/internal/component"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/server"
	"github.com/cnartlu/area-service/internal/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	pkgcompant.ProviderSet,
	component.ProviderSet,
	biz.ProviderSet,
	service.ProviderSet,
	server.ProviderSet,
)
