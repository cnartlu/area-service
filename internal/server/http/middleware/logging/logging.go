package logging

import (
	"time"

	"github.com/cnartlu/area-service/component/log"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	"go.uber.org/zap"
)

func ServerByGin(logger *log.Logger) gin.HandlerFunc {
	l := logger.AddCallerSkip(1)
	return func(c *gin.Context) {
		var (
			code      int32
			reason    string
			kind      string
			operation string
		)
		startTime := time.Now()
		if info, ok := transport.FromServerContext(c.Request.Context()); ok {
			kind = info.Kind().String()
			operation = info.Operation()
		}

		c.Next()

		if se := errors.FromError(c.Err()); se != nil {
			code = se.Code
			reason = se.Reason
		}

		l.Info("http request",
			zap.String("component", kind),
			zap.String("operation", operation),
			zap.Int32("code", code),
			zap.String("reason", reason),
			zap.Float64("latency", time.Since(startTime).Seconds()),
		)
	}
}
