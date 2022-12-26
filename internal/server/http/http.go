package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"syscall"

	v1 "github.com/cnartlu/area-service/api/manage/v1"
	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/server/http/gin/middleware"
	"github.com/cnartlu/area-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*http.Server
	// logger 日志
	logger *log.Logger
	// config 配置
	config *config.Http
	// router 引擎
	router *gin.Engine
}

// Start start the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	addr := s.config.GetAddr()
	if addr == "" {
		addr = ":http"
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	s.logger.Info("[HTTP] server listening on: [::]" + tcpAddr.String())
	// 判断是否为子进程
	// 1、当为子进程时，需要继承父进程的socket监听
	var lis net.Listener
	if lis != nil {
		lisAddr := lis.Addr()
		if lisAddr.Network() != tcpAddr.Network() || lisAddr.String() != tcpAddr.String() {
			if err := lis.Close(); err != nil {
				return err
			}
			lis = nil
		}
	}
	// 默认监听器
	if lis == nil {
		l := net.ListenConfig{
			Control: func(network, address string, c syscall.RawConn) error {
				return nil
			},
		}
		var err error
		lis, err = l.Listen(ctx, tcpAddr.Network(), tcpAddr.String())
		if err != nil {
			return err
		}
		defer lis.Close()
		// ln := lis.(*net.TCPListener)
		// _, err = ln.File()
		// if err != nil && syscall.geterror() != err {
		// 	return err
		// }
		// if err := os.Setenv(env.NameSocketHttp, strconv.FormatInt(int64(fd.Fd()), 10)); err != nil {
		// 	return err
		// }
	}
	if err := s.Serve(lis); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}
	return nil
}

// Stop stop the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("[HTTP] server stopping")
	return s.Shutdown(ctx)
}

// NewServer new a HTTP server.
func NewServer(
	a *app.App,
	logger *log.Logger,
	httpConfig *config.Http,
	// 其他数据
	areaService *service.AreaService,
) *Server {
	switch a.GetEnv() {
	case app.EnvName_dev:
		gin.SetMode(gin.DebugMode)
	case app.EnvName_test:
		gin.SetMode(gin.TestMode)
	case app.EnvName_sit:
		gin.SetMode(gin.DebugMode)
	case app.EnvName_uat:
		gin.SetMode(gin.ReleaseMode)
	case app.EnvName_prod:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.New()
	g1 := e.Group("/")
	{
		g1.Use(middleware.Recover(logger), middleware.Logger(logger), middleware.Response())
		v1.RegisterAreaGinServer(g1, areaService)
	}
	h := http.Server{
		Handler: e,
	}
	srv := Server{
		Server: &h,
		logger: logger,
		config: httpConfig,
	}
	return &srv
}
