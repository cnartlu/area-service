package server

import (
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/log"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Server struct {
	logger *log.Logger
	server *kratos.App
}

func (s *Server) Start() error {
	s.logger.Info("starting server")
	return s.server.Run()
}

func (s *Server) Stop() error {
	s.logger.Info("stop server")
	return s.server.Stop()
}

func NewServer(
	logger *log.Logger,
	config *config.Config,
	gs *grpc.Server,
	hs *http.Server,
	cn *Cron,
) *Server {
	var servers []transport.Server
	// gs := NewGRPCServer(logger, config.Config.GetGrpc())
	// hs := NewHTTPServer(logger, config.Config.GetHttp())
	// cn := NewCronServer(logger)
	if hs != nil {
		servers = append(servers, hs)
	}
	if gs != nil {
		servers = append(servers, gs)
	}
	if cn != nil {
		servers = append(servers, cn)
	}

	options := []kratos.Option{
		kratos.ID("app.id"),
		kratos.Name("app.name"),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
	}

	// if discovery != nil {
	// 	options = append(options, kratos.Registrar(discovery))
	// }

	server := kratos.New(options...)
	return &Server{
		logger: logger,
		server: server,
	}
}
