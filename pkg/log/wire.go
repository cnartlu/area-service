// The build tag makes sure the stub is not built in the final build.

package log

import (
	"github.com/google/wire"
)

// ProviderSet 日志的功能
var ProviderSet = wire.NewSet(
// wire.Bind(new(log.Logger), new(*Logger)),
)
