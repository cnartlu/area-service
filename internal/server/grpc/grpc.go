package grpc

import (
	"context"
	"fmt"
	"net"
	"syscall"

	v1 "github.com/cnartlu/area-service/api/v1"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/service"
	"github.com/cnartlu/area-service/component/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"golang.org/x/sys/windows"
)

type Server = grpc.Server

// NewServer new a gRPC server.
func NewServer(
	logger *log.Logger,
	c *config.Grpc,

	// 服务
	sArea *service.AreaService,
) *Server {
	var network = c.GetNetwork()
	var addr = c.GetAddr()
	if network == "" {
		network = "tcp"
	}
	if addr == "" {
		addr = ":0"
	}
	l := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			c.Control(func(fd uintptr) {
				err := windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
				if err != nil {
					panic(err)
				}
				// syscall.SetsockoptInt(syscall.Handle(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
				// syscall.SetsockoptInt(syscall.Handle(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
				fmt.Println("open grpc fd", int(fd))
			})
			return nil
		},
	}
	lis, _ := l.Listen(context.Background(), network, addr)
	var opts = []grpc.ServerOption{
		grpc.Network(network),
		grpc.Address(addr),
		grpc.Listener(lis),
		grpc.Logger(log.NewKratosLogger(logger.AddCallerSkip(1))),
		grpc.Middleware(recovery.Recovery()),
	}
	if c.GetTimeout() != nil {
		opts = append(opts, grpc.Timeout(c.GetTimeout().AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	v1.RegisterAreaServer(srv, sArea)

	return srv
}
