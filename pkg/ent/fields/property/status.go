package property

import (
	"database/sql/driver"
	"strconv"
)

// Status 通用的状态码
type Status uint8

const (
	WaitLoaded   Status = 0
	FinishLoaded Status = 1
)

func (p Status) String() string {
	return strconv.Itoa(int(p))
}

// Values provides list valid values for Enum.
func (Status) Values() []string {
	return []string{
		WaitLoaded.String(),
		FinishLoaded.String(),
	}
}

// Value provides the DB a string from int.
func (p Status) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *Status) Scan(val interface{}) error {
	var s string
	switch v := val.(type) {
	case nil:
		return nil
	case string:
		s = v
	case []uint8:
		s = string(v)
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*p = Status(v)
	return nil
}
