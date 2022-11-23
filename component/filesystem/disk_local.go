package filesystem

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
)

func init() {
	RegisterTarget(localRegister(0))
}

type localRegister int

func (localRegister) Name() string {
	return "local"
}

func (l localRegister) Register(data map[string]interface{}) (Disk, error) {
	f := LocalDisk{}
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
	return f, nil
}

type LocalResult string

func (l LocalResult) Name() string {
	return "local"
}

func (l LocalResult) Raw() any {
	return string(l)
}

type LocalDisk struct {
	Enable   bool   `json:"enable,omitempty"`
	Root     string `json:"root,omitempty"`
	Domain   string `json:"domain,omitempty"`
	BasePath string `json:"basePath,omitempty"`
}

func (l LocalDisk) IsEnable() bool {
	return l.Enable
}

func (l LocalDisk) Fullname(key string) string {
	return filepath.Join(l.Root, l.BasePath, key)
}

func (l LocalDisk) isFile(filename string) bool {
	f, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

func (l LocalDisk) isDir(dir string) bool {
	f, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return f.IsDir()
}

func (l LocalDisk) Exists(key string, options ...HandleFunc) bool {
	return l.isFile(l.Fullname(key))
}

func (l LocalDisk) Upload(filename, key string, options ...HandleFunc) (Result, error) {
	if !l.IsEnable() {
		return nil, os.ErrInvalid
	}
	if !l.isFile(filename) {
		return nil, os.ErrNotExist
	}
	var fullkey = l.Fullname(key)
	if err := os.Remove(fullkey); err != nil {
		return nil, err
	}
	filedir := filepath.Dir(fullkey)
	if !l.isDir(filedir) {
		if err := os.MkdirAll(filedir, os.ModeDir); err != nil {
			return nil, err
		}
	}
	if err := os.Rename(filename, fullkey); err != nil {
		return nil, err
	}
	return LocalResult("SUCCESS"), nil
}

func (l LocalDisk) Url(key string, options ...HandleFunc) string {
	for _, o := range options {
		_ = o
	}
	buf := bytes.Buffer{}
	buf.WriteString(l.Domain)
	buf.WriteString(key)
	return buf.String()
}

func (l LocalDisk) Delete(key string, options ...HandleFunc) error {
	if l.Exists(key) {
		return os.Remove(l.Fullname(key))
	}
	return nil
}
