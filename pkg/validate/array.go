package validate

// t 通用的基础类型
type t interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | string | float32 | float64
}

// InArray 判断数据是否在数组内
func InArray[T t](s T, ss []T) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}

// ArraySearch 查找数据所在位置
// 返回-1表示未找到
func ArraySearch[T t](s T, ss []T) int {
	for k, v := range ss {
		if s == v {
			return k
		}
	}
	return -1
}
