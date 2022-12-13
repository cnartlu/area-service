package config

import (
	"os"

	app "github.com/cnartlu/area-service/component/app"
	db "github.com/cnartlu/area-service/component/database"
	filesystem "github.com/cnartlu/area-service/component/filesystem"
	log "github.com/cnartlu/area-service/component/log"
	redis "github.com/cnartlu/area-service/component/redis"
	kconfig "github.com/go-kratos/kratos/v2/config"
	kconfigFile "github.com/go-kratos/kratos/v2/config/file"
	klog "github.com/go-kratos/kratos/v2/log"
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

func NewKratos(config string) (kconfig.Config, error) {
	var sources []kconfig.Source
	var filenames []string = []string{
		config,
		"config.yaml",
		"config.json",
	}
	for _, filename := range filenames {
		if filename == "" {
			continue
		}
		if _, err := os.Stat(filename); err != nil {
			continue
		}
		sources = append(sources, kconfigFile.NewSource(filename))
	}
	logger := klog.NewFilter(klog.DefaultLogger, klog.FilterLevel(klog.LevelError))
	klog.SetLogger(logger)
	var c = kconfig.New(
		kconfig.WithLogger(logger),
		kconfig.WithSource(sources...),
	)
	if err := c.Load(); err != nil {
		_ = c.Close()
		return nil, err
	}
	return c, nil
}

func New(c kconfig.Config) (*Config, error) {
	var config = Config{}
	if c != nil {
		if err := c.Scan(&config); err != nil {
			return nil, err
		}
		if err := c.Close(); err != nil {
			return nil, err
		}
	}
	return &config, nil
}

func NewByString(s string) (*Config, error) {
	c, err := NewKratos(s)
	if err != nil {
		return nil, err
	}
	return New(c)
}
