package target

import (
	"path/filepath"

	"github.com/cnartlu/area-service/pkg/utils"
	"gopkg.in/natefinch/lumberjack.v2"
)

type File struct {
	// 文件名称
	Filename string
	// 激活轮询
	EnableRotation *bool
	// 文件最大大小
	MaxFileSize *int
	// 日志文件最大数量
	MaxLogFiles *int
	FileMode    *int
	DirMode     *int

	logger *lumberjack.Logger
}

func (f *File) Init() error {
	f.logger = &lumberjack.Logger{
		Filename:   filepath.Join(utils.RootPath(), "logs", "app.log"),
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}
	return nil
}

func (f *File) Write(p []byte) (int, error) {
	return f.logger.Write(p)
}

func (f *File) Close() error {
	return f.logger.Close()
}
