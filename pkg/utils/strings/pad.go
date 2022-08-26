package strings

import (
	"strconv"
	"strings"
)

const (
	STR_PAD_RIGHT = 0b01
	STR_PAD_LEFT  = 0b10
	STR_PAD_BOTH  = 0b11
)

// Pad 使用另一个字符串填充字符串为指定长度
// s 是输入的值
// 如果 padLength 的值是负数，小于或者等于输入字符串的长度，不会发生任何填充，并会返回 s 。
func Pad[T interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | string
}](s T, padLength int, padStr string, padType int) string {
	var (
		str string
	)
	switch ss := (interface{})(s).(type) {
	case string:
		str = ss
	case int:
		str = strconv.FormatInt(int64(ss), 10)
	case int8:
		str = strconv.FormatInt(int64(ss), 10)
	case int16:
		str = strconv.FormatInt(int64(ss), 10)
	case int32:
		str = strconv.FormatInt(int64(ss), 10)
	case int64:
		str = strconv.FormatInt(ss, 10)
	case uint:
		str = strconv.FormatUint(uint64(ss), 10)
	case uint8:
		str = strconv.FormatUint(uint64(ss), 10)
	case uint16:
		str = strconv.FormatUint(uint64(ss), 10)
	case uint32:
		str = strconv.FormatUint(uint64(ss), 10)
	case uint64:
		str = strconv.FormatUint(ss, 10)
	}
	var (
		strLen   = len(str)
		forCount = padLength - strLen
	)
	if forCount > 0 {
		var (
			b  = strings.Builder{}
			lb = strings.Builder{}
			rb = strings.Builder{}
		)
		switch padType {
		case STR_PAD_LEFT:
			lb.WriteString(strings.Repeat(padStr, forCount))
		case STR_PAD_RIGHT:
			rb.WriteString(strings.Repeat(padStr, forCount))
		case STR_PAD_BOTH:
			for i := 0; i < (forCount); i++ {
				if i%2 == 0 {
					rb.WriteString(padStr)
				} else {
					lb.WriteString(padStr)
				}
			}
		}
		b.WriteString(lb.String())
		b.WriteString(str)
		b.WriteString(rb.String())
		str = b.String()
	}
	return str
}
