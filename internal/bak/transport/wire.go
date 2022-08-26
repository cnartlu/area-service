// The build tag makes sure the stub is not built in the final build.
package transport

import (
	gtr "github.com/cnartlu/area-service/internal/transport/grpc"
	htr "github.com/cnartlu/area-service/internal/transport/http"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	gtr.ProviderSet,
	htr.ProviderSet,
	New,
)
