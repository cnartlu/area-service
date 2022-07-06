package config

import (
	"fmt"

	"github.com/cnartlu/area-service/pkg/path"
	kconfig "github.com/go-kratos/kratos/v2/config"
	kconfigFile "github.com/go-kratos/kratos/v2/config/file"
)

type Config struct {
	Config    kconfig.Config
	Bootstrap *Bootstrap
}

func New() (*Config, func(), error) {
	rootPath := path.RootPath()
	var configPaths = []string{"config.yaml", rootPath + "/config.yaml", rootPath + "/etc/config.yaml"}
	var options []kconfig.Option
	for _, path := range configPaths {
		options = append(options, kconfig.WithSource(kconfigFile.NewSource(path)))
	}
	config := kconfig.New(options...)
	if err := config.Load(); err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		config.Close()
	}
	var v interface{}
	config.Scan(&v)
	fmt.Println(v)
	var bootstrap Bootstrap
	if err := config.Scan(&bootstrap); err != nil {
		return nil, cleanup, err
	}
	c := &Config{
		Config:    config,
		Bootstrap: &bootstrap,
	}
	return c, cleanup, nil
}
