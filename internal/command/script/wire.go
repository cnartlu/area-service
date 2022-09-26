//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package script

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewS0000000000,
)
