//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package handler

import (
	"github.com/cnartlu/area-service/internal/command/handler/github"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	github.NewHandler,
	New,
)
