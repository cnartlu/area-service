// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/component/filesystem"
	"github.com/cnartlu/area-service/component/github"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/component/proxy"
	"github.com/cnartlu/area-service/component/redis"
	area2 "github.com/cnartlu/area-service/internal/biz/area"
	github3 "github.com/cnartlu/area-service/internal/biz/github"
	"github.com/cnartlu/area-service/internal/command"
	"github.com/cnartlu/area-service/internal/command/handler"
	github4 "github.com/cnartlu/area-service/internal/command/handler/github"
	"github.com/cnartlu/area-service/internal/command/script"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/data/area"
	"github.com/cnartlu/area-service/internal/data/area/release"
	"github.com/cnartlu/area-service/internal/data/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/data"
	github2 "github.com/cnartlu/area-service/internal/data/github"
	"github.com/cnartlu/area-service/internal/server"
	"github.com/cnartlu/area-service/internal/server/cron"
	"github.com/cnartlu/area-service/internal/server/cron/job"
	"github.com/cnartlu/area-service/internal/server/grpc"
	"github.com/cnartlu/area-service/internal/server/http"
	"github.com/cnartlu/area-service/internal/server/http/router"
	"github.com/cnartlu/area-service/internal/service"
)

// Injectors from wire.go:

// initApp 初始化应用
func initApp(contextContext context.Context, string2 string) (*server.Server, func(), error) {
	configConfig, err := config.NewByString(string2)
	if err != nil {
		return nil, nil, err
	}
	appConfig := config.GetApp(configConfig)
	appApp := app.New(appConfig)
	logConfig := config.GetLogger(configConfig)
	logger, err := log.New(logConfig)
	if err != nil {
		return nil, nil, err
	}
	configGrpc := config.GetGrpc(configConfig)
	redisConfig := config.GetRedis(configConfig)
	client, cleanup, err := redis.New(redisConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	databaseConfig := config.GetDb(configConfig)
	dataData, cleanup2, err := data.NewData(logger, client, databaseConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	areaRepo := area.NewAreaRepo(dataData)
	areaUsecase := area2.NewAreaUsecase(areaRepo)
	areaService := service.NewAreaService(areaUsecase)
	grpcServer := grpc.NewServer(logger, configGrpc, areaService)
	configHttp := config.GetHttp(configConfig)
	routerArea := router.NewArea(areaService)
	v := router.NewRouter(routerArea)
	httpServer := http.NewServer(appApp, logger, configHttp, v)
	daily := job.NewDaily(logger)
	cronServer := cron.NewServer(logger, daily)
	serverServer := server.NewServer(contextContext, appApp, logger, grpcServer, httpServer, cronServer)
	return serverServer, func() {
		cleanup2()
		cleanup()
	}, nil
}

// initCommand 初始化命令行
func initCommand(string2 string) (*command.Command, func(), error) {
	configConfig, err := config.NewByString(string2)
	if err != nil {
		return nil, nil, err
	}
	appConfig := config.GetApp(configConfig)
	appApp := app.New(appConfig)
	logConfig := config.GetLogger(configConfig)
	logger, err := log.New(logConfig)
	if err != nil {
		return nil, nil, err
	}
	redisConfig := config.GetRedis(configConfig)
	client, cleanup, err := redis.New(redisConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	databaseConfig := config.GetDb(configConfig)
	dataData, cleanup2, err := data.NewData(logger, client, databaseConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	filesystemConfig := config.GetFileSystem(configConfig)
	proxyClient, err := proxy.NewByAppConfig(appConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	fileSystem := filesystem.New(filesystemConfig, proxyClient)
	githubClient := github.New(proxyClient)
	xiangyuecnRepo := github2.NewXiangyuecnRepo(githubClient)
	areaRepo := area.NewAreaRepo(dataData)
	releaseRepo := release.NewRepository(dataData)
	assetRepo := asset.NewAssetRepo(dataData)
	githubUsecase := github3.NewGithubUsecase(appApp, logger, dataData, fileSystem, xiangyuecnRepo, areaRepo, releaseRepo, assetRepo)
	githubHandler := github4.NewHandler(githubUsecase)
	handlerHandler := handler.New(githubHandler)
	scriptScript := script.New()
	commandCommand := command.New(handlerHandler, scriptScript)
	return commandCommand, func() {
		cleanup2()
		cleanup()
	}, nil
}
