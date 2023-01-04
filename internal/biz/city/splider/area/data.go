package area

import (
	"strconv"
	"strings"
)

const (
	RegionIDLength = 10
)

type Area struct {
	ID             int
	ParentID       int
	RegionID       string
	ParentRegionID string
	Title          string
	Level          int
	Lng            float64
	Lat            float64
}

// 将字符串转为区域地址ID
func ToConvertRegionID(s string) string {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return ""
	}
	idStr := strconv.FormatUint(id, 10)
	if len(idStr) >= RegionIDLength {
		idStr = idStr[0:RegionIDLength]
	} else {
		idStr = idStr + strings.Repeat("0", RegionIDLength-len(idStr))
	}
	return idStr
}
