package log

import (
	"encoding/json"
	"io"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	RegisterTarget(fileRegister(0))
}

type fileRegister int8

func (fileRegister) Name() string {
	return "file"
}

func (fileRegister) Register(data []byte) (io.Writer, error) {
	f := lumberjack.Logger{}
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	if f.Filename == "" {
		f.Filename = filepath.Join("logs", "app.log")
	}
	return &f, nil
}
