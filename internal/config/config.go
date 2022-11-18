package config

import (
	"fmt"
	"os"

	app "github.com/cnartlu/area-service/component/app"
	db "github.com/cnartlu/area-service/component/db"
	log "github.com/cnartlu/area-service/component/log"
	redis "github.com/cnartlu/area-service/component/redis"
	"github.com/cnartlu/area-service/pkg/env"
	kconfig "github.com/go-kratos/kratos/v2/config"
)

func GetApp(c *Config) *app.Config {
	return c.GetApp()
}

func GetHttp(c *Config) *Http {
	return c.GetHttp()
}

func GetGrpc(c *Config) *Grpc {
	return c.GetGrpc()
}

func GetCron(c *Config) *Cron {
	return c.GetCron()
}

func GetLogger(c *Config) *log.Config {
	return c.GetLogger()
}

func GetRedis(c *Config) *redis.Config {
	return c.GetRedis()
}

func GetDb(c *Config) *db.Config {
	return c.GetDb()
}

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
