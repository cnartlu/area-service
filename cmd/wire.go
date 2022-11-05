//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/cnartlu/area-service/internal"
	// "github.com/cnartlu/area-service/internal/command"
	"github.com/cnartlu/area-service/internal/server"
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/google/wire"
)

// initApp 初始化应用
func initApp(kconfig.Config) (*server.Server, func(), error) {
	panic(
		wire.Build(
			internal.ProviderSet,
			server.NewServer,
		),
	)
}

// initCommand 初始化命令行
// func initCommand(kconfig.Config) *command.Command {
// 	panic(wire.Build(
// 		command.ProviderSet,
// 		command.New,
// 	))
// }
