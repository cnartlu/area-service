package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	targets map[string]io.WriteCloser
	// 日志配置器
	zap *zap.Logger
}

// new 实例化组件
func (l *Logger) new() (err error) {
	l.once.Do(func() {
		if l.c == nil {
			l.c = &Config{}
		}
		for field, value := range l.c.GetTargets() {
			r := value.GetFields()
			targetType := field
			if t, ok := r["type"]; ok {
				targetType = t.String()
			}
			// 检查类型是否被注册
			target := GetTarget(targetType)
			if target == nil {
				err = fmt.Errorf("logger target not registered, type is \"%s\"", field)
				return
			}
			b, _ := value.MarshalJSON()
			err = json.Unmarshal(b, target)
			if err != nil {
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
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
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
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.zap.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case log.LevelDebug:
		l.zap.Debug("", data...)
	case log.LevelInfo:
		l.zap.Info("", data...)
	case log.LevelWarn:
		l.zap.Warn("", data...)
	case log.LevelError:
		l.zap.Error("", data...)
	case log.LevelFatal:
		l.zap.Fatal("", data...)
	}
	return nil
}

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
