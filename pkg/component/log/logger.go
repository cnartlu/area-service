package log

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cnartlu/area-service/pkg/path"
	ppath "github.com/cnartlu/area-service/pkg/path"
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	configure kconfig.Config
	config    *Config
	log       *zap.Logger
	// 日志配置器
	loggers map[string]*zap.Logger
}

func (l *Logger) Use(name string) *Logger {
	adefault := l.config.Default
	logger := Logger{
		configure: l.configure,
		config:    l.config,
		log:       l.Zap(name),
		loggers: map[string]*zap.Logger{
			adefault: l.Zap(name),
		},
	}
	return &logger
}

// Zap 获取日志配置器
func (l *Logger) Zap(name string) *zap.Logger {
	if name == "" {
		return l.log
	} else if logger, ok := l.loggers[name]; ok {
		return logger
	}
	return l.log
}

// 实现日志器的方法
func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	case log.LevelFatal:
		l.log.Fatal("", data...)
	}
	return nil
}

func (l *Logger) DebugLog(data ...interface{}) {
	zap := l.log.WithOptions(zap.AddCallerSkip(1))
	clone := &Logger{log: zap}
	clone.Log(log.LevelDebug, data...)
}

func (l *Logger) Debug(msg string, fields ...zapcore.Field) {
	l.log.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zapcore.Field) {
	l.log.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zapcore.Field) {
	l.log.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zapcore.Field) {
	l.log.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...zapcore.Field) {
	l.log.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zapcore.Field) {
	l.log.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zapcore.Field) {
	l.log.Fatal(msg, fields...)
}

// Printf 实现Log日志
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Log(log.LevelInfo, "", fmt.Sprintf(format, v...))
}

func (l *Logger) Sync() error {
	return l.log.Sync()
}

func (l *Logger) Close() error {
	return l.log.Sync()
}

func New(options ...Option) (*Logger, error) {
	logger := Logger{}
	for _, option := range options {
		option(&logger)
	}
	// 初始化配置
	if logger.configure != nil {
		// 监听配置文件的修改
		logger.configure.Watch("logger", func(key string, v kconfig.Value) {

		})
	}
	// 实例化zap日志
	logger.log = newLogger(nil)
	if logger.config != nil {
		config := logger.config
		if config.Loggers != nil {
			for name, loggerConfig := range config.Loggers {
				zapLogger := newLogger(loggerConfig)
				logger.loggers[name] = zapLogger
			}
		}
		// 默认配置
		if config.Default != "" {
			if defaultLogger, ok := logger.loggers[config.Default]; ok {
				logger.log = defaultLogger
			}
		}
	}

	return &logger, nil
}

func NewDefault() (*Logger, error) {
	return New()
}

func newLogger(c *Config_Logger) *zap.Logger {
	var (
		encoderConfig zapcore.EncoderConfig
		encoder       zapcore.Encoder
		writeSyncer   zapcore.WriteSyncer
		core          zapcore.Core
		zapOptions    []zap.Option
		zapFields     []zapcore.Field
	)
	encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoder = zapcore.NewJSONEncoder(encoderConfig)
	writeSyncer = zapcore.AddSync(os.Stdout)
	zapOptions = append(zapOptions, zap.AddCaller(), zap.AddCallerSkip(1))

	if c != nil {
		if c.Path != "" {
			c.Path = ppath.RootPath()
			if c.File == "" {
				c.File = "{Y-m-d}.log"
			}
			writeSyncer = zapcore.AddSync(os.Stdout)
		}
		if c.Stdout == nil || c.Stdout.GetValue() {
			writeSyncer = zapcore.AddSync(os.Stdout)
		}
		if c.Header == nil || c.Header.GetValue() {
			zapFields = append(zapFields, zap.Any("headers", ""))
		}
		if len(c.CtxKeys) > 0 {
			for k, v := range c.CtxKeys {
				zapFields = append(zapFields, zap.Any(k, v.GetValue()))
			}
		}
	}

	_ = os.MkdirAll(filepath.Join(path.RootPath(), "logs"), 0666)
	core = zapcore.NewTee(
		zapcore.NewCore(
			encoder,
			writeSyncer,
			zap.DebugLevel,
		),
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   filepath.Join(path.RootPath(), "logs", "app.log"),
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     28,   //days
				Compress:   true, // disabled by default
			}),
			zap.DebugLevel,
		),
	)

	core.With(zapFields)
	zap := zap.New(core, zapOptions...)
	return zap
}
