package validates

import (
	"errors"
	"unicode/utf8"
)

// MinRuneCount 最大字符长度
func MinRuneCount(maxLen int) func(s string) error {
	return func(s string) error {
		if utf8.RuneCountInString(s) < maxLen {
			return errors.New("value is less than minimum length")
		}
		return nil
	}
}

// MaxRuneCount 最大字符长度
func MaxRuneCount(maxLen int) func(s string) error {
	return func(s string) error {
		if utf8.RuneCountInString(s) > maxLen {
			return errors.New("value is more than the max length")
		}
		return nil
	}
}
