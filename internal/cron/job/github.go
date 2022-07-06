package job

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/cases"

	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/cnartlu/area-service/internal/component/ent/area"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/cnartlu/area-service/pkg/path"
	"github.com/gen2brain/go-unarr"
	"github.com/google/go-github/v45/github"
	"go.uber.org/zap"
)

var (
	github_org     = "xiangyuecn"
	github_package = "AreaCity-JsSpider-StatsGov"
)

type areaLevel struct {
	Id      string
	Pid     string
	Level   int
	Name    string
	Pinyin  string
	ExtId   string
	ExtName string
}

// StatsGov 省市区数据采集并标注拼音、坐标和边界范围
type Github struct {
	logger     *log.Logger
	db         *db.DB
	httpClient *http.Client
	client     *github.Client
	ctx        context.Context
}

// NewExample 构造函数
func NewGithub(
	logger *log.Logger,
	db *db.DB,
	httpClient *http.Client,
) *Github {
	var (
		ctx    = context.Background()
		client = github.NewClient(httpClient)
		g      = Github{
			logger:     logger,
			httpClient: httpClient,
			client:     client,
			db:         db,
			ctx:        ctx,
		}
	)
	if httpClient == nil {
		g.httpClient = &http.Client{}
	}
	return &g
}

// Run 任务执行方法
func (s Github) Run() {
	s.logger.Info("start 开始执行计划任务")
	// 第一步 拉取最新的数据资料
	s.logger.Debug("获取区域地址资料")
	release, err := s.GetRepository()
	if err != nil {
		return
	}
	pathdir := filepath.Join(path.RootPath(), "/assets", github_org)
	dirpath := filepath.Join(pathdir, strconv.Itoa(int(*release.ID)))
	if err := s.DownloadRelease(dirpath, release); err != nil {
		return
	}
	// 开启
	filenames, _ := filepath.Glob(dirpath + "/ok_*")
	for _, filename := range filenames {
		s.logger.Debug("Open file", zap.String("filename", filename))
		file := filepath.Base(filename)
		switch file {
		case "ok_data_level3.csv":
			if err := s.ImportLevel3(filename); err != nil {
				s.logger.Error("import file error", zap.String("filename", filename), zap.Error(err))
			}
		case "ok_data_level4.csv":
			if err := s.ImportLevel4(filename); err != nil {
				s.logger.Error("import file error", zap.String("filename", filename), zap.Error(err))
			}
		case "ok_geo.csv.7z":
			s.logger.Debug("unzip file", zap.String("file", file))
			archive, err := unarr.NewArchive(filename)
			if err != nil {
				s.logger.Error("archive error", zap.String("filename", filename), zap.Error(err))
				continue
			}
			defer archive.Close()
			_, err = archive.Extract(dirpath)
			if err != nil {
				s.logger.Error("archive file extract error", zap.String("filename", filename), zap.Error(err))
				continue
			}
			if err := s.ImportGeo(filename); err != nil {
				s.logger.Error("import file error", zap.String("filename", filename), zap.Error(err))
			}
		}
	}

	s.logger.Info("end 任务执行成功")
}

func (s Github) GetRepository() (*github.RepositoryRelease, error) {
	rep, _, err := s.client.Repositories.GetLatestRelease(s.ctx, github_org, github_package)
	if err != nil {
		s.logger.Info("get github package failed", zap.Error(err))
		return nil, err
	}
	return rep, nil
}

func (s *Github) DownloadRelease(path string, release *github.RepositoryRelease) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			s.logger.Error("", zap.Error(err))
			return err
		}
		err := os.MkdirAll(path, 0644)
		if err != nil {
			s.logger.Error("创建目录失败", zap.Error(err))
			return err
		}
	}
	for _, asset := range release.Assets {
		browserDownloadURL := asset.BrowserDownloadURL
		uri, _ := url.Parse(*browserDownloadURL)
		filename := filepath.Join(path, filepath.Base(uri.Path))
		// 检查文件是否已经下载成功
		fstat, err := os.Stat(filename)
		if err != nil && os.IsExist(err) {
			s.logger.Error("stat file exist error", zap.String("filename", filename), zap.Error(err))
			continue
		}
		if fstat != nil {
			assetSize := asset.GetSize()
			filesize := int(fstat.Size())
			if assetSize <= filesize {
				continue
			}
			// 已经存在部分下载的数据
			if err := os.Remove(filename); err != nil {
				s.logger.Error("delete os file error", zap.String("filename", filename), zap.Error(err))
				continue
			}
		}
		out, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			s.logger.Error("open file error", zap.String("filename", filename), zap.String("url", *browserDownloadURL), zap.Error(err))
			continue
		}
		defer out.Close()
		httpClient := s.httpClient
		if httpClient == nil {
			httpClient = &http.Client{}
		}
		s.logger.Debug("download file", zap.String("filename", filename), zap.String("url", *browserDownloadURL))
		resp, err := httpClient.Get(*browserDownloadURL)
		if err != nil {
			defer func() {
				out.Close()
				os.Remove(filename)
			}()
			s.logger.Error("request file url error", zap.String("url", *browserDownloadURL), zap.Error(err))
			continue
		}
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			defer func() {
				out.Close()
				os.Remove(filename)
			}()
			s.logger.Error("file copy failed", zap.String("filename", filename), zap.Error(err))
			continue
		}
	}
	return nil
}

func (s *Github) ImportLevel3(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	rowIndex := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		s.logger.Debug("level3 line", zap.String("filename", filename), zap.Strings("line", line))
		data, err := s.formatImportLevelRow(line)
		if err != nil {
			s.logger.Error("level3 row error", zap.String("filename", filename), zap.Int("row", rowIndex), zap.Strings("row", line))
			continue
		}
		s.logger.Info("format row data", zap.Any("data", data))
	}
	return nil
}

func (s *Github) ImportLevel4(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	rowIndex := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		data, err := s.formatImportLevelRow(line)
		if err != nil {
			s.logger.Error("level3 row error", zap.String("filename", filename), zap.Int("row", rowIndex), zap.Strings("row", line))
			continue
		}
		// 插入更新记录
		s.db.DB().Area.Query().Where(area.RegionIDEQ(data.Id)).Exist(s.ctx)
	}
	return nil
}

func (s *Github) ImportGeo(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		s.logger.Debug("geo line", zap.String("filename", filename), zap.Strings("line", line))

	}

	return nil
}

func (s *Github) formatImportLevelRow(row []string) (*areaLevel, error) {
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
	levelIdLen := 12
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
		pid = id[0:levelIdLen]
	} else if pidLen < levelIdLen {
		pid = id + strings.Repeat("0", levelIdLen-pidLen)
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
		pinyin = cases.Title(pinyin)
	}
	// 原始ID
	extId := data[6]
	// 原始名称
	extName := data[7]

	return &areaLevel{
		Id:      id,
		Pid:     pid,
		Level:   int(level + 1),
		Pinyin:  pinyin,
		Name:    name,
		ExtId:   extId,
		ExtName: extName,
	}, nil
}
