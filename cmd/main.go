package main

import (
	"github.com/cnartlu/area-service/internal/command"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/log"
	klog "github.com/go-kratos/kratos/v2/log"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string = "0.0.1beta"
	// Config 配置文件
	Config string
	// logger 配置器
	logger, _ = log.NewLogger(nil)
	// configure 配置器
	configure *config.Config
	// cmd 命令行组件
	cmd = &cobra.Command{
		Use:           "",
		Short:         ``,
		Long:          ``,
		Version:       Version,
		SilenceErrors: true,
		SilenceUsage:  true,
		FParseErrWhitelist: cobra.FParseErrWhitelist{
			UnknownFlags: true,
		},

		PreRunE: func(cmd *cobra.Command, args []string) error {
			logger.Debug("", zap.Strings("args", args))
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			appServ, appCleanup, err := initApp(logger, configure)
			if err != nil {
				return err
			}
			defer appCleanup()
			return appServ.Start()
		},
	}
)

func main() {
	{
		cmd.Flags().StringVarP(&Config, "config", "c", "config.yaml", "config paths")
		c, err := config.NewConfig(Config)
		if err != nil {
			logger.Error("load config error", zap.Error(err))
			return
		}
		configure = c
		defer configure.Close()
	}

	// 日志初始化
	{
		l, err := log.NewLogger(configure.App.GetLogger())
		if err != nil {
			logger.Error("logger init failed", zap.Error(err))
			return
		}
		logger = l
		klog.SetLogger(logger)
	}

	command.Setup(cmd, func() (*command.Command, func(), error) {
		return initCommand(logger, configure)
	})

	if err := cmd.Execute(); err != nil {
		logger.Error("execute app error", zap.Error(err))
	}
}
