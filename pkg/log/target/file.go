package target

import (
	"encoding/json"
	"path/filepath"
	"time"

	"github.com/cnartlu/area-service/pkg/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	log.RegisterTarget(NewFile())
}

type File struct {
	// 文件名称
	Filename string
	// 激活轮询
	EnableRotation bool
	// 文件最大大小
	MaxFileSize int
	// 日志文件最大数量
	MaxLogFiles int
	FileMode    int
	DirMode     int

	logger *lumberjack.Logger
}

func (f *File) Name() string {
	return "file"
}

func (f *File) UnmarshalJSON(data []byte) error {
	if f == nil {
		*f = File{}
	}
	type a *File
	b := a(f)
	err := json.Unmarshal(data, &b)
	if err != nil {
		return err
	}
	*f = File(*b)
	return nil
}

func (f *File) Write(p []byte) (int, error) {
	return f.logger.Write(p)
}

func (f *File) Close() error {
	return f.logger.Close()
}

func NewFile() *File {
	return &File{
		logger: &lumberjack.Logger{
			Filename:   filepath.Join("logs", time.Now().Format("log.20060102")+".log"),
			MaxSize:    50,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		},
	}
}
