package log

import (
	"io"
)

type Target interface {
	io.Writer
	Name() string
	Clone() Target
}

// targets 标签项
var targets = make(map[string]Target)

// Register 注册Target
func Register(t Target) {
	targets[t.Name()] = t
}

// GetTarget 获取Target
func GetTarget(name string) Target {
	i, ok := targets[name]
	if !ok {
		panic("unknown target " + name)
	}
	return i
}
