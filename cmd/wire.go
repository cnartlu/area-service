//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/cnartlu/area-service/internal/app"
	"github.com/cnartlu/area-service/internal/command"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/google/wire"
)

// initApp 初始化应用
func initApp(*log.Logger, *config.Config) (*app.App, func(), error) {
	panic(
		wire.Build(
			app.ProviderSet,
			app.New,
		),
	)
}

// initCommand 初始化命令行
func initCommand(*log.Logger, *config.Config) (*command.Command, func(), error) {
	panic(wire.Build(
		command.ProviderSet,
		command.New,
	))
}
