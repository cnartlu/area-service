package log

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	// 是否初始化
	once sync.Once
	// c 配置装置
	c *Config
	// targets 记录目标
	targets map[string]Target
	// 日志配置器
	zap *zap.Logger
}

// new 实例化组件
func (l *Logger) new() (err error) {
	l.once.Do(func() {
		if l.c == nil {
			l.c = &Config{}
		}
		configTargets := l.c.GetTargets()
		l.targets = make(map[string]Target, len(configTargets))
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
				err = fmt.Errorf("logger target name [%s] not registered, type is [%s] ", field, targetType)
				return
			}
			b, _ := value.MarshalJSON()
			if err := target.UnmarshalJSON(b); err != nil {
				return
			}
			l.targets[field] = target
		}
		var (
			encoderConfig zapcore.EncoderConfig
			encoder       zapcore.Encoder
			cores         []zapcore.Core
			zapOptions    []zap.Option
		)
		// 解析配置
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.MessageKey = log.DefaultMessageKey
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			s := caller.FullPath()
			// 清除项目mod路径
			s = strings.TrimPrefix(s, "github.com/cnartlu/area-service/")
			enc.AppendString(s)
		}
		encoder = zapcore.NewJSONEncoder(encoderConfig)
		// zap选项
		zapOptions = append(zapOptions, zap.AddCaller(), zap.AddCallerSkip(1+int(l.c.GetTraceLevel())))
		// 输出到控制台
		if l.c.Stdout == nil || *l.c.Stdout {
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel))
		}
		// 输出到其他路径
		for _, target := range l.targets {
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(target), zap.DebugLevel))
		}
		core := zapcore.NewTee(cores...)
		l.zap = zap.New(core, zapOptions...)
	})
	return
}

// 实现日志器的方法
func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
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
	l = l.AddCallerSkip(1)
	switch level {
	case log.LevelDebug:
		l.zap.Debug(msg, data...)
	case log.LevelInfo:
		l.zap.Info(msg, data...)
	case log.LevelWarn:
		l.zap.Warn(msg, data...)
	case log.LevelError:
		l.zap.Error(msg, data...)
	case log.LevelFatal:
		l.zap.Fatal(msg, data...)
	}
	return nil
}

// AddCallerSkip 增加过滤长度
func (l *Logger) AddCallerSkip(skip int) *Logger {
	logger := &Logger{
		once:    l.once,
		targets: l.targets,
		c:       l.c,
		zap:     nil,
	}
	logger.zap = l.zap.WithOptions(zap.AddCallerSkip(skip))
	return logger
}

// Debug 记录debug日志
func (l *Logger) Debug(msg string, fields ...zapcore.Field) {
	l.zap.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zapcore.Field) {
	l.zap.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zapcore.Field) {
	l.zap.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zapcore.Field) {
	l.zap.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...zapcore.Field) {
	l.zap.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zapcore.Field) {
	l.zap.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zapcore.Field) {
	l.zap.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.zap.Sync()
}

func (l *Logger) Close() error {
	return l.zap.Sync()
}

// NewLogger 日志器
func NewLogger(c *Config) (*Logger, error) {
	l := &Logger{
		c: c,
	}
	err := l.new()
	if err != nil {
		return nil, err
	}
	return l, nil
}
