package github_release

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/go-github/v45/github"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	DefaultOwner = "xiangyuecn"
	DefaultRepo  = "AreaCity-JsSpider-StatsGov"
	regionIDLen  = 10
)

// GetLatestRelease 获取最后一条数据
func GetLatestRelease(client *github.Client, ctx context.Context, owner, repo string) (*github.RepositoryRelease, error) {
	if owner == "" {
		owner = DefaultOwner
		repo = DefaultRepo
	}
	rep, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	return rep, nil
}

// DownloadAsset 下载资源文件
func DownloadAsset(path string, asset *github.ReleaseAsset, httpClient *http.Client) (string, error) {
	if asset.BrowserDownloadURL == nil {
		return "", errors.New("")
	}
	browserDownloadURL := *asset.BrowserDownloadURL
	uri, err := url.Parse(browserDownloadURL)
	if err != nil {
		return "", err
	}
	// 文件保存的完整地址
	filename := filepath.Join(path, filepath.Base(uri.Path))
	// 请求客户端
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	resp, err := httpClient.Get(browserDownloadURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// 打开文件
	out, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		defer os.Remove(filename)
		return "", err
	}
	return filename, nil
}

// FormatToLevelRow 格式化导入的等级区域地址数据行
func FormatToLevelRow(row []string, regionIDLen int) (*LevelRow, error) {
	rowLen := len(row)
	if rowLen == 0 {
		return nil, errors.New("empty row")
	}
	var data = []string{}
	for _, value := range row {
		value = strings.TrimSpace(value)
		data = append(data, value)
	}
	id := data[0]
	if id == "" {
		return nil, errors.New("empty row")
	}
	id = strings.ToLower(id)
	if id == "id" {
		return nil, errors.New("empty row")
	}
	levelIdLen := regionIDLen
	// ID长度补全
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	id = strconv.FormatUint(idUint, 10)
	idLen := len(id)
	if idLen > levelIdLen {
		id = id[0:levelIdLen]
	} else if idLen < levelIdLen {
		id = id + strings.Repeat("0", levelIdLen-idLen)
	}
	// PID长度补全
	pidUint, err := strconv.ParseUint(data[1], 10, 64)
	if err != nil {
		return nil, err
	}
	pid := strconv.FormatUint(pidUint, 10)
	pidLen := len(pid)
	if pidLen > levelIdLen {
		pid = pid[0:levelIdLen]
	} else if pidLen < levelIdLen {
		pid = pid + strings.Repeat("0", levelIdLen-pidLen)
	}
	// 深度和级别
	level, err := strconv.ParseInt(data[2], 10, 64)
	if err != nil {
		return nil, err
	}
	// 名称
	name := data[3]
	// 拼音
	pinyin := data[5]
	if pinyin == "" || pinyin[0:1] == "-" {
		pinyin = ""
	} else {
		pinyin = cases.Title(language.Make(pinyin)).String(pinyin)
	}
	// 原始ID
	extId := data[6]
	// 原始名称
	extName := data[7]

	return &LevelRow{
		ID:      id,
		Pid:     pid,
		Level:   int(level + 1),
		Pinyin:  pinyin,
		Name:    name,
		ExtId:   extId,
		ExtName: extName,
	}, nil
}

// FormatToGeoRow 格式化导入的等级区域地址数据行
func FormatToGeoRow(row []string) (*GeoRow, error) {
	rowLen := len(row)
	if rowLen == 0 {
		return nil, DataEmptyError
	}
	id := strings.TrimSpace(row[0])
	if id == "" {
		return nil, DataEmptyError
	}
	id = strings.ToLower(id)
	if id == "id" {
		return nil, DataHeaderError
	}
	levelIdLen := regionIDLen

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	id = strconv.FormatUint(idUint, 10)
	idLen := len(id)
	if idLen > levelIdLen {
		id = id[0:levelIdLen]
	} else if idLen < levelIdLen {
		id = id + strings.Repeat("0", levelIdLen-idLen)
	}

	pidUint, err := strconv.ParseUint(row[1], 10, 64)
	if err != nil {
		return nil, err
	}
	pid := strconv.FormatUint(pidUint, 10)
	pidLen := len(pid)
	if pidLen > levelIdLen {
		pid = pid[0:levelIdLen]
	} else if pidLen < levelIdLen {
		pid = pid + strings.Repeat("0", levelIdLen-pidLen)
	}
	// 深度和级别
	level, err := strconv.ParseInt(row[2], 10, 64)
	if err != nil {
		return nil, err
	}
	// 名称
	name := row[3]
	// Geo
	geoData := toGeo(row[5])
	// Polygon
	polygons := toPolygons(row[6])
	return &GeoRow{
		ID:       id,
		Pid:      pid,
		Level:    int(level + 1),
		Name:     name,
		Geo:      geoData,
		Polygons: polygons,
	}, nil
}
