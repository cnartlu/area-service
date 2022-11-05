package app

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	pkgpath "github.com/cnartlu/area-service/pkg/path"
	kconfig "github.com/go-kratos/kratos/v2/config"
	kconfigFile "github.com/go-kratos/kratos/v2/config/file"
	klog "github.com/go-kratos/kratos/v2/log"
)

type App struct {
	mux sync.Mutex
	// name 应用名称
	name string
	// version 版本号
	version string
	// pid 文件
	pid *Pid
	// pwd
	pwd string
	// config 配置
	config kconfig.Config
	// flag 解析值
	flag Flag
	// start 执行启动函数
	start func() error
}

// Name 应用名称
func (a *App) Name() string {
	if a.name == "" {
		s := filepath.Base(os.Args[0])
		if i := strings.Index(s, "."); i >= 0 {
			a.name = s[:i]
		} else {
			a.name = s
		}
	}
	return a.name
}

// Version returns the version
func (a *App) Version() string {
	return a.version
}

// Config returns the config
func (a *App) Config() kconfig.Config {
	return a.config
}

// WithOptions 设置Options
func (a *App) WithOptions(options ...Option) {
	for _, o := range options {
		o(a)
	}
}

// Run the application
func (a *App) Run() (err error) {
	if a.flag.Test || a.flag.Help || a.flag.Version {
		return
	}
	if a.flag.Signal != "" {
		switch a.flag.Signal {
		case "quit":
			process, err := a.pid.Process()
			if err != nil {
				return err
			}
			return process.Signal(os.Interrupt)
		case "stop":
			process, err := a.pid.Process()
			if err != nil {
				return err
			}
			return process.Kill()
		case "reload":
			process, err := a.pid.Process()
			if err != nil {
				return err
			}
			return process.Signal(syscall.SIGHUP)
		default:
			buf := strings.Builder{}
			buf.WriteString("nginx")
			buf.WriteString(": invalid option: \"-")
			buf.WriteString("s ")
			buf.WriteString(a.flag.Signal)
			buf.WriteString("\"")
			return errors.New(buf.String())
		}
	}
	// 初始化配置
	var sources []kconfig.Source
	var filenames []string = []string{
		a.flag.Config,
		filepath.Join("../etc", a.flag.Config),
		"config.yaml",
		"config.json",
	}
	for _, filename := range filenames {
		if err := isFile(filename); err == nil {
			sources = append(sources, kconfigFile.NewSource(filename))
		}
	}
	logger := klog.NewFilter(klog.DefaultLogger, klog.FilterLevel(klog.LevelError))
	klog.SetLogger(logger)
	a.config = kconfig.New(
		kconfig.WithLogger(logger),
		kconfig.WithSource(sources...),
	)
	if err := a.config.Load(); err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	// 初始化PID
	a.pid = NewPid(a.Name(), a.Config())
	// 设置pid文件
	if err := a.pid.WriteFile(); err != nil {
		return err
	}
	defer func() {
		if err := a.pid.Remove(); err != nil {
			os.Stdout.WriteString(err.Error())
		}
	}()
	// 监听信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			sig := <-c
			switch sig {
			case syscall.SIGHUP:
				// 1、将当前pid文件迁移到 old
				if err := a.pid.Move(); err != nil {
					// 命令行输出错误信息，继续监听
					fmt.Println("Error: ", err)
					continue
				}
				// 获取当前进程的socket数据

				cmd := exec.Command(os.Args[0], os.Args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				err := cmd.Start()
				if err != nil {
					err1 := a.pid.Recover()
					fmt.Println("Error:", err)
					fmt.Println("remove pid file failed:", err1)
					continue
				}
				// 等待子进程发出信号停止
			default:
				// ingrone ...
				return
			}
		}
	}()
	if a.start != nil {
		return a.start()
	}
	return nil
}

// New returns app
func New(options ...Option) *App {
	wd, _ := os.Getwd()
	v := &App{
		pwd: wd,
	}
	os.Chdir(pkgpath.RootPath())
	for _, fun := range options {
		fun(v)
	}
	return v
}
