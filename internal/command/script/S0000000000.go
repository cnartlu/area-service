package script

import (
	"github.com/cnartlu/area-service/pkg/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// S0000000000 示例脚本
// 脚本说明：TODO
type S0000000000 struct {
	logger *log.Logger
}

// NewS0000000000 构造函数
func NewS0000000000(logger *log.Logger) *S0000000000 {
	return &S0000000000{
		logger: logger,
	}
}

// Run 脚本执行方法
func (s S0000000000) Run(cmd *cobra.Command, args []string) {
	s.logger.Info("脚本调用成功", zap.String("use", cmd.Use))
}
