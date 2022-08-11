package config

import (
	"os"
	"path/filepath"

	"github.com/cnartlu/area-service/pkg/utils"
	kconfig "github.com/go-kratos/kratos/v2/config"
	kconfigFile "github.com/go-kratos/kratos/v2/config/file"
)

type Config struct {
	Config    kconfig.Config
	Bootstrap *Bootstrap
}

func (c *Config) Close() error {
	return c.Config.Close()
}

func New(filename string) (*Config, error) {
	// 初始化配置器
	var options []kconfig.Option
	if filename == "" {
		filename = "config.yaml"
	}
	var (
		sources   []kconfig.Source
		filenames []string = []string{filename}
	)
	filenames = append(
		filenames,
		filepath.Join("config", filename),
		filepath.Join("..", "config", filename),
		filepath.Join("etc", filename),
		filepath.Join(utils.RootPath(), "etc", filename),
		filepath.Join(utils.RootPath(), "../etc", filename),
	)
	for _, filename := range filenames {
		_, err := os.Stat(filename)
		if err == nil {
			sources = append(sources, kconfigFile.NewSource(filename))
			break
		}
	}
	options = append(options, kconfig.WithSource(sources...))
	config := kconfig.New(options...)
	if err := config.Load(); err != nil {
		return nil, err
	}
	var bootstrap Bootstrap
	if err := config.Scan(&bootstrap); err != nil {
		defer config.Close()
		return nil, err
	}
	// 初始化默认配置数据
	if bootstrap.GetApplication() == nil {
		bootstrap.Application = &Application{
			Name:  "",
			Debug: false,
			Env:   "Prod",
			Proxy: "",
		}
	}
	return &Config{
		Config:    config,
		Bootstrap: &bootstrap,
	}, nil
}
