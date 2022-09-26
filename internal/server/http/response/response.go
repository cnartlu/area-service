package response

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
)

func NewSuccessResponse() *errors.Error {
	return errors.New(http.StatusOK, "SUCCESS", "操作成功")
}

func NewResponse(err int, msg string) *errors.Error {
	return errors.New(err, "UNKNOWN", msg)
}

type DataResponse struct {
	*errors.Error
	Data any `json:"data,omitempty"`
}

func NewDataResponse(err int, msg string, data any) *DataResponse {
	return &DataResponse{
		Error: errors.New(err, "UNKNOWN", msg),
		Data:  data,
	}
}

func NewSuccessDataResponse(data any) *DataResponse {
	return &DataResponse{
		Error: NewSuccessResponse(),
		Data:  data,
	}
}
