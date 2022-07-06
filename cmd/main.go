package main

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/cnartlu/area-service/internal/component/ent"
	appconfig "github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/component/log"
	kconfig "github.com/go-kratos/kratos/v2/config"
	kconfigFile "github.com/go-kratos/kratos/v2/config/file"

	"github.com/spf13/cobra"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string = "0.0.1beta"
	// cmd 命令行组件
	cmd = &cobra.Command{
		Version: Version,
		FParseErrWhitelist: cobra.FParseErrWhitelist{
			UnknownFlags: true,
		},
	}
)

func main() {
	var (
		// configPaths 配置文件路径
		configPaths []string
	)
	cmd.Flags().StringSliceVarP(&configPaths, "config", "c", []string{"config.yaml"}, "config paths")
	// 初始化配置器
	var options []kconfig.Option
	for _, path := range configPaths {
		options = append(options, kconfig.WithSource(kconfigFile.NewSource(path)))
	}
	config := kconfig.New(options...)
	if err := config.Load(); err != nil {
		panic(err)
	}
	defer config.Close()
	var bootstrap appconfig.Bootstrap
	if err := config.Scan(bootstrap); err != nil {
		panic(err)
	}
	// 初始化日志器
	logger, err := log.New(
		log.WithConfigure(config),
		log.WithConfig(bootstrap.GetLogger()),
	)
	if err != nil {
		panic(err)
	}
	ent.NewClient()
	// 配置执行方法
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		// 初始化
		logger.Info("starting app ...")
		// 监听退出信号
		signalCtx, signalStop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer signalStop()
		appServ, appCleanup, err := initApp(logger, &bootstrap)
		if err != nil {
			return err
		}
		defer appCleanup()
		// 调用 app 启动钩子
		if err := appServ.Start(signalStop); err != nil {
			panic(err)
		}
		// 等待退出信号
		<-signalCtx.Done()
		signalStop()

		logger.Info("the app is shutting down ...")

		// time.Duration(bootstrap.Application.Timeout)
		ctx, cancel := context.WithTimeout(
			context.Background(),
			30*time.Second,
		)
		defer cancel()

		// 关闭应用
		if err := appServ.Stop(ctx); err != nil {
			panic(err)
		}
		return nil
	}

	// command.Setup(cmd, func() (*command.Command, func(), error) {
	// 	return initCommand(loggerWriter, logger, zLogger, configModel)
	// })

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

// getExecutableName 获取执行命令的名称
func getExecutableName() string {
	binPath, err := os.Executable()
	if err != nil {
		return ""
	}
	filenameWithSuffix := filepath.Base(binPath)
	return strings.TrimSuffix(filenameWithSuffix, "."+filepath.Ext(binPath))
}

func NewConfig() {

}
