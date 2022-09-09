package app

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	kconfig "github.com/go-kratos/kratos/v2/config"
)

type App struct {
	// Name 应用名称
	Name string
	// ProxyUrl 请求代理
	ProxyUrl string
	// Env 应用环境
	Env string
	// Debug 开启调试
	Debug bool

	// ExecPath 执行目录相对路径
	execPath string
	// RuntimePath 缓存文件路径
	runtimePath string

	// 配置器
	config kconfig.Config
}

// Reset 重置app
func (a *App) Reset() *App {
	*a = App{
		config: a.config,
	}
	return a
}

// RootPath 获取此项目的绝对路径
// 如果是以 go build 生成的二进制文件运行，则返回 二进制文件的目录路径
// 如果是以 go run 运行，则返回在此项目的绝对路径
func (a *App) RootPath() string {
	if a.execPath != "" {
		return a.execPath
	}
	var binDir string
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	binDir = filepath.Dir(exePath)
	tmpDir := os.TempDir()
	if strings.Contains(exePath, tmpDir) {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			binDir = filepath.Dir(filepath.Dir(filepath.Dir(filename)))
		}
	}
	a.execPath = binDir
	return binDir
}

// RuntimePath 获取缓存数据的绝对路径
func (a *App) RuntimePath() string {
	if a.runtimePath != "" {
		return a.runtimePath
	}
	runtimePath := filepath.Join(a.RootPath(), "runtime")
	a.runtimePath = runtimePath
	return runtimePath
}

func New(config kconfig.Config) *App {
	c := Config{}
	err := config.Scan(&c)
	if err != nil {
		panic(err)
	}
	env := "prod"
	switch strings.ToLower(strings.TrimSpace(c.Env)) {
	case "dev":
		env = "dev"
	case "uat":
		env = "uat"
	case "test":
		env = "test"
	case "prod":
		fallthrough
	default:
		env = "prod"
	}
	debug := true
	if c.GetDebug() != nil {
		debug = c.Debug.GetValue()
	}
	app := &App{
		config:      config,
		Name:        c.Name,
		ProxyUrl:    c.ProxyUrl,
		Env:         env,
		Debug:       debug,
		runtimePath: c.RuntimePath,
	}
	config.Watch("app", func(k string, v kconfig.Value) {

	})
	return app
}
