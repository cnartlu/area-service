package storage

import "net/http"

// SetResponse 设置响应请求
func SetResponse(res *http.Response) HandlerFunc {
	return func(c *Context) {
		c.res = res
	}
}
