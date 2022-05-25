package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

// Watcher 监听者
type Watcher func(*Bootstrap)

var (
	watchers = make(map[string]Watcher)
)

func NewConfig(filename string) config.Config {
	c := config.New(
		config.WithSource(
			file.NewSource(filename),
		),
	)
	return c
}

func Scan(c config.Config, v interface{}) error {
	return c.Scan(v)
}

// AddWatcher 增加监听者
func AddWatcher(key string, watcher Watcher) {
	watchers[key] = watcher
}

// RemoveWatcher 移除监听者
func RemoveWatcher(key string) {
	delete(watchers, key)
}
