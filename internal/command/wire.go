//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package command

import (
	pkgcompant "github.com/cnartlu/area-service/component"
	"github.com/cnartlu/area-service/internal/biz"
	"github.com/cnartlu/area-service/internal/command/handler"
	"github.com/cnartlu/area-service/internal/command/script"
	"github.com/cnartlu/area-service/internal/component"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/data"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	pkgcompant.ProviderSet,
	component.ProviderSet,
	data.ProviderSet,
	biz.ProviderSet,
	script.ProviderSet,
	handler.ProviderSet,
)
