// The build tag makes sure the stub is not built in the final build.

package config

import (
	"github.com/google/wire"
)

// ProviderSet 配置项的功能
var ProviderSet = wire.NewSet(
	// wire.NewSet(wire.Bind(new(kconfig.Config), new(*Config))),
	wire.FieldsOf(new(*Config), "Config"),
	wire.FieldsOf(new(*App), "Http", "Grpc", "Cron", "Logger", "Redis", "Db"),
)
