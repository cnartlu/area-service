package main

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/component/log"
	"go.uber.org/zap"

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
		// configfile 配置文件路径
		configfile string
	)
	cmd.Flags().StringVarP(&configfile, "config", "c", "config.yaml", "config paths")
	conf, err := config.New(configfile)
	if err != nil {
		log.Error("load config error", zap.Error(err))
		return
	}
	defer conf.Close()
	// 初始化日志器
	logger, err := log.New(
		log.WithConfigure(conf.Config),
		log.WithConfig(conf.Bootstrap.GetLogger()),
	)
	if err != nil {
		log.Error("init logger error", zap.Error(err))
		return
	}
	// 配置执行方法
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		// 初始化
		logger.Info("starting app ...")
		// 监听退出信号
		signalCtx, signalStop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer signalStop()
		appServ, appCleanup, err := initApp(logger, conf)
		if err != nil {
			return err
		}
		defer appCleanup()
		// 调用 app 启动钩子
		if err := appServ.Start(signalStop); err != nil {
			return err
		}
		// 等待退出信号
		<-signalCtx.Done()
		signalStop()

		logger.Info("the app is shutting down ...")

		ctx, cancel := context.WithTimeout(
			context.Background(),
			30*time.Second,
		)
		defer cancel()

		// 关闭应用
		if err := appServ.Stop(ctx); err != nil {
			return err
		}
		return nil
	}

	// command.Setup(cmd, func() (*command.Command, func(), error) {
	// 	return initCommand(loggerWriter, logger, zLogger, configModel)
	// })

	if err := cmd.Execute(); err != nil {
		logger.Error("execute app error", zap.Error(err))
		return
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
