// The build tag makes sure the stub is not built in the final build.

package http

import (
	"github.com/google/wire"
)

// ProviderSet Http服务注入器
var ProviderSet = wire.NewSet(
	NewHTTPServer,
)
