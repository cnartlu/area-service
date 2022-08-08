package asset

import (
	"context"
	"path/filepath"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/areareleaseasset"
	"github.com/cnartlu/area-service/pkg/utils"
)

type Importer interface {
	Import(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset) error
}

func (s *Service) Import(ctx context.Context, areaReleaseAsset *ent.AreaReleaseAsset) error {
	if areaReleaseAsset.Status == areareleaseasset.StatusWaitLoaded {
		filePath := filepath.Join(utils.RuntimePath(), areaReleaseAsset.FilePath)
		if !utils.IsFile(filePath) {
			err := utils.Download(ctx, areaReleaseAsset.DownloadURL, filePath)
			if err != nil {
				return err
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
