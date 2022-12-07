package github

import (
	"context"
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/cnartlu/area-service/api"
	zip7 "github.com/cnartlu/area-service/component/7zip"
	"github.com/cnartlu/area-service/component/app"

	"github.com/cnartlu/area-service/component/filesystem"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/biz/transaction"
	"go.uber.org/zap"
)

type XiangyuecnRepository interface {
	GetLatestRelease(ctx context.Context) (*GithubRepositoryRelease, error)
}

type GithubUsecase struct {
	app         *app.App
	logger      *log.Logger
	filesystem  *filesystem.FileSystem
	xiangyuecn  XiangyuecnRepository
	transaction transaction.Transaction
	areaRepo    area.AreaRepo
	releaseRepo release.ReleaseRepo
	assetRepo   asset.AssetRepo
}

func (g *GithubUsecase) GetLatestRelease(ctx context.Context) (*release.Release, error) {
	githubReponseRelease, err := g.xiangyuecn.GetLatestRelease(ctx)
	if err != nil {
		return nil, err
	}
	rep := githubReponseRelease.RepositoryRelease
	result, err := g.releaseRepo.FindOne(ctx, release.ReleaseIDEQ(uint64(rep.GetID())))
	if err != nil {
		if !api.IsDataNotFound(err) {
			return nil, err
		}
		err = g.transaction.Transaction(ctx, func(ctx context.Context) error {
			if err != nil {
				if !api.IsDataNotFound(err) {
					return err
				}
				result, err = g.releaseRepo.Save(ctx, &release.Release{
					Owner:              githubReponseRelease.Owner,
					Repo:               githubReponseRelease.Repo,
					ReleaseID:          uint64(rep.GetID()),
					ReleaseName:        rep.GetName(),
					ReleaseNodeID:      rep.GetNodeID(),
					ReleasePublishedAt: uint64(rep.GetPublishedAt().Unix()),
					ReleaseContent:     rep.GetBody(),
					Status:             release.StatusWaitSync,
				})
				if err != nil {
					return err
				}
				for _, repAsset := range rep.Assets {
					repAsset := repAsset
					downloadUrl := repAsset.GetBrowserDownloadURL()
					uri, err := url.Parse(downloadUrl)
					if err != nil {
						return err
					}
					var filepath string = filepath.Join("github.com", result.Owner, strconv.FormatUint(result.ReleaseID, 10), filepath.Base(uri.Path))
					_, err = g.assetRepo.Save(ctx, &asset.Asset{
						AreaReleaseID: result.ID,
						AssetName:     repAsset.GetName(),
						AssetID:       uint64(repAsset.GetID()),
						AssetLabel:    repAsset.GetLabel(),
						AssetState:    repAsset.GetState(),
						FileSize:      uint(repAsset.GetSize()),
						FilePath:      filepath,
						DownloadURL:   downloadUrl,
						Status:        asset.StatusWaitSync,
					})
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (g *GithubUsecase) Download(ctx context.Context, data *release.Release) error {
	assets, err := g.assetRepo.FindList(ctx, asset.AreaReleaseIDEQ(data.ID), asset.StatusEQ(asset.StatusWaitSync), asset.Limit(10), asset.Order("-id"))
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	for _, result := range assets {
		result := result
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := g.filesystem.Download(ctx, result.DownloadURL, filepath.Join(g.app.GetRuntimePath(), result.FilePath)); err != nil {
				g.logger.Error("download asset file failed", zap.Error(err), zap.String("filename", result.FilePath), zap.Uint64("areaReleaseAreaID", result.ID))
				return
			}
			result.Status = asset.StatusWaitLoaded
			if _, err1 := g.assetRepo.Save(ctx, result); err1 != nil {
				g.logger.Error("save asset download status failed", zap.Error(err1), zap.Uint64("areaReleaseAreaID", result.ID))
				return
			}
		}()
	}
	wg.Wait()
	return nil
}

func (g *GithubUsecase) LoadReleaseAssets(ctx context.Context, data *release.Release) error {
	assets, err := g.assetRepo.FindList(ctx, asset.StatusEQ(asset.StatusWaitLoaded), asset.Limit(10), asset.Order("-id"))
	if err != nil {
		return err
	}
	for _, result := range assets {
		result := result
		if err := g.LoadFile(ctx, filepath.Join(g.app.GetRuntimePath(), result.FilePath)); err != nil {
			return err
		}
	}
	return nil
}

func (g *GithubUsecase) LoadDir(ctx context.Context, path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		var err error
		if file.IsDir() {
			err = g.LoadDir(ctx, filepath.Join(path, file.Name()))
		} else {
			err = g.LoadFile(ctx, filepath.Join(path, file.Name()))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GithubUsecase) LoadFile(ctx context.Context, filename string) error {
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".7z":
		absPath, _ := filepath.Abs(filename)
		b := md5.Sum([]byte(filename))
		basedir := filepath.Join(filepath.Dir(filename), hex.EncodeToString(b[0:md5.Size-1]))
		if err := zip7.Extract(absPath, basedir); err != nil {
			g.logger.Error("7zip extract failed", zap.Error(err), zap.String("filename", filename))
			return err
		}
		if err := g.LoadDir(ctx, basedir); err != nil {
			return err
		}
	case ".csv":
		f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		r := csv.NewReader(f)
		headers, err := r.Read()
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
		if headers[0][0] == 0xef && headers[0][1] == 0xbb && headers[0][2] == 0xbf {
			headers[0] = headers[0][3:]
		}
		for k, v := range headers {
			v = strings.ToLower(strings.TrimSpace(v))
			headers[k] = v
		}
		headerStr := strings.Join(headers, ",")
		switch headerStr {
		case areaHeaders:
			// 加载区域数据
			if err := g.LoadFileWithAreaData(ctx, filename); err != nil {
				return err
			}
		case areaGeoHeader:
			// 加载数据数据
			if err := g.LoadFileWithAreaGeoData(ctx, filename); err != nil {
				return err
			}
		default:
			g.logger.Warn("unknown csv file format", zap.String("file", filename), zap.String("headers", headerStr))
		}
	default:
		g.logger.Warn("unknown file", zap.String("file", filename))
	}
	return nil
}

// LoadFileWithAreaData 加载区域列表
func (g *GithubUsecase) LoadFileWithAreaData(ctx context.Context, filename string) error {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	idx := 0
	for {
		headers, err := r.Read()
		idx++
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if idx == 1 {
			continue
		}
		id := ToConvertAreaRegionID(headers[0])
		if id == "" {
			return fmt.Errorf("")
		}
		pid := ToConvertAreaRegionID(headers[1])
		if pid == "" {
			return fmt.Errorf("")
		}
		deep, err := strconv.ParseInt(headers[2], 10, 8)
		name := strings.ToLower(headers[3])
		if err := g.transaction.Transaction(ctx, func(ctx context.Context) error {
			areaData, err := g.areaRepo.FindOne(ctx, area.RegionIDEQ(id), area.LevelEQ(int(deep)+1))
			if err != nil {
				if !api.IsDataNotFound(err) {
					return err
				}
				areaData, err = g.areaRepo.Save(ctx, &area.Area{})
				if err != nil {
					return err
				}
			}
			_ = name
			_ = areaData

			return nil
		}); err != nil {
			return err
		}
		// g.logger.Debug("row data", zap.Int("idx", idx), zap.Strings("rows", headers))
	}
	return nil
}

// LoadFileWithAreaGeoData 加载区域范围坐标
func (g *GithubUsecase) LoadFileWithAreaGeoData(ctx context.Context, filename string) error {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	idx := 0
	for {
		headers, err := r.Read()
		idx++
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if idx == 1 {
			continue
		}
		_ = headers
		// g.logger.Debug("row data", zap.Int("idx", idx), zap.Strings("rows", headers))
	}
	return nil
}

func NewGithubUsecase(
	app *app.App,
	logger *log.Logger,
	transaction transaction.Transaction,
	filesystem *filesystem.FileSystem,
	xiangyuecn XiangyuecnRepository,
	areaRepo area.AreaRepo,
	releaseRepo release.ReleaseRepo,
	assetRepo asset.AssetRepo,
) *GithubUsecase {
	return &GithubUsecase{
		app:        app,
		logger:     logger,
		filesystem: filesystem,
		xiangyuecn: xiangyuecn,

		transaction: transaction,
		releaseRepo: releaseRepo,
		assetRepo:   assetRepo,
		areaRepo:    areaRepo,
	}
}
