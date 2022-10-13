//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package config

import (
	// kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/google/wire"
)

// ProviderSet 配置项的功能
var ProviderSet = wire.NewSet(
	// wire.NewSet(wire.Value("config.yaml"), NewConfig),
	wire.FieldsOf(new(*Config), "App"),
	wire.FieldsOf(new(*App), "Http", "Grpc", "Cron", "Logger", "Redis", "Db"),
)
