package response

type Repository interface{}

type Response struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

func NewSuccessResponse() *Response {
	return &Response{
		Error: 0,
		Msg:   "操作成功",
	}
}

func NewResponse(err int, msg string) Repository {
	return &Response{
		Error: err,
		Msg:   msg,
	}
}

type DataResponse struct {
	*Response
	Data any `json:"data,omitempty"`
}

func NewDataResponse(err int, msg string, data any) Repository {
	return &DataResponse{
		Response: &Response{
			Error: err,
			Msg:   msg,
		},
		Data: data,
	}
}

func NewSuccessDataResponse(data any) Repository {
	return &DataResponse{
		Response: NewSuccessResponse(),
		Data:     data,
	}
}
