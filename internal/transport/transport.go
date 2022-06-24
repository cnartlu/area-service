package transport

import (
	"os"

	// "github.com/cnartlu/area-service/internal/component/discovery"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var hostname, _ = os.Hostname()

type Transport struct {
	logger *log.Logger
	server *kratos.App
}

func New(
	logger *log.Logger,
	appConf *config.Application,
	hs *http.Server,
	gs *grpc.Server,
	// discovery discovery.Discovery,
) *Transport {
	var servers []transport.Server
	if hs != nil {
		servers = append(servers, hs)
	}
	if gs != nil {
		servers = append(servers, gs)
	}

	options := []kratos.Option{
		kratos.ID(hostname),
		kratos.Name(appConf.Name),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
	}

	// if discovery != nil {
	// 	options = append(options, kratos.Registrar(discovery))
	// }

	server := kratos.New(options...)

	return &Transport{
		logger: logger,
		server: server,
	}
}

func (t *Transport) Start() error {
	t.logger.Info("transport server starting ...")

	if err := t.server.Run(); err != nil {
		return err
	}
	return nil
}

func (t *Transport) Stop() error {
	if err := t.server.Stop(); err != nil {
		return err
	}

	t.logger.Info("transport server stopping ...")
	return nil
}
