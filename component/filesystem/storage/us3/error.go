package us3

import (
	"encoding/json"
	"fmt"
	"io"
)

type Error struct {
	// 返回状态码，为 0 则为成功返回，非 0 为失败
	RetCode int `json:"RetCode"`
	// 返回错误消息，当 RetCode 非 0 时提供详细的描述信息
	ErrMsg string `json:"ErrMsg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("err code %d, message: %s", e.RetCode, e.ErrMsg)
}

func newResponseError(data io.Reader) error {
	bs, err := io.ReadAll(data)
	if err != nil {
		if io.EOF == err {
			return nil
		}
		return err
	}
	e := Error{}
	err = json.Unmarshal(bs, &data)
	if err != nil {
		return err
	}
	return e
}
