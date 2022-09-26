package strings

import (
	"encoding/json"
)

type OrderField struct {
	// field 字段名称
	Field string
	// IsASC 是否正序排序
	IsASC bool
}

type Order struct {
	str string
}

func (o *Order) String() string {
	return o.str
}

func (o *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}

func (o *Order) UnmarshalJSON(data []byte) error {
	return nil
}
