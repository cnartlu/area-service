//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer)
