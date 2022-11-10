package log

import (
	"fmt"

	"github.com/cnartlu/area-service/component/log"
	"go.uber.org/zap"
)

type Ent struct {
	*log.Logger
}

// DebugLog 实现ent的日志记录器方法
func (l *Ent) DebugLog(keyvals ...interface{}) {
	length := len(keyvals)
	switch length {
	case 0:
	case 1:
		l.AddCallerSkip(1).Debug(fmt.Sprint(keyvals[0]))
	default:
		var (
			msg  string
			data []zap.Field
		)
		if length%2 == 0 {
			for i := 0; i < len(keyvals); i += 2 {
				data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
			}
		} else {
			for i := 1; i < len(keyvals); i += 2 {
				data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
			}
		}
		l.AddCallerSkip(1).Debug(msg, data...)
	}
}

func NewEnt(l *log.Logger) *Ent {
	return &Ent{
		Logger: l,
	}
}
