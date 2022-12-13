package grpc

import (
	"context"
	"net"
	"syscall"

	v1 "github.com/cnartlu/area-service/api/v1"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/service"
	"github.com/go-kratos/kratos/v2/transport/grpc"
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
			return nil
		},
	}
	lis, _ := l.Listen(context.Background(), network, addr)
	var opts = []grpc.ServerOption{
		grpc.Network(network),
		grpc.Address(addr),
		grpc.Listener(lis),
		// grpc.Middleware(recovery.Recovery(), logging.Server(ilog.NewKratosLogger(logger.AddCallerSkip(1)))),
	}
	if c.GetTimeout() != nil {
		opts = append(opts, grpc.Timeout(c.GetTimeout().AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	v1.RegisterAreaServer(srv, sArea)

	return srv
}
