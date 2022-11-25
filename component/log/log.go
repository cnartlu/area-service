package log

import (
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/encoding/protojson"
)

var packagepath = "github.com/cnartlu/area-service/"

type levelEnable struct {
	level  zapcore.Level
	levels []zapcore.Level
}

func (l levelEnable) Enabled(lvl zapcore.Level) bool {
	if len(l.levels) > 0 {
		for _, level := range l.levels {
			if level == lvl {
				return lvl >= l.level
			}
		}
		return false
	}
	return lvl >= l.level
}

func newLevelEnable(oldLevelEnable levelEnable, level Level, levels []Level) levelEnable {
	var newLevelEnable = levelEnable{
		level:  oldLevelEnable.level,
		levels: oldLevelEnable.levels,
	}
	switch level {
	case Level_debug:
		newLevelEnable.level = zap.DebugLevel
	case Level_info:
		newLevelEnable.level = zap.InfoLevel
	case Level_warning:
		newLevelEnable.level = zap.WarnLevel
	case Level_error:
		newLevelEnable.level = zap.ErrorLevel
	}
	if len(levels) > 0 {
		var levels_bak = make([]zapcore.Level, len(levels))
		for k, level := range levels {
			var breakFor = false
			switch level {
			case Level_debug:
				levels_bak[k] = zap.DebugLevel
			case Level_info:
				levels_bak[k] = zap.InfoLevel
			case Level_warning:
				levels_bak[k] = zap.WarnLevel
			case Level_error:
				levels_bak[k] = zap.ErrorLevel
			case Level_all:
				levels_bak = make([]zapcore.Level, len(levels))
				breakFor = true
			}
			if breakFor {
				break
			}
		}
		newLevelEnable.levels = levels_bak
	}
	return newLevelEnable
}

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
			encoderConfig zapcore.EncoderConfig
			encoder       zapcore.Encoder
			cores         []zapcore.Core
			zapOptions    []zap.Option
		)
		encoderConfig = zap.NewProductionEncoderConfig()
		// encoderConfig.MessageKey = klog.DefaultMessageKey
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
		zapOptions = append(zapOptions, zap.AddCaller(), zap.AddCallerSkip(1+int(l.c.GetTraceLevel())))
		defaultLevel := levelEnable{}
		defaultLevel = newLevelEnable(defaultLevel, l.c.GetLevel(), l.c.GetLevels())
		if l.c == nil || l.c.Stdout == nil || *l.c.Stdout {
			cores = append(cores, zapcore.NewCore(encoder, os.Stdout, defaultLevel))
		}
		configTargets := l.c.GetTargets()
		protoUnmarshalOption := protojson.UnmarshalOptions{
			// RecursionLimit: 1,
			DiscardUnknown: true,
		}
		for field, targetStruct := range configTargets {
			targetStruct := targetStruct
			var targetConfig TargetConfig
			targetStructBytes, _ := targetStruct.MarshalJSON()
			if err1 := protoUnmarshalOption.Unmarshal(targetStructBytes, &targetConfig); err1 != nil {
				err = err1
				return
			}
			targetType := targetConfig.GetType()
			if targetType == "" {
				targetType = field
			}
			targetType = strings.ToLower(targetType)
			target := GetTarget(targetType)
			if target == nil {
				target = GetTarget("file")
			}
			writer, err1 := target.Register(targetStructBytes)
			if err1 != nil {
				err = err1
				break
			}
			var targetLevel = newLevelEnable(defaultLevel, targetConfig.GetLevel(), targetConfig.GetLevels())
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(writer), targetLevel))
		}
		core := zapcore.NewTee(cores...)
		l.zap = zap.New(core, zapOptions...)
	})
	return
}

// AddCallerSkip increases the number of callers skipped by caller annotation
// (as enabled by the AddCaller option). When building wrappers around the
// Logger and SugaredLogger, supplying this Option prevents zap from always
// reporting the wrapper code as the caller.
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
