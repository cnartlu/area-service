package http

import (
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	logger *log.Logger,
	c *config.Server_HTTP,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c == nil {
		c = &config.Server_HTTP{}
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

	return srv
}
