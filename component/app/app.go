package app

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	rootPath string
	pwdPath  string
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pwdPath = pwd
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
	rootPath = binDir
	if err := os.Chdir(rootPath); err != nil {
		panic(err)
	}
}

type App struct {
	// Name 应用名称
	name string
	// Env 应用环境
	env EnvName
	// Debug 是否开启调试
	debug bool
	// rootPath 应用根路径
	rootPath string
	// pwdPath 命令行执行路径
	pwdPath string
	// runtimePath 缓存文件路径
	runtimePath string
	// process 进程
	process *process
	// config 原始配置
	config *Config
}

func (a *App) GetName() string {
	if a != nil {
		return a.name
	}
	return ""
}

func (a *App) GetEnv() EnvName {
	if a != nil {
		return a.env
	}
	return EnvName_prod
}

func (a *App) GetDebug() bool {
	if a != nil {
		return a.debug
	}
	return false
}

func (a *App) GetRootPath() string {
	if a != nil {
		return a.rootPath
	}
	return ""
}

func (a *App) GetPwdPath() string {
	if a != nil {
		return a.pwdPath
	}
	return ""
}

func (a *App) GetRuntimePath() string {
	if a != nil {
		return a.runtimePath
	}
	return ""
}

func (a *App) GetProcess() *process {
	if a != nil {
		return a.process
	}
	return nil
}

func (a *App) GetConfig() *Config {
	if a != nil {
		return a.config
	}
	return nil
}

func New(c *Config) *App {
	var name string
	s := filepath.Base(os.Args[0])
	if i := strings.Index(s, "."); i >= 0 {
		name = s[:i]
	} else {
		name = s
	}
	var a = App{
		name:        name,
		debug:       false,
		env:         EnvName_prod,
		config:      c,
		rootPath:    rootPath,
		pwdPath:     pwdPath,
		runtimePath: "runtime",
	}
	if c != nil {
		if c.GetName() != "" {
			a.name = c.GetName()
		}
		a.debug = c.GetDebug()
		a.env = c.GetEnv()
		if c.GetRuntimePath() != "" {
			a.runtimePath = c.GetRuntimePath()
		}
	}
	if !filepath.IsAbs(a.runtimePath) {
		a.runtimePath = filepath.Join(a.rootPath, a.runtimePath)
	}
	if f, err := os.Stat(a.runtimePath); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		if err1 := os.MkdirAll(a.runtimePath, os.ModeDir); err1 != nil {
			panic(err)
		}
	} else if !f.IsDir() {
		panic(os.ErrInvalid)
	}
	a.process = newProcess(&a)
	return &a
}
