//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package config

import (
	"github.com/google/wire"
)

// ProviderSet 配置项的功能
var ProviderSet = wire.NewSet(
	New,
	wire.FieldsOf(new(*Config), "Http", "Grpc", "Cron", "Logger", "Redis", "Db"),
)
