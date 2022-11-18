package config

import (
	"os"

	kconfig "github.com/go-kratos/kratos/v2/config"
	kconfigFile "github.com/go-kratos/kratos/v2/config/file"
	klog "github.com/go-kratos/kratos/v2/log"
)

func NewKratos(config string) (kconfig.Config, func(), error) {
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
	cleanup := func() {
		if err := c.Close(); err != nil {
			klog.Error("error", err)
		}
	}
	if err := c.Load(); err != nil {
		cleanup()
		return nil, nil, err
	}
	return c, cleanup, nil
}
