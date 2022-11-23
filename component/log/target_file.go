package log

import (
	"io"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	RegisterTarget(fileRegister(0))
}

type fileRegister int8

func (fileRegister) Name() string {
	return "file"
}

func (fileRegister) Register(data map[string]interface{}) (io.Writer, error) {
	f := lumberjack.Logger{}
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		ZeroFields:       true,
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           &f,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(data); err != nil {
		return nil, err
	}
	if f.Filename == "" {
		f.Filename = filepath.Join("logs", "app.log")
	}
	return &f, nil
}
