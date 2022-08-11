package release

import (
	"context"

	"github.com/cnartlu/area-service/internal/component/ent"
)

// const (
// 	defaultOwner = "xiangyuecn"
// 	defaultRepo  = "AreaCity-JsSpider-StatsGov"
// 	regionIDLen  = 10
// )

type Querier interface {
	FindOneWithLastAdded(ctx context.Context) (*ent.AreaRelease, error)
}

func (s *Service) FindOneWithLastRecord(ctx context.Context) (*ent.AreaRelease, error) {
	return s.releaseRepo.FindOneWithLastRecord(ctx)
}

func (s *Service) FindOneWithLastAdded(ctx context.Context) (*ent.AreaRelease, error) {
	return s.LoadLatestRelease(ctx)
	// client := github.NewClient(nil)
	// rep, _, err := client.Repositories.GetLatestRelease(ctx, defaultOwner, defaultRepo)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to getLatestRelease by github client")
	// }
	// areaRelease, err := s.releaseRepo.FindByReleaseID(ctx, uint64(rep.GetID()))
	// if err != nil {
	// 	if !ent.IsNotFound(err) {
	// 		return nil, errors.Wrap(err, "failed to FindByReleaseID by release repository")
	// 	}
	// 	areaRelease, err = s.releaseRepo.Create(ctx, &ent.AreaRelease{
	// 		Owner:              defaultOwner,
	// 		Repository:               defaultRepo,
	// 		ReleaseID:          uint64(rep.GetID()),
	// 		ReleaseName:        rep.GetName(),
	// 		ReleaseNodeID:      rep.GetNodeID(),
	// 		ReleasePublishedAt: uint64(rep.GetPublishedAt().Unix()),
	// 		ReleaseContent:     rep.GetBody(),
	// 		Status:             arearelease.StatusWaitLoaded,
	// 	})
	// 	if err != nil {
	// 		return nil, errors.Wrap(err, "failed to create area.release repository")
	// 	}
	// }
	// // 只是存在ID
	// if areaRelease.Status == arearelease.StatusWaitLoaded {
	// 	for _, asset := range rep.Assets {
	// 		_, err := s.assetRepo.FindOneByAssetID(ctx, uint64(asset.GetID()))
	// 		if err != nil {
	// 			if !ent.IsNotFound(err) {
	// 				return nil, errors.Wrap(err, "failed to FindOneByAssetID by area.release.asset repository")
	// 			}
	// 			downloadUrl := asset.GetBrowserDownloadURL()
	// 			uri, _ := url.Parse(downloadUrl)
	// 			_, err = s.assetRepo.Create(ctx, &ent.AreaReleaseAsset{
	// 				AreaReleaseID: areaRelease.ID,
	// 				AssetID:       uint64(asset.GetID()),
	// 				AssetName:     asset.GetName(),
	// 				AssetLabel:    asset.GetLabel(),
	// 				AssetState:    asset.GetState(),
	// 				DownloadURL:   downloadUrl,
	// 				FilePath:      utils.RelativePath(filepath.Join(utils.RuntimePath(), areaRelease.ReleaseNodeID, filepath.Base(uri.Path))),
	// 				FileSize:      uint(asset.GetSize()),
	// 				Status:        areareleaseasset.StatusWaitLoaded,
	// 			})
	// 			if err != nil {
	// 				return nil, errors.Wrap(err, "failed to create area.release.asset repository")
	// 			}
	// 		}
	// 	}
	// }

	// return areaRelease, nil
}
