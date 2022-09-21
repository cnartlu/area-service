package server

import (
	"github.com/cnartlu/area-service/internal/config"
	serverhttp "github.com/cnartlu/area-service/internal/server/http"
	"github.com/cnartlu/area-service/pkg/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.uber.org/zap"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	logger *log.Logger,
	g *grpc.Server,
	c *config.Http,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c == nil {
		c = &config.Http{}
	}
	if c.Network != "" {
		opts = append(opts, http.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, http.Address(c.Addr))
	}
	if c.Timeout != nil {
		opts = append(opts, http.Timeout(c.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	grpcEndpoint, _ := g.Endpoint()
	engine, err := serverhttp.NewGin(grpcEndpoint.Host)
	if err != nil {
		logger.Error("http init failed", zap.Error(err))
		return srv
	}
	srv.HandlePrefix("", engine)
	return srv
}
