package config

import (
	app "github.com/cnartlu/area-service/component/app"
	db "github.com/cnartlu/area-service/component/db"
	filesystem "github.com/cnartlu/area-service/component/filesystem"
	log "github.com/cnartlu/area-service/component/log"
	redis "github.com/cnartlu/area-service/component/redis"
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

func GetFileSystem(c *Config) *filesystem.Config {
	return c.GetFilesystem()
}

func New(c kconfig.Config) (*Config, error) {
	var config = Config{}
	if c != nil {
		if err := c.Scan(&config); err != nil {
			return nil, err
		}
	}
	return &config, nil
}
