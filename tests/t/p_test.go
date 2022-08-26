package t

import (
	"encoding/json"
	"fmt"
	"testing"
)

type a interface {
	Name() string
}

type c struct {
	Content a
}

type jsons struct {
	Name string       `json:"name"`
	S    map[string]c `json:"s,omitempty"`
}

func (jsons) UnmarshalJSON(b []byte) error {
	fmt.Println("我是自定义解析方法", string(b))
	return nil
}

func Test_A(t *testing.T) {
	var z jsons
	json.Unmarshal([]byte(`{
		"name": "name",
		"s": {
			"c": {
				"name": "i"
			}
		}
		}`), &z)
	fmt.Println(z)
}
