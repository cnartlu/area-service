package log

import (
	"github.com/cnartlu/area-service/pkg/config/logger"
	kconfig "github.com/go-kratos/kratos/v2/config"
)

type Option func(*Logger)

func WithConfigure(c kconfig.Config) Option {
	return func(o *Logger) {
		o.configure = c
	}
}

func WithConfig(c *logger.Config) Option {
	return func(o *Logger) {
		o.config = c
	}
}
