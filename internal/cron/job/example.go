package job

import (
	"github.com/cnartlu/area-service/pkg/log"
)

// Example 示例任务
// 任务说明：TODO
type Example struct {
	logger *log.Logger
}

// NewExample 构造函数
func NewExample(logger *log.Logger) *Example {
	return &Example{
		logger: logger,
	}
}

// Run 任务执行方法
func (s Example) Run() {
	s.logger.Info("Example 任务执行成功")
}
