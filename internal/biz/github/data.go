package github

import (
	"strconv"
	"strings"

	"github.com/google/go-github/v45/github"
)

var (
	areaHeaders   = strings.Join([]string{"id", "pid", "deep", "name", "pinyin_prefix", "pinyin", "ext_id", "ext_name"}, ",")
	areaGeoHeader = strings.Join([]string{"id", "pid", "deep", "name", "ext_path", "geo", "polygon"}, ",")
)

type GithubRepositoryRelease struct {
	Owner             string
	Repo              string
	RepositoryRelease *github.RepositoryRelease
}

// ToConvertAreaRegionID 将字符串转为区域地址ID
func ToConvertAreaRegionID(s string) string {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return ""
	}
	idStr := strconv.FormatUint(id, 10)
	if len(idStr) >= 10 {
		idStr = idStr[0:10]
	} else {
		idStr = idStr + strings.Repeat("0", 10-len(idStr))
	}
	return idStr
}
