package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"

	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/component/log"
	"github.com/gin-gonic/gin"
	grpctransport "github.com/go-kratos/kratos/v2/transport/grpc"
	"golang.org/x/sys/windows"
)

type Server struct {
	*http.Server
	// logger 日志
	logger *log.Logger
	// config 配置
	config *config.Config
	// router 引擎
	router *gin.Engine
	// lis 监听器
	lis net.Listener
	// fd 文件
	fd *os.File
}

// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s.Handler.ServeHTTP(res, req)
}

// Start start the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	// log.Infof("[HTTP] server listening on: %s", s.lis.Addr().String())
	// 实例化gin
	// if c.App.GetHttp() != nil {
	// 	// ch := c.App.GetHttp()
	// 	if c.App.GetDebug() {
	// 		gin.SetMode(gin.DebugMode)
	// 	} else {
	// 		switch c.App.GetEnv() {
	// 		case "prod", "production":
	// 			gin.SetMode(gin.ReleaseMode)
	// 		default:
	// 			gin.SetMode(gin.TestMode)
	// 		}
	// 	}
	// }

	httpConfig := s.config.GetHttp()
	addr := httpConfig.GetAddr()
	if addr == "" {
		addr = ":http"
	}

	// 需要先检查相关操作
	// 是否存在fd的数据入参
	if s.lis != nil {
		tcpAddr, err := net.ResolveTCPAddr(httpConfig.GetNetwork(), addr)
		if err == nil {
			return err
		}
		addr := s.lis.Addr()
		if addr.Network() == tcpAddr.Network() && addr.String() == tcpAddr.String() {
			s.lis.Close()
			s.lis = nil
		}
	}

	// 启动默认的监听器
	if s.lis == nil {
		l := net.ListenConfig{
			Control: func(network, address string, c syscall.RawConn) error {
				c.Control(func(fd uintptr) {
					err := windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
					if err != nil {
						panic(err)
					}
					// syscall.SetsockoptInt(syscall.Handle(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
					// syscall.SetsockoptInt(syscall.Handle(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
					// unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
					// unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
					fmt.Println("open http fd", int(fd))
				})
				return nil
			},
		}
		ln, err := l.Listen(ctx, "tcp", addr)
		if err != nil {
			return err
		}
		var ff *os.File
		switch v := ln.(type) {
		case *net.TCPListener:
			ff, err = v.File()
		case *net.UnixListener:
			ff, err = v.File()
		}
		if err != nil {
			fmt.Println("os failed.error", err)
		} else {
			s.fd = ff
			lis, err := net.FileListener(s.fd)
			if err != nil {
				fmt.Println("FileListener.error", err)
			} else {
				fmt.Println("lis", lis.Addr())
			}
		}

		s.lis = ln
	}

	if err := s.Serve(s.lis); err != nil {
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
	logger *log.Logger,
	c *config.Config,
	g *grpctransport.Server,
	// 其他数据
) *Server {
	h := http.Server{
		Handler: gin.New(),
	}
	srv := Server{
		Server: &h,
		logger: logger,
	}
	return &srv
}
