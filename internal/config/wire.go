// The build tag makes sure the stub is not built in the final build.

package config

import (
	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/google/wire"
)

// ProviderSet 配置项的功能
var ProviderSet = wire.NewSet(
	New,
	wire.FieldsOf(new(*Config), "Config", "Bootstrap"),
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
	wire.FieldsOf(new(*Application), "Proxy"),
	wire.FieldsOf(new(*Server), "HTTP", "GRPC"),
	wire.FieldsOf(new(*db.Config), "Connections"),
	wire.FieldsOf(new(*FileSystem), "Disks"),
	wire.FieldsOf(new(*Cache), "Stores"),
	wire.FieldsOf(new(*log.Config), "Loggers"),
)
