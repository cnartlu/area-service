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
				return true
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

type logger struct {
	once   sync.Once
	config *Config
	core   *zap.Logger
}

func (l *logger) Setup() (err error) {
	l.once.Do(func() {
		if l.config == nil {
			l.config = &Config{
				Stdout: true,
			}
		}
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
		zapOptions = append(zapOptions, zap.AddCaller(), zap.AddCallerSkip(1+int(l.config.GetTraceLevel())))
		defaultLevelEnabler := newLevelEnable(levelEnable{}, l.config.GetLevel(), l.config.GetLevels())
		if l.config.GetStdout() {
			cores = append(cores, zapcore.NewCore(encoder, os.Stdout, defaultLevelEnabler))
		}
		configTargets := l.config.GetTargets()
		protoUnmarshalOption := protojson.UnmarshalOptions{
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
			levelEnabler := newLevelEnable(defaultLevelEnabler, targetConfig.GetLevel(), targetConfig.GetLevels())
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(writer), levelEnabler))
		}
		core := zapcore.NewTee(cores...)
		l.core = zap.New(core, zapOptions...)
	})
	return
}

func New(config *Config) (*zap.Logger, error) {
	l := &logger{
		once:   sync.Once{},
		config: config,
		core:   nil,
	}
	err := l.Setup()
	if err != nil {
		return nil, err
	}
	return l.core, nil
}
