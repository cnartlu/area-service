package job

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/cnartlu/area-service/pkg/path"
	"github.com/google/go-github/v45/github"
	"go.uber.org/zap"
)

var (
	github_org     = "xiangyuecn"
	github_package = "AreaCity-JsSpider-StatsGov"
)

// StatsGov 省市区数据采集并标注拼音、坐标和边界范围
type Github struct {
	logger    *log.Logger
	db        *db.DB
	httpProxy *http.Client
	client    *github.Client
	ctx       context.Context
}

// NewExample 构造函数
func NewGithub(
	logger *log.Logger,
	db *db.DB,
	httpProxy *http.Client,
) *Github {
	var (
		ctx    = context.Background()
		client = github.NewClient(httpProxy)
		g      = Github{
			logger:    logger,
			httpProxy: httpProxy,
			client:    client,
			db:        db,
			ctx:       ctx,
		}
	)

	if httpProxy == nil {
		uri, _ := url.Parse("http://127.0.0.1:7890")
		g.httpProxy = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(uri),
			},
		}
	}
	return &g
}

// Run 任务执行方法
func (s Github) Run() {
	s.logger.Info("start 开始执行计划任务")
	// 第一步 拉取最新的数据资料
	s.logger.Debug("获取区域地址资料")
	relase, err := s.GetRepository()
	if err != nil {
		return
	}
	pathdir := filepath.Join(path.RootPath(), "/assets", github_org)
	if err := s.DownloadRelease(pathdir, relase); err != nil {
		return
	}
	// 开启

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
	dirpath := filepath.Join(path, strconv.Itoa(int(*release.ID)))
	if _, err := os.Stat(dirpath); err != nil {
		if os.IsExist(err) {
			s.logger.Error("", zap.Error(err))
			return err
		}
		err := os.MkdirAll(dirpath, 0644)
		if err != nil {
			s.logger.Error("创建目录失败", zap.Error(err))
			return err
		}
	}
	for _, asset := range release.Assets {
		browserDownloadURL := asset.BrowserDownloadURL
		uri, _ := url.Parse(*browserDownloadURL)
		filename := filepath.Join(dirpath, filepath.Base(uri.Path))
		// 检查文件是否已经下载成功
		fstat, err := os.Stat(filename)
		if err != nil && os.IsExist(err) {
			s.logger.Error("stat file exist error", zap.String("filename", filename), zap.Error(err))
			continue
		}
		if fstat != nil {
			if asset.GetSize() > 0 && int(fstat.Size()) < asset.GetSize() {
				// 已经存在部分下载的数据
				if err := os.Remove(filename); err != nil {
					s.logger.Error("delete os file error", zap.String("filename", filename), zap.Error(err))
					continue
				}
			}
		}
		out, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			s.logger.Error("open file error", zap.String("filename", filename), zap.String("url", *browserDownloadURL), zap.Error(err))
			continue
		}
		defer out.Close()
		resp, err := s.httpProxy.Get(*browserDownloadURL)
		if err != nil {
			s.logger.Error("request file url error", zap.String("url", *browserDownloadURL), zap.Error(err))
			continue
		}
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			s.logger.Error("file copy failed", zap.String("filename", filename), zap.Error(err))
			continue
		}
	}
	return nil
}
