package config

import (
	"fmt"
	"os"

	"github.com/cnartlu/area-service/pkg/env"
	kconfig "github.com/go-kratos/kratos/v2/config"
)

func New(c kconfig.Config) (*Config, error) {
	var config = Config{
		Debug:  false,
		Env:    "production",
		Name:   "app",
		Logger: nil,
	}
	if c != nil {
		if err := c.Scan(&config); err != nil {
			return nil, err
		}
	}
	var debug string = "false"
	if config.Debug {
		debug = "true"
	}
	if err := os.Setenv(env.NameDebug, debug); err != nil {
		fmt.Println("set debug", err)
	}
	if err := os.Setenv(env.NameEnv, config.Env); err != nil {
		fmt.Println("set environment", err)
	}
	return &config, nil
}
