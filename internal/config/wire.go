// The build tag makes sure the stub is not built in the final build.

package config

import (
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/google/wire"
)

// ProviderSet 配置项的功能
var ProviderSet = wire.NewSet(
	wire.FieldsOf(new(*Config), "Bootstrap"),
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
	wire.FieldsOf(new(*Server), "Http", "Grpc"),
	wire.FieldsOf(new(*FileSystem), "Disks"),
	wire.FieldsOf(new(*Cache), "Stores"),
	wire.FieldsOf(new(*log.Config), "Loggers"),
)
