// The build tag makes sure the stub is not built in the final build.

package http

import (
	"github.com/cnartlu/area-service/internal/transport/http/handler"

	"github.com/cnartlu/area-service/internal/transport/http/router"
	"github.com/google/wire"
)

// ProviderSet Http服务注入器
var ProviderSet = wire.NewSet(
	handler.ProviderSet,
	router.ProviderSet,
	NewHTTPServer,
)
