package middleware

import (
	"net"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	errors "github.com/go-kratos/kratos/v2/errors"
	"go.uber.org/zap"
)

func Recover(logger *zap.Logger) gin.HandlerFunc {
	l := logger.WithOptions(zap.AddCallerSkip(1))
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if !brokenPipe {
					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					headers := strings.Split(string(httpRequest), "\r\n")
					for idx, header := range headers {
						current := strings.Split(header, ":")
						if current[0] == "Authorization" {
							headers[idx] = current[0] + ": *"
						}
					}
					headersToStr := strings.Join(headers, "\r\n")
					l.Error("[Recovery from panic]", zap.Any("panic", err), zap.String("headers", headersToStr), zap.Stack("stack"))
				}
				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
				} else {
					err1 := errors.InternalServer("InternalServer", "the service is abnormal, please wait for a while and try again.")
					c.AbortWithError(int(err1.GetCode()), err1)
				}
			}
		}()
		c.Next()
	}
}
