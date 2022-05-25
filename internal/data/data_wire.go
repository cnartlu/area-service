//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
)
