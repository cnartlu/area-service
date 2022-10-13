package grpc

import (
	v1 "github.com/cnartlu/area-service/api/v1"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/service"
	"github.com/cnartlu/area-service/pkg/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
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
	var opts = []grpc.ServerOption{
		grpc.Logger(logger.AddCallerSkip(1)),
		grpc.Middleware(recovery.Recovery()),
	}
	if c.GetNetwork() != "" {
		opts = append(opts, grpc.Network(c.GetNetwork()))
	}
	if c.GetAddr() != "" {
		opts = append(opts, grpc.Address(c.GetAddr()))
	}
	if c.GetTimeout() != nil {
		opts = append(opts, grpc.Timeout(c.GetTimeout().AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	v1.RegisterAreaServer(srv, sArea)

	return srv
}
