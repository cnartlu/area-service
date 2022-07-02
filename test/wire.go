//go:build wireinject
// +build wireinject

package tests

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewZapLogger,
	NewLogger,
	NewDB,
	NewRDB,
)

func Init() (*Tests, func(), error) {
	panic(wire.Build(
		ProviderSet,
		New,
	))
}
