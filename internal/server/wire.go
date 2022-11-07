//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package server

import (
	"github.com/cnartlu/area-service/internal/server/cron"
	"github.com/cnartlu/area-service/internal/server/grpc"
	"github.com/cnartlu/area-service/internal/server/http"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	cron.ProviderSet,
	grpc.NewServer,
	http.ProviderSet,
)
