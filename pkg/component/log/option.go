package log

import (
	kconfig "github.com/go-kratos/kratos/v2/config"
)

type Option func(*Logger)

func WithConfigure(c kconfig.Config) Option {
	return func(o *Logger) {
		o.configure = c
	}
}

func WithConfig(c *Config) Option {
	return func(o *Logger) {
		o.config = c
	}
}
