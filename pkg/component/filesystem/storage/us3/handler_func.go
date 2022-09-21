package us3

import (
	"strings"

	"github.com/cnartlu/area-service/pkg/component/filesystem/storage"
)

//WithStorage 文件存储类型，分别是标准、低频、归档，对应有效值：STANDARD, IA, ARCHIVE
func WithStorage(s string) storage.HandlerFunc {
	return func(c *storage.Context) {
		if s != "" {
			s = strings.ToUpper(strings.TrimSpace(s))
		}
		switch s {
		case "STANDARD", "IA", "ARCHIVE":
		default:
			s = "STANDARD"
		}
		c.Request.Header.Set("X-Ufile-Storage-Class", s)
	}
}
