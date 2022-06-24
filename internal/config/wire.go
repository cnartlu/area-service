// The build tag makes sure the stub is not built in the final build.

package config

import (
	"github.com/cnartlu/area-service/pkg/config/logger"
	"github.com/google/wire"
)

// ProviderSet 配置项的功能
var ProviderSet = wire.NewSet(
	// new(*Redis),
	wire.FieldsOf(
		new(*Bootstrap),
		"Application",
		"Server",
		"Database",
		"FileSystem",
		"Logger",
		"Cache",
		"Redis",
	),
	wire.FieldsOf(new(*Server), "HTTP", "GRPC"),
	wire.FieldsOf(new(*Database), "Connections"),
	wire.FieldsOf(new(*FileSystem), "Disks"),
	wire.FieldsOf(new(*Cache), "Stores"),
	wire.FieldsOf(new(*logger.Config), "Loggers"),
)
