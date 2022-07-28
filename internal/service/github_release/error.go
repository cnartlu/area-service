package github_release

type e string

func (s e) Error() string {
	return string(s)
}

const (
	DataHeaderError e = "the data is table header"
	DataEmptyError  e = "the data is empty"
)

// IsDataHeaderError 是否为表格表头
func IsDataHeaderError(err error) bool {
	return err == DataHeaderError
}

// IsDataEmptyError 数据是否为空值
func IsDataEmptyError(err error) bool {
	return err == DataEmptyError
}
