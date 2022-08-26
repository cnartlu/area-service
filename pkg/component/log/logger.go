package log

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	// 日志配置
	config kconfig.Config
	// 日志器名称
	name string
	// 日志配置器
	zap *zap.Logger
	// Targets 记录器
	Targets map[string]Target `json:"targets,omitempty"`
	// Stdout 输出到控制台
	Stdout *bool `json:"stdout,omitempty"`
	// TraceLevel 记录堆栈行号
	TraceLevel int `json:"trace_level,omitempty"`
	// Messages 记录固定的消息
	Messages map[string]string `json:"messages,omitempty"`
}

type TargetContent struct {
	Type string `json:"type,omitempty"`
	Target
}

// init 实例化日志
func (l *Logger) init() {
	err := l.config.Value(l.name).Scan(l)
	if err != nil {
		panic(err)
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
	zapOptions = append(zapOptions, zap.AddCaller(), zap.AddCallerSkip(1))
	// 输出到控制台
	if l.Stdout == nil || *l.Stdout {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel))
	}
	// 输出到其他路径
	for _, target := range l.Targets {
		// 查找数据
		l.config.Value(fmt.Sprintf("%s.%s.%s", l.name, "targets", target.Name())).String()

		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(target), zap.DebugLevel))
	}

	core := zapcore.NewTee(cores...)
	l.zap = zap.New(core, zapOptions...)
}

// UnmarshalJSON 实现json反解
func (l *Logger) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Logger
		Targets map[string][]byte `json:"targets,omitempty"`
	}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	l.Stdout = tmp.Stdout
	l.TraceLevel = tmp.TraceLevel
	l.Messages = tmp.Messages
	if tmp.Targets != nil {
		l.Targets = make(map[string]Target)
		for name, bytes := range tmp.Targets {
			var p struct {
				Type string `json:"type,omitempty"`
			}
			err := json.Unmarshal(bytes, &p)
			if err != nil {
				return err
			}
			t := strings.ToLower(strings.TrimSpace(p.Type))
			ta := GetTarget(t)
			err = json.Unmarshal(bytes, &ta)
			if err != nil {
				return err
			}
			l.Targets[name] = ta
		}
	}
	return nil
}

// Use 使用其他的日志
func (l *Logger) Use(name string) *Logger {
	return l
}

// Clone 复制日志
func (l *Logger) Clone() *Logger {
	targets = make(map[string]Target)
	if l.Targets != nil && len(l.Targets) > 0 {
		for k, v := range l.Targets {
			targets[k] = v.Clone()
		}
	}
	stdout := *l.Stdout
	c := &Logger{
		config:   l.config,
		name:     l.name,
		zap:      nil,
		Targets:  targets,
		Stdout:   &stdout,
		Messages: l.Messages,
	}
	c.init()
	return c
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

func (l *Logger) DebugLog(data ...interface{}) {
	clone := l.Clone()
	clone.zap = clone.zap.WithOptions(zap.AddCallerSkip(1))
	clone.Log(log.LevelDebug, data...)
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

// Printf 实现Log日志
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Log(log.LevelInfo, "", fmt.Sprintf(format, v...))
}

func (l *Logger) Sync() error {
	return l.zap.Sync()
}

func (l *Logger) Close() error {
	return l.zap.Sync()
}

// NewLogger 实例化日志器
func NewLogger(config kconfig.Config) *Logger {
	l := &Logger{
		config: config,
		name:   "logger",
	}
	l.init()
	return l
}
