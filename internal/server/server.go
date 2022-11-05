package server

import (
	"os"
	"os/exec"

	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/server/cron"
	"github.com/cnartlu/area-service/internal/server/grpc"
	"github.com/cnartlu/area-service/internal/server/http"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
)

type Server struct {
	logger *log.Logger
	server *kratos.App
	//
	apps []interface{ Restart() }
}

func (s *Server) Start() error {
	return s.server.Run()
}

func (s *Server) Restart() error {
	if os.Getenv("parent") == "" {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = append(os.Environ(), "listener=", "parent=1")
		return cmd.Run()
	}
	return nil
}

func (s *Server) Stop() error {
	return s.server.Stop()
}

func NewServer(
	logger *log.Logger,
	config *config.Config,
	// 服务对象
	gs *grpc.Server,
	hs *http.Server,
	cn *cron.Server,
) *Server {
	var servers []transport.Server

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
		kratos.ID(config.GetName()),
		kratos.Name(config.GetName()),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(log.NewKratosLogger(logger)),
		kratos.Server(servers...),
	}

	server := kratos.New(options...)
	return &Server{
		logger: logger,
		server: server,
	}
}
