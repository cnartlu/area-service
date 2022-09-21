package storage

import (
	"math"
	"net/http"
)

// abortIndex represents a typical value used in abort functions.
const abortIndex int8 = math.MaxInt8 >> 1

type HandlerFunc func(*Context)

type Context struct {
	// Request 请求对象
	Request *http.Request
	// Error 请求错误
	Errors *ErrHandlers
	// handlers 请求中间件
	handlers []HandlerFunc
	// index 请求中间件索引
	index int8
	// res 请求响应
	res *http.Response
	// authorization 开启鉴权
	authorization *bool
}

// Use 加载请求中间件
func (c *Context) Use(handlers ...HandlerFunc) {
	if c == nil {
		panic("cannot register nil")
	}
	c.handlers = append(c.handlers, handlers...)
}

// Run 执行中间件
func (c *Context) Run() error {
	c.handlers[0](c)
	_ = c.Next()
	return nil
}

// Next 执行下一个中间件
func (c *Context) Next() *http.Response {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
	return c.res
}

// IsAborted returns true if the current context was aborted.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
func (c *Context) Abort() {
	c.index = abortIndex
}

// AbortWithError calls `AbortWithStatus()` and `Error()` internally.
// This method stops the chain, writes the status code and pushes the specified error to `c.Errors`.
// See Context.Error() for more details.
func (c *Context) AbortWithError(err error) {
	c.Abort()
	c.Error(err)
}

// Error attaches an error to the current context. The error is pushed to a list of errors.
// It's a good idea to call Error for each error that occurred during the resolution of a request.
// A middleware can be used to collect all the errors and push them to a database together,
// print a log, or append it in the HTTP response.
// Error will panic if err is nil.
func (c *Context) Error(err error) {
	if err == nil {
		panic("err is nil")
	}
	c.Errors.Append(err)
}

// CheckAbortIndex 检查请求中间件数量
func (c *Context) CheckAbortIndex(number int) error {
	if int8(number) >= abortIndex {
		return ErrTooManyHandler
	}
	return nil
}

// Authorization 权限鉴权开启
func (c *Context) Authorization(b bool) {
	c.authorization = &b
}

// EnableAuthorization 激活鉴权
func (c *Context) EnableAuthorization() bool {
	return !(c.authorization != nil && *c.authorization)
}
