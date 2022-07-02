package app

import (
	"context"

	"github.com/cnartlu/area-service/internal/cron"
	"github.com/cnartlu/area-service/internal/transport"
	"github.com/cnartlu/area-service/pkg/component/log"
)

//go:generate swag fmt -g app.go
//go:generate swag init -g app.go -o transport/http/api/docs --parseInternal

// @title                       API 接口文档
// @description                 API 接口文档
// @version                     0.0.0
// @host                        localhost
// @BasePath                    /api
// @schemes                     http https
// @accept                      json
// @accept                      x-www-form-urlencoded
// @produce                     json
// @securityDefinitions.apikey  Authorization
// @in                          header
// @name                        Token

type App struct {
	logger    *log.Logger
	cron      *cron.Cron
	transport *transport.Transport
}

// Start 启动应用
func (a *App) Start(cancel context.CancelFunc) (err error) {
	// 设置 tracer
	// if a.trace != nil {
	// 	otel.SetTracerProvider(a.trace.TracerProvider())
	// 	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
	// 		propagation.TraceContext{},
	// 		propagation.Baggage{},
	// 	))
	// }

	// 启动 cron 服务
	if err = a.cron.Start(); err != nil {
		return
	}

	// 启动 transport 服务
	go func() {
		if err = a.transport.Start(); err != nil {
			a.logger.Error(err.Error())
			cancel()
			return
		}
	}()

	return nil
}

// Stop 停止应用
func (a *App) Stop(ctx context.Context) (err error) {
	// 关闭 cron 服务
	if err = a.cron.Stop(ctx); err != nil {
		return
	}

	// 关闭 transport 服务
	if err = a.transport.Stop(); err != nil {
		return
	}

	return nil
}

// New 实例化应用
func New(
	logger *log.Logger,
	cron *cron.Cron,
	transport *transport.Transport,
) *App {
	return &App{
		logger:    logger,
		cron:      cron,
		transport: transport,
	}
}
