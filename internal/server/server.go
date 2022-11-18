package server

import (
	"errors"
	"os"
	"strings"
	"syscall"

	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/internal/server/cron"
	"github.com/cnartlu/area-service/internal/server/grpc"
	"github.com/cnartlu/area-service/internal/server/http"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
)

type Server struct {
	app    *app.App
	logger *log.Logger
	server *kratos.App
}

func (s *Server) Run(signal string) error {
	if signal != "" {
		switch signal {
		case "quit":
			process, err := s.app.GetProcess().GetPidFile().GetProcess()
			if err != nil {
				return err
			}
			return process.Signal(os.Interrupt)
		case "stop":
			process, err := s.app.GetProcess().GetPidFile().GetProcess()
			if err != nil {
				return err
			}
			return process.Kill()
		case "reload":
			process, err := s.app.GetProcess().GetPidFile().GetProcess()
			if err != nil {
				return err
			}
			return process.Signal(syscall.SIGHUP)
		default:
			buf := strings.Builder{}
			buf.WriteString("nginx")
			buf.WriteString(": invalid option: \"-")
			buf.WriteString("s ")
			buf.WriteString(signal)
			buf.WriteString("\"")
			return errors.New(buf.String())
		}
	}
	return s.server.Run()
}

func (s *Server) Stop() error {
	return s.server.Stop()
}

func NewServer(
	app *app.App,
	logger *log.Logger,
	config *config.Config,
	// 服务对象
	gs *grpc.Server,
	hs *http.Server,
	cn *cron.Server,
) *Server {
	var servers []transport.Server
	if cn != nil {
		servers = append(servers, cn)
	}
	if gs != nil {
		servers = append(servers, gs)
	}
	if hs != nil {
		servers = append(servers, hs)
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
		app:    app,
		logger: logger,
		server: server,
	}
}
