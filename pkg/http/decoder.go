package http

import (
	"net/http"
	"unsafe"
)

type Response struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

// DefaultResponseEncoder 请求通过
func DefaultResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	// 通过Request Header的Accept中提取出对应的编码器
	// 如果找不到则忽略报错，并使用默认json编码器
	_ = unsafe.Pointer(&v)
	// is := *(*v)(unsafe.Pointer(&v))
	// if is.pt == 0 && is.pv == 0 {
	// 	//is nil do something
	// }
	// codec, _ := a.CodecForRequest(r, "Accept")
	// data, err := codec.Marshal(v)
	// if err != nil {
	// 	return err
	// }
	// // 在Response Header中写入编码器的scheme
	// w.Header().Set("Content-Type", httputil.ContentType(codec.Name()))
	// w.Write(data)
	return nil
}

// DefaultErrorEncoder 请求错误
func DefaultErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	// 拿到error并转换成kratos Error实体
	// se := errors.FromError(err)
	// // 通过Request Header的Accept中提取出对应的编码器
	// codec, _ := a.CodecForRequest(r, "Accept")
	// body, err := codec.Marshal(se)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", httputil.ContentType(codec.Name()))
	// // 设置HTTP Status Code
	// w.WriteHeader(int(se.Code))
	// w.Write(body)
}
