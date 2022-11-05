package log

import (
	"os"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	klog "github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var packagepath = "github.com/cnartlu/area-service/"

type Logger struct {
	// 是否初始化
	once sync.Once
	// c 配置装置
	c *Config
	// 日志配置器
	zap *zap.Logger
}

func (l *Logger) core() *zap.Logger {
	if l == nil {
		l = &Logger{}
	}
	if l.zap == nil {
		l.Setup()
	}
	return l.zap
}

// clone
func (l *Logger) clone() *Logger {
	c := *l
	return &c
}

func (l *Logger) Setup() (err error) {
	l.once.Do(func() {
		var (
			encoderConfig     zapcore.EncoderConfig
			encoder           zapcore.Encoder
			cores             []zapcore.Core
			zapOptions        []zap.Option
			coverLogLevelFunc = func(s string) zapcore.Level {
				var defaultLevel zapcore.Level
				s = strings.TrimSpace(s)
				if s == "" {
					return defaultLevel
				}
				switch strings.ToLower(s) {
				case "info":
					defaultLevel = zap.DebugLevel
				case "warn", "warning":
					defaultLevel = zap.DebugLevel
				case "error":
					defaultLevel = zap.DebugLevel
				case "debug":
					fallthrough
				default:
					defaultLevel = zap.DebugLevel
				}
				return defaultLevel
			}
		)
		// 解析配置
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.MessageKey = klog.DefaultMessageKey
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		if packagepath != "" {
			encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
				s := caller.FullPath()
				s = strings.TrimPrefix(s, packagepath)
				enc.AppendString(s)
			}
		}
		encoder = zapcore.NewJSONEncoder(encoderConfig)
		// zap选项
		zapOptions = append(zapOptions, zap.AddCaller(), zap.AddCallerSkip(1+int(l.c.GetTraceLevel())))
		// 默认日志级别
		defaultLevel := coverLogLevelFunc(l.c.GetLogLevel())
		// 输出到控制台
		if l.c == nil || l.c.Stdout == nil || *l.c.Stdout {
			cores = append(cores, zapcore.NewCore(encoder, os.Stdout, defaultLevel))
		}
		// 使用其他的输出项输出
		{
			configTargets := l.c.GetTargets()
			for field, value := range configTargets {
				value := value
				r := value.GetFields()
				targetType := field
				if t, ok := r["type"]; ok {
					targetType = t.GetStringValue()
				}
				// 检查类型是否被注册
				target := GetTarget(targetType)
				if target == nil {
					buf := strings.Builder{}
					buf.WriteString("logger target name [")
					buf.WriteString(field)
					buf.WriteString("] not registered, type is [")
					buf.WriteString(targetType)
					buf.WriteString("]")
					log.Log(klog.LevelDebug, buf.String())
					err = nil
					continue
				}
				writer, err := target.Register(value.AsMap())
				if err != nil {
					buf := strings.Builder{}
					buf.WriteString("logger target name [")
					buf.WriteString(field)
					buf.WriteString("] not registered, type is [")
					buf.WriteString(targetType)
					buf.WriteString("]")
					log.Log(klog.LevelDebug, buf.String())
					err = nil
					continue
				}
				// 该target适用的日志级别
				level := defaultLevel.String()
				if t, ok := r["level"]; ok {
					level = t.GetStringValue()
				}
				targetLevel := coverLogLevelFunc(level)
				// 核心
				cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(writer), targetLevel))
			}

		}
		core := zapcore.NewTee(cores...)
		l.zap = zap.New(core, zapOptions...)
	})
	return
}

// AddCallerSkip 增加过滤长度
func (l *Logger) AddCallerSkip(skip int) *Logger {
	c := l.clone()
	c.zap = l.core().WithOptions(zap.AddCallerSkip(skip))
	return c
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Debug(msg string, fields ...zapcore.Field) {
	l.core().Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Info(msg string, fields ...zapcore.Field) {
	l.core().Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Warn(msg string, fields ...zapcore.Field) {
	l.core().Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Error(msg string, fields ...zapcore.Field) {
	l.core().Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func (l *Logger) DPanic(msg string, fields ...zapcore.Field) {
	l.core().DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (l *Logger) Panic(msg string, fields ...zapcore.Field) {
	l.core().Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (l *Logger) Fatal(msg string, fields ...zapcore.Field) {
	l.core().Fatal(msg, fields...)
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func (l *Logger) Sync() error {
	return l.core().Sync()
}

// Close calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func (l *Logger) Close() error {
	return l.Sync()
}

// New 日志器
func New(c *Config) (*Logger, error) {
	l := &Logger{
		once: sync.Once{},
		c:    c,
		zap:  nil,
	}
	err := l.Setup()
	if err != nil {
		return nil, err
	}
	return l, nil
}
