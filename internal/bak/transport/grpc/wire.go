// The build tag makes sure the stub is not built in the final build.

package grpc

import (
	"github.com/google/wire"
)

// ProviderSet Grpc服务注入器
var ProviderSet = wire.NewSet(
	NewGRPCServer,
)
