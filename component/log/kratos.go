package log

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

type kratosLogger struct {
	l *Logger
}

// Log logs a message at the specified level. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
func (l *kratosLogger) Log(level log.Level, keyvals ...interface{}) error {
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
				case log.DefaultMessageKey:
					msg = fmt.Sprint(keyvals[i+1])
				default:
					data = append(data, zap.Any(key, keyvals[i+1]))
				}
			}
		}
	}
	switch level {
	case log.LevelDebug:
		l.l.Debug(msg, data...)
	case log.LevelInfo:
		l.l.Info(msg, data...)
	case log.LevelWarn:
		l.l.Warn(msg, data...)
	case log.LevelError:
		l.l.Error(msg, data...)
	case log.LevelFatal:
		l.l.Fatal(msg, data...)
	}
	return nil
}

func NewKratosLogger(l *Logger) log.Logger {
	var klogger = kratosLogger{l: l.AddCallerSkip(1)}
	return &klogger
}
