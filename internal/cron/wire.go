// The build tag makes sure the stub is not built in the final build.

package cron

import (
	"github.com/cnartlu/area-service/internal/cron/job"
	"github.com/google/wire"
)

// ProviderSet 命令行注入方法
var ProviderSet = wire.NewSet(
	job.NewGithub,
	New,
)
