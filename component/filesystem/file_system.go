package filesystem

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	kconfig "github.com/go-kratos/kratos/v2/config"

	"context"
)

type FileSystem struct {
	Default string   `json:"default,omitempty"`
	Stores  []string `json:"stores,omitempty"`
}

func (f *FileSystem) Upload(ctx context.Context, key string, body io.Reader, hooks ...HookFunc) (result *UploadResult, err error) {
	key = ""
	length := len(f.Stores)
	switch length {
	case 0:
		return nil, fmt.Errorf("store is empty")
	case 1:
		for _, name := range f.Stores {
			storage := GetStorage(name)
			err := storage.Upload(ctx, key, body)
			if err != nil {
				return nil, err
			}
		}
	default:
		// 判断文件打大小以验证是否需要开启临时存在在本地磁盘
		// 如果开启临时存储，则需要
		file, err := ioutil.TempFile(os.TempDir(), "go_*")
		if err != nil {
			return nil, err
		}
		defer func() {
			file.Close()
			os.Remove(file.Name())
		}()
		_, err = file.ReadFrom(body)
		if err != nil {
			return nil, err
		}
		for _, name := range f.Stores {
			// 指针转到初始位置
			_, err := file.Seek(0, 0)
			if err != nil {
				continue
			}
			storage := GetStorage(name)
			err = storage.Upload(ctx, key, body)
			if err != nil {
				return nil, err
			}
		}
	}
	return
}

func (f *FileSystem) Url(filename string, hooks ...HookFunc) {

}

func New(config kconfig.Config) *FileSystem {
	v := &FileSystem{}
	err := config.Value("file_system").Scan(v)
	if err != nil {
		panic(err)
	}
	return v
}
