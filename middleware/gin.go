package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware wraps the request context with a timeout
func LoggerMiddleware(logger *zap.Logger, notlogged ...string) func(c *gin.Context) {
	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()

		{
			ctxerr := c.Request.Context().Err()
			if ctxerr != nil {
				c.Error(ctxerr)
			}
		}

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			endtime := time.Now()
			if raw != "" {
				path = path + "?" + raw
			}
			logger.Info(
				"request info",
				zap.String("requestID", c.GetString("requestID")),
				zap.Duration("latency", endtime.Sub(start)),
				zap.String("clientIP", c.ClientIP()),
				zap.String("method", c.Request.Method),
				zap.Int("status_code", c.Writer.Status()),
				zap.Int("body_size", c.Writer.Size()),
				zap.String("path", path),
				zap.String("msg", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			)
		}
	}
}

// TimeoutMiddleware wraps the request context with a timeout
func TimeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	// 控制请求
	return func(c *gin.Context) {
		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {
				c.Abort()
			}
			cancel()
		}()

		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
