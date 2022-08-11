package importer

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/component/log"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/areareleaseasset"
	"github.com/cnartlu/area-service/pkg/utils"
	"github.com/pkg/errors"
)

type Servicer interface {
	Import(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset, httpClient *http.Client) error
}

type Service struct {
	logger      *log.Logger
	application *config.Application
}

func (s *Service) Import(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset, httpClient *http.Client) error {
	if areaReleaseAsset.Status == areareleaseasset.StatusFinishedLoaded {
		return nil
	}
	filePath := filepath.FromSlash(filepath.Join(utils.RootPath(), areaReleaseAsset.FilePath))
	if !utils.IsFile(filePath) {
		if httpClient == nil && s.application != nil {
			httpClient = s.application.ProxyClient()
		}
		d := utils.NewHttpClient(httpClient)
		err := d.Download(ctx, areaReleaseAsset.DownloadURL, filePath)
		if err != nil {
			return errors.Wrap(err, "utils.HttpClient Download")
		}
	}
	// 下载完成后执行导入
	switch filepath.Base(filePath) {
	case "ok_data_level3.csv":
	case "ok_data_level4.csv":
	case "ok_geo.csv.7z":
	}
	return nil
}

func NewService(
	logger *log.Logger,
	application *config.Application,
) *Service {
	return &Service{
		logger:      logger,
		application: application,
	}
}
