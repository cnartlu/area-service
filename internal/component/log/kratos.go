package log

import (
	"fmt"

	"github.com/cnartlu/area-service/component/log"
	klog "github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

type kratosLogger struct {
	l *log.Logger
}

// Log logs a message at the specified level. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
func (l *kratosLogger) Log(level klog.Level, keyvals ...interface{}) error {
	msg := ""
	data := []zap.Field{}
	switch len(keyvals) {
	case 0:
		return nil
	case 1:
		msg = fmt.Sprint(keyvals[0])
	default:
		if len(keyvals)%2 != 0 {
			msg = fmt.Sprint(keyvals[0])
			for i := 1; i < len(keyvals); i += 2 {
				data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
			}
		} else {
			for i := 0; i < len(keyvals); i += 2 {
				key := fmt.Sprint(keyvals[i])
				switch key {
				case klog.DefaultMessageKey:
					msg = fmt.Sprint(keyvals[i+1])
				default:
					data = append(data, zap.Any(key, keyvals[i+1]))
				}
			}
		}
	}
	switch level {
	case klog.LevelDebug:
		l.l.Debug(msg, data...)
	case klog.LevelInfo:
		l.l.Info(msg, data...)
	case klog.LevelWarn:
		l.l.Warn(msg, data...)
	case klog.LevelError:
		l.l.Error(msg, data...)
	case klog.LevelFatal:
		l.l.Fatal(msg, data...)
	}
	return nil
}

func NewKratosLogger(l *log.Logger) klog.Logger {
	var klogger = kratosLogger{l: l.AddCallerSkip(2)}
	return &klogger
}
