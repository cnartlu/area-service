package discovery

import (
	"github.com/cnartlu/area-service/component/discovery/consul"
	"github.com/cnartlu/area-service/component/discovery/etcd"

	"github.com/go-kratos/kratos/v2/registry"
	"go.uber.org/zap"
)

type Discovery interface {
	registry.Registrar
	registry.Discovery
}

type Config struct {
	Etcd   *etcd.Config
	Consul *consul.Config
}

func New(config *Config, logger *zap.Logger) (Discovery, error) {
	if config == nil {
		return nil, nil
	}

	if config.Etcd != nil {
		return etcd.New(config.Etcd, logger)
	}

	if config.Consul != nil {
		return consul.New(config.Consul)
	}

	return nil, nil
}
