package server

import (
	v1 "github.com/cnartlu/area-service/api/v1"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/service"
	"github.com/cnartlu/area-service/pkg/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	logger *log.Logger,
	c *config.Grpc,

	// 服务
	sArea *service.AreaService,
) *grpc.Server {
	if c == nil {
		c = &config.Grpc{}
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Network != "" {
		opts = append(opts, grpc.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, grpc.Address(c.Addr))
	}
	if c.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	v1.RegisterAreaServer(srv, sArea)

	return srv
}
