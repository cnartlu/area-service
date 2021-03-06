// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/cnartlu/area-service/internal/app"
	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/cron"
	"github.com/cnartlu/area-service/internal/cron/job"
	"github.com/cnartlu/area-service/internal/transport"
	"github.com/cnartlu/area-service/internal/transport/grpc"
	"github.com/cnartlu/area-service/internal/transport/http"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/cnartlu/area-service/pkg/component/proxy"
	"github.com/cnartlu/area-service/pkg/component/redis"
)

// Injectors from wire.go:

// initApp 初始化应用
func initApp(logger *log.Logger, bootstrap *config.Bootstrap) (*app.App, func(), error) {
	redisConfig := bootstrap.Redis
	client, cleanup, err := redis.New(redisConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	dbConfig := bootstrap.Database
	dbDB, cleanup2, err := db.New(logger, dbConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	application := bootstrap.Application
	proxyConfig := application.Proxy
	httpClient := proxy.New(proxyConfig)
	github := job.NewGithub(logger, dbDB, httpClient)
	cronCron, err := cron.New(logger, client, dbDB, github)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server := bootstrap.Server
	server_HTTP := server.Http
	httpServer := http.NewHTTPServer(logger, server_HTTP)
	server_GRPC := server.Grpc
	grpcServer := grpc.NewGRPCServer(logger, server_GRPC)
	transportTransport := transport.New(logger, application, httpServer, grpcServer)
	appApp := app.New(logger, cronCron, transportTransport)
	return appApp, func() {
		cleanup2()
		cleanup()
	}, nil
}
