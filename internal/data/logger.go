package data

import (
	klog "github.com/go-kratos/kratos/v2/log"
)

// Logger ent日志处理器
type Logger struct {
	l klog.Logger
}

// DebugLog 实现ent的日志记录器方法
func (l *Logger) DebugLog(keyvals ...interface{}) {
	l.l.Log(klog.LevelDebug, keyvals...)
}
