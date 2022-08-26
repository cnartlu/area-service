package server

import (
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/transport"
)

type Server struct {
	logger *log.Logger
	server *kratos.App
}

func (s *Server) Start() {
	s.server.Run()
}

func (s *Server) Stop() error {
	return s.server.Stop()
}

func NewServer(
	logger *log.Logger,
	config config.Config,
) *Server {
	var servers []transport.Server
	var hs transport.Server
	var gs transport.Server
	if hs != nil {
		servers = append(servers, hs)
	}
	if gs != nil {
		servers = append(servers, gs)
	}

	options := []kratos.Option{
		// kratos.ID(hostname),
		// kratos.Name(appConf.Name),
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
