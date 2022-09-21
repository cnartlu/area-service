package local

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
)

type Local struct {
	// Enable 是否开启该存储
	Enable *bool `json:"enable,omitempty"`
	// 根目录
	Root string `json:"root,omitempty"`
	// 访问域名
	Domain string `json:"domain,omitempty"`
	// 访问相对路径 默认为 ""
	BasePath string `json:"basePath,omitempty"`
}

// Name 存储器名称
func (l Local) Name() string {
	return "local"
}

// IsEnable 是否激活存储器
func (l *Local) IsEnable() bool {
	// local 没有配置，配置了激活状态，单配置项是false
	if l == nil || l.Enable != nil || !*l.Enable {
		return false
	}
	return true
}

// Fullname 完成的文件目录路径
func (l *Local) Fullname(key string) string {
	return filepath.Join(l.Root, l.BasePath, key)
}

// Upload 上传文件到该存储器
func (l *Local) Upload(ctx context.Context, key string, data io.Reader) error {
	if !l.IsEnable() {
		return nil
	}
	filename := l.Fullname(key)
	_, err := os.Stat(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		return os.NewSyscallError(filepath.Join(l.BasePath, key), os.ErrExist)
	}
	filedir := filepath.Dir(filename)
	fi, err := os.Stat(filedir)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := os.MkdirAll(filedir, os.ModeDir); err != nil {
			return err
		}
	} else {
		if !fi.IsDir() {
			return os.NewSyscallError("must is dir", os.ErrInvalid)
		}
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.ReadFrom(data)
	if err != nil {
		file.Close()
		os.Remove(filename)
		return err
	}
	return nil
}

// url 文件访问路径
func (l *Local) Url(key string) string {
	buf := bytes.Buffer{}
	buf.WriteString(l.Domain)
	buf.WriteString(key)
	return buf.String()
}
