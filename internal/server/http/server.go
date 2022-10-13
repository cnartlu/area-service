package http

import (
	"strings"
	"time"

	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/server/http/router"
	"github.com/cnartlu/area-service/pkg/log"
	"github.com/cnartlu/area-service/pkg/middleware"
	"github.com/gin-gonic/gin"
	grpctransport "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server = http.Server

// NewServer new a HTTP server.
func NewServer(
	logger *log.Logger,
	c *config.Config,
	g *grpctransport.Server,
	// 其他数据
) *Server {
	var opts = []http.ServerOption{
		http.Timeout(0),
		http.Logger(logger),
	}
	if c.App.GetHttp() != nil {
		ch := c.App.GetHttp()
		if ch.GetNetwork() != "" {
			opts = append(opts, http.Network(ch.GetNetwork()))
		}
		if ch.GetAddr() != "" {
			opts = append(opts, http.Address(ch.GetAddr()))
		}
		if c.App.GetDebug() {
			gin.SetMode(gin.DebugMode)
		} else {
			switch c.App.GetEnv() {
			case "prod", "production":
				gin.SetMode(gin.ReleaseMode)
			default:
				gin.SetMode(gin.TestMode)
			}
		}
	}

	// 实例化gin
	var e = gin.New()
	{
		// 实例化数据
		e.HandleMethodNotAllowed = true
		rootGroup := e.Group("/")
		rootGroup.Use(
			func(c *gin.Context) {
				c.Request.Context()
			},
			middleware.LoggerMiddleware(logger.AddCallerSkip(1)),
			middleware.TimeoutMiddleware(time.Second*10),
			gin.Recovery(),
		)

		endpoint, _ := g.Endpoint()
		endpoint.Scheme = ""
		logger.Debug(endpoint.String())
		gr, err := grpc.Dial(strings.TrimLeft(endpoint.String(), "//"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Error("grpc endpoint path is invalid", zap.Error(err))
			return nil
		}
		// 注册其他路由
		router.NewArea(gr, rootGroup)
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("", e)

	return srv
}
