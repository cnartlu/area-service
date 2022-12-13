package sort

import "strings"

type Sort struct {
	// 字段名
	Field string `json:"field,omitempty"`
	// 倒序
	Desc bool `json:"desc,omitempty"`
}

func (s Sort) String() string {
	if s.Field != "" {
		if s.Desc {
			return "-" + s.Field
		}
		return s.Field
	}
	return ""
}

type Sorts []Sort

func (s Sorts) String() string {
	var str = ""
	for _, ss := range s {
		s1 := ss.String()
		if s1 != "" {
			str = str + "," + s1
		}
	}
	return strings.TrimPrefix(str, ",")
}

func Parse(s string) Sort {
	if s == "" {
		panic("data.sort cannot be empty field")
	}
	var r = Sort{}
	var field = ""
	sLength := len(s)
	if s[0] == '-' {
		if sLength == 1 {
			panic("data.sort cannot be empty field")
		}
		field = strings.TrimSpace(s[1:])
		if field == "" {
			panic("data.sort cannot be empty field")
		}
		r.Desc = true
	} else {
		field = s
	}
	r.Field = field

	return r
}

func ParseArray(s string) Sorts {
	strs := strings.Split(s, ",")
	if len(strs) > 0 {
		sorts := make(Sorts, len(strs))
		for idx, str := range strs {
			str := strings.TrimSpace(str)
			if str == "" {
				continue
			}
			sort := Parse(str)
			sorts[idx] = sort
		}
		return sorts
	}
	return []Sort{}
}
