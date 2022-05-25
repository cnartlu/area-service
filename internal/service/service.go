//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewGreeterService,
)
