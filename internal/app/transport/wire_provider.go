//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package transport

import (
	// gtr "github.com/cnartlu/area-service/internal/app/transport/grpc"
	// htr "github.com/cnartlu/area-service/internal/app/transport/http"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	// gtr.ProviderSet,
	// htr.ProviderSet,
	New,
)
