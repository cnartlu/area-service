package filesystem

import (
	"net/http"
	sync "sync"

	"github.com/cnartlu/area-service/component/proxy"
)

type FileSystem struct {
	once       sync.Once
	disk       multiDisks
	config     *Config
	httpClient *http.Client
}

func (f *FileSystem) Setup() (err error) {
	f.once.Do(func() {
		for field, value := range f.config.GetDisks() {
			value := value
			r := value.GetFields()
			targetType := field
			if t, ok := r["type"]; ok {
				targetType = t.GetStringValue()
			}
			target := GetTarget(targetType)
			if target == nil {
				continue
			}
			disk, err := target.Register(value.AsMap())
			if err != nil {
				continue
			}
			f.disk[field] = disk
		}
	})
	return
}

func (f *FileSystem) Use(name string) Disk {
	if f != nil {
		return f.disk.GetDisk(name)
	}
	return nil
}

func (f *FileSystem) Exists(filename string, options ...HandleFunc) bool {
	if f != nil {
		return f.disk.Exists(filename, options...)
	}
	return false
}

func (f *FileSystem) Upload(filename string, key string, options ...HandleFunc) (Result, error) {
	if f != nil {
		return f.disk.Upload(filename, key, options...)
	}
	return multiResults{}, nil
}

func (f *FileSystem) Url(key string, options ...HandleFunc) string {
	if f != nil {
		return f.disk.Url(key, options...)
	}
	return ""
}

func (f *FileSystem) Delete(key string, options ...HandleFunc) error {
	if f != nil {
		return f.disk.Delete(key, options...)
	}
	return nil
}

func New(config *Config, p *proxy.Client) *FileSystem {
	var result = &FileSystem{
		config:     config,
		httpClient: http.DefaultClient,
	}
	if p != nil {
		result.httpClient = p.Client
	}
	return result
}
