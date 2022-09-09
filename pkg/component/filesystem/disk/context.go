package disk

import (
	"fmt"
	"math"
	"net/http"
	"strings"
)

// abortIndex represents a typical value used in abort functions.
const abortIndex int8 = math.MaxInt8 >> 1

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

type errorMsgs []error

func (e errorMsgs) Error() string {
	return e.String()
}

func (a errorMsgs) String() string {
	if len(a) == 0 {
		return ""
	}
	var buffer strings.Builder
	for i, msg := range a {
		fmt.Fprintf(&buffer, "Error #%02d: %s\n", i+1, msg)
	}
	return buffer.String()
}

type Context struct {
	Request *http.Request
	// 错误
	Errors errorMsgs

	// 请求回调
	response *http.Response
	// 是否需要鉴权,默认开启
	withAuthorization bool
	// 请求处理中间件
	handlers HandlersChain
	// 处理索引
	index int8
}

// Next 执行下一个中间件
func (c *Context) Next() *http.Response {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
	return c.response
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

/************************************/
/********* ERROR MANAGEMENT *********/
/************************************/

// Error attaches an error to the current context. The error is pushed to a list of errors.
// It's a good idea to call Error for each error that occurred during the resolution of a request.
// A middleware can be used to collect all the errors and push them to a database together,
// print a log, or append it in the HTTP response.
// Error will panic if err is nil.
func (c *Context) Error(err error) {
	if err == nil {
		panic("err is nil")
	}
	c.Errors = append(c.Errors, err)
	return
}

func (c *Context) WithAuthorization(b bool) *Context {
	c.withAuthorization = b
	return c
}
