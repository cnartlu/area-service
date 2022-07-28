package utils

// InArray 判断数据是否在数组内
func InArray[T interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | string
}](s T, ss []T) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}
