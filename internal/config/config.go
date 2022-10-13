package config

import (
	"os"
	"path/filepath"

	pkgfilepath "github.com/cnartlu/area-service/pkg/filepath"

	_ "github.com/cnartlu/area-service/pkg/log/target"
	kconfig "github.com/go-kratos/kratos/v2/config"
	kconfigFile "github.com/go-kratos/kratos/v2/config/file"
)

var (
	_ Configure = (*Config)(nil)
)

type Configure interface {
	kconfig.Config
}

type Config struct {
	c   kconfig.Config
	App *App
}

// Load 加载数据
func (c *Config) Load() error {
	return c.c.Load()
}

// Scan 解析赋值
func (c *Config) Scan(v interface{}) error {
	return c.c.Scan(v)
}

// Value 获取值
func (c *Config) Value(key string) kconfig.Value {
	return c.c.Value(key)
}

// Watch 监听修改
func (c *Config) Watch(key string, w kconfig.Observer) error {
	return c.c.Watch(key, w)
}

// Close 关闭配置器
func (c *Config) Close() error {
	return c.c.Close()
}

// Config 获取应用配置
func (c *Config) GetApp() *App {
	if c.App == nil {
		c.App.Reset()
	}
	return c.App
}

func NewConfig(filename string) (*Config, error) {
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
		filepath.Join(pkgfilepath.RootPath(), "etc", filename),
		filepath.Join(pkgfilepath.RootPath(), "../etc", filename),
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
	c := &Config{
		c:   config,
		App: &App{},
	}
	if err := config.Scan(c.App); err != nil {
		config.Close()
		return nil, err
	}
	return c, nil
}
