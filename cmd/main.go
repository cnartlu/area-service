package main

import (
	"os"
	"path/filepath"
	"strings"

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
	// cmd 命令行组件
	cmd = &cobra.Command{
		Version: Version,
		FParseErrWhitelist: cobra.FParseErrWhitelist{
			UnknownFlags: true,
		},
	}
	logger, _ = log.NewLogger(nil)
)

func main() {
	configFilename := ""
	cmd.Flags().StringVarP(&configFilename, "config", "c", "config.yaml", "config paths")
	configure, err := config.NewConfig(configFilename)
	if err != nil {
		logger.Error("load config error", zap.Error(err))
		return
	}
	defer configure.Close()

	// 日志初始化
	{
		l, err := log.NewLogger(configure.GetConfig().GetLogger())
		if err != nil {
			logger.Error("logger init failed", zap.Error(err))
			return
		}
		logger = l
		klog.SetLogger(logger)
	}

	// 配置执行方法
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		appServ, appCleanup, err := initApp(logger, configure)
		if err != nil {
			return err
		}
		defer appCleanup()
		if err := appServ.Start(); err != nil {
			return err
		}
		if err := appServ.Stop(); err != nil {
			return err
		}
		return nil
	}

	// command.Setup(cmd, func() (*command.Command, func(), error) {
	// 	return initCommand(logger, configure)
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
