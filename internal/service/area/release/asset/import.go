package asset

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/areareleaseasset"
	"github.com/cnartlu/area-service/pkg/utils"
	"github.com/pkg/errors"
)

type Importer interface {
	Import(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset) error
}

func (s *Service) Import(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset) error {
	if areaReleaseAsset.Status == areareleaseasset.StatusWaitLoaded {
		filePath := filepath.FromSlash(filepath.Join(utils.RootPath(), areaReleaseAsset.FilePath))
		if !utils.IsFile(filePath) {
			// 代理下载
			var p *http.Client
			if s.application != nil && s.application.Proxy != "" {
				p = s.application.ProxyClient()
			}
			httpClient := utils.NewHttpClient(p)
			err := httpClient.Download(ctx, areaReleaseAsset.DownloadURL, filePath)
			if err != nil {
				return errors.Wrap(err, "download file failed")
			}
		}
		// 下载完成后执行导入
		switch filepath.Base(filePath) {
		case "ok_data_level3.csv":
		case "ok_data_level4.csv":
		case "ok_geo.csv.7z":
		}
		areaReleaseAsset.Status = areareleaseasset.StatusFinishedLoaded
		_, err := s.repo.Update(ctx, areaReleaseAsset)
		if err != nil {
			return err
		}
	}
	return nil
}
