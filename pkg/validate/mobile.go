package validate

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

// IsMobilePhone 校验是否为手机号
// 基于 google 的 libphonenumber 库的 go 版本
func IsMobilePhone(phone string) error {
	phoneNumber, err := phonenumbers.Parse(phone, "CN")
	if err != nil {
		return err
	}

	if !phonenumbers.IsValidNumber(phoneNumber) {
		return fmt.Errorf("手机号码格式无效")
	}

	return nil
}
