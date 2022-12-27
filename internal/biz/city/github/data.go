package github

import (
	"strings"

	"github.com/cnartlu/area-service/errors"
	"github.com/cnartlu/area-service/internal/biz/city/splider"
	"github.com/cnartlu/area-service/internal/biz/city/splider/asset"
)

const Source = splider.SourceGithub

type Account struct {
	User string
	Repo string
}

type Github struct {
	*splider.Splider
	Assets []*asset.Asset
}

// ReaderFileType 读取文件类型
type ReaderFileType int

const (
	ReaderFileTypeUnknow ReaderFileType = iota
	ReaderFileTypeArea
	ReaderFileTypeGeo
)

var (
	areaHeaders          = []string{"id", "pid", "deep", "name", "pinyin_prefix", "pinyin", "ext_id", "ext_name"}
	geoHeaders           = []string{"id", "pid", "deep", "name", "ext_path", "geo", "polygon"}
	areaHeaderStr string = strings.Join(areaHeaders, ",")
	geoHeaderStr  string = strings.Join(geoHeaders, ",")
)

var (
	Err7zipExtractError     = errors.ErrorServerError("7zip extract error")
	ErrUnsupportedFileType  = errors.ErrorServerError("unsupported file type")
	ErrUnsupportedSheetFile = errors.ErrorServerError("unsupported sheet file")
)
