package storage

import (
	"fmt"
	"strings"
)

// ErrTooManyHandler 太多的请求中间件
var ErrTooManyHandler = fmt.Errorf("too many handlers")

type ErrStorage map[string]*ErrHandlers

func (e ErrStorage) Error() string {
	return ""
}

// Get 获取错误
func (e ErrStorage) Get(key string) *ErrHandlers {
	v := e[key]
	if v == nil {
		return &ErrHandlers{}
	}
	return v
}

// Set 设置错误
func (e *ErrStorage) Set(key string, err *ErrHandlers) {
	key = strings.ToLower(key)
	v := *e
	if _, ok := v[key]; ok {
		v[key] = err
	}
	*e = v
}

// Append 增加错误
func (e *ErrStorage) Append(key string, err error) {
	key = strings.ToLower(key)
	v := *e
	if _, ok := v[key]; ok {
		v[key] = &ErrHandlers{}
	}
	v[key].Append(err)
	*e = v
}

func (e ErrStorage) Len() int {
	return len(e)
}

type ErrHandlers []error

func (e ErrHandlers) Error() string {
	return ""
}

func (e ErrHandlers) String() string {
	return ""
}

func (e *ErrHandlers) Append(err error) {
	if e == nil {
		e = new(ErrHandlers)
	}
	*e = append(*e, err)
}
