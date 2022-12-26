package middleware

import (
	"net/http"

	stderrors "github.com/cnartlu/area-service/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
)

const ResponseDataKey = "_gin-gonic/gin/responseDataKey"

var SuccessResponse = stderrors.ErrorSuccess(stderrors.Error_SUCCESS.String())

type response struct {
	errors.Error
	Data any `json:"data,omitempty"`
}

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			ginErr := c.Errors.Last()
			var err *errors.Error
			switch ginErr.Type {
			case gin.ErrorTypeBind:
				err = stderrors.ErrorParamFormat(ginErr.Error())
			case gin.ErrorTypeRender:
				err = stderrors.ErrorServerError(ginErr.Error())
			case gin.ErrorTypePrivate:
				err = errors.FromError(ginErr)
			case gin.ErrorTypePublic:
				fallthrough
			case gin.ErrorTypeAny:
				fallthrough
			default:
				err = errors.FromError(ginErr)
			}
			c.AbortWithStatusJSON(int(err.Code), err)
			return
		}
		if res, ok := c.Get(ResponseDataKey); ok {
			c.JSON(http.StatusOK, response{
				Error: *SuccessResponse,
				Data:  res,
			})
			return
		}
		c.JSON(http.StatusOK, SuccessResponse)
	}
}
