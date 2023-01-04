package github

import (
	"context"
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	zip7 "github.com/cnartlu/area-service/component/7zip"
	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/component/filesystem"
	"github.com/cnartlu/area-service/errors"
	"github.com/cnartlu/area-service/internal/biz/city/splider"
	"github.com/cnartlu/area-service/internal/biz/city/splider/area"
	"github.com/cnartlu/area-service/internal/biz/city/splider/area/polygon"
	"github.com/cnartlu/area-service/internal/biz/city/splider/asset"
	"github.com/cnartlu/area-service/internal/biz/transaction"
	"github.com/google/go-github/v45/github"
	"golang.org/x/sync/errgroup"
)

type GithubRepo interface {
	GetLatestRelease(ctx context.Context, owner string, repo string) (*github.RepositoryRelease, error)
}

type GithubUsecase struct {
	account            Account
	repo               GithubRepo
	app                *app.App
	filesystem         *filesystem.FileSystem
	transaction        transaction.Transaction
	spliderUsecase     *splider.SpliderUsecase
	assetUsecase       *asset.AssetUsecase
	areaUsecase        *area.AreaUsecase
	areaPolygonUsecase *polygon.PolygonUsecase
}

func (g *GithubUsecase) LatestRelease(ctx context.Context) (*Github, error) {
	var result Github
	release, err := g.repo.GetLatestRelease(ctx, g.account.User, g.account.Repo)
	if err != nil {
		return nil, err
	}
	spliderResponse, err := g.spliderUsecase.FindOneWithInstance(ctx,
		splider.SourceIDEQ(uint64(release.GetID())),
		splider.SourceEQ(Source),
		splider.OwnerEQ(g.account.User),
		splider.RepoEQ(g.account.Repo),
		splider.Cache(-1),
	)
	if err != nil {
		if !errors.IsDataNotFound(err) {
			return nil, err
		}
		spliderResponse, err = g.spliderUsecase.Save(ctx, &splider.Splider{
			Source:      Source,
			Owner:       g.account.User,
			Repo:        g.account.Repo,
			SourceID:    uint64(release.GetID()),
			Title:       release.GetName(),
			Draft:       release.GetDraft(),
			PreRelease:  release.GetPrerelease(),
			PublishedAt: release.GetPublishedAt().Time,
			Status:      splider.StatusWaitSync,
			CreatedAt:   release.GetCreatedAt().Time,
		})
		if err != nil {
			return nil, err
		}
	}
	result.Splider = spliderResponse
	if len(release.Assets) > 0 {
		var assets = make([]*asset.Asset, len(release.Assets))
		for k, data := range release.Assets {
			data := data
			assetData, err := g.assetUsecase.FindOneWithInstance(ctx, asset.CitySpliderIDEQ(spliderResponse.ID), asset.SourceIDEQ(uint64(data.GetID())))
			if err != nil {
				if !errors.IsDataNotFound(err) {
					return nil, err
				}
				assetData = &asset.Asset{
					CitySpliderID: spliderResponse.ID,
					SourceID:      uint64(data.GetID()),
					FileTitle:     data.GetName(),
					FilePath:      data.GetBrowserDownloadURL(),
					FileSize:      uint(data.GetSize()),
					Status:        asset.StatusWaitSync,
				}
				assetData, err = g.assetUsecase.Create(ctx, assetData)
				if err != nil {
					return nil, err
				}
			}
			assets[k] = assetData
		}
		result.Assets = assets
	}
	return &result, nil
}

func (g *GithubUsecase) LoadLatestRelease(ctx context.Context) error {
	result, err := g.spliderUsecase.FindOneWithInstance(ctx, splider.SourceEQ(splider.SourceGithub), splider.Order("-id"))
	if err != nil {
		return err
	}
	assets, err := g.assetUsecase.FindList(ctx, asset.FindListParams{
		CitySpliderID: result.ID,
	})
	if err != nil {
		return err
	}
	eg, fctx := errgroup.WithContext(ctx)
	for _, data := range assets {
		data := data
		eg.Go(func() error {
			return g.LoadReleaseAsset(fctx, data, false)
		})
	}
	return eg.Wait()
}

func (g *GithubUsecase) LoadReleaseAsset(ctx context.Context, data *asset.Asset, force bool) error {
	if force {
		data.Status = asset.StatusWaitSync
	}
	switch data.Status {
	case asset.StatusFinished:
	case asset.StatusWaitSync:
		var err error
		err = g.DownloadAsset(ctx, data)
		if err != nil {
			return err
		}
		data.Status = asset.StatusWaitLoaded
		data, err = g.assetUsecase.Update(ctx, data)
		if err != nil {
			return err
		}
		fallthrough
	case asset.StatusWaitLoaded:
		var err error
		var filename string = g.GetFilename(ctx, data, true)
		if err := g.WriterFile(ctx, filename); err != nil {
			return err
		}
		data.Status = asset.StatusFinished
		data, err = g.assetUsecase.Update(ctx, data)
		if err != nil {
			return err
		}
		_ = data
	default:
	}
	return nil
}

func (g *GithubUsecase) GetFilename(ctx context.Context, data *asset.Asset, shortPath bool) string {
	var filename = filepath.Join("city", "splider", "asset", strconv.FormatUint(uint64(data.ID), 10), strconv.FormatUint(data.SourceID, 10)+filepath.Ext(data.FileTitle))
	var runtimePath = g.app.GetRuntimePath()
	var fullFilename = filepath.Join(runtimePath, filename)
	if shortPath {
		var rootPath = g.app.GetRootPath()
		if strings.HasPrefix(fullFilename, rootPath) {
			return fullFilename[len(rootPath)+1:]
		}
	}
	return fullFilename
}

func (g *GithubUsecase) DownloadAsset(ctx context.Context, data *asset.Asset) error {
	var fullFilename = g.GetFilename(ctx, data, false)
	if f, err := os.Stat(fullFilename); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := g.filesystem.Download(ctx, data.FilePath, fullFilename); err != nil {
			return err
		}
	} else if f.IsDir() {
		return os.ErrInvalid
	}
	return nil
}

func (g *GithubUsecase) WriterFile(ctx context.Context, filename string) error {
	{
		fstat, err := os.Stat(filename)
		if err != nil {
			return err
		}
		if fstat.IsDir() {
			files, err := os.ReadDir(filename)
			if err != nil {
				return err
			}
			for _, file := range files {
				err := g.WriterFile(ctx, filepath.Join(filename, file.Name()))
				if err != nil && err != ErrUnsupportedFileType {
					return err
				}
			}
			return nil
		}
	}
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".7z":
		var dirExists = true
		var md5Bytes = md5.Sum([]byte(filepath.Base(filename)))
		var baseDir = filepath.Join(filepath.Dir(filename), hex.EncodeToString(md5Bytes[0:md5.Size-1]))
		if f, err := os.Stat(baseDir); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
			dirExists = false
		} else if !f.IsDir() {
			return &os.PathError{Op: "StatFile", Path: baseDir, Err: os.ErrExist}
		}
		if !dirExists {
			fullFilename := filename
			if !filepath.IsAbs(fullFilename) {
				fullFilename, _ = filepath.Abs(fullFilename)
			}
			fullBaseDir := filepath.Join(filepath.Dir(fullFilename), hex.EncodeToString(md5Bytes[0:md5.Size-1]))
			if err := zip7.Extract(fullFilename, fullBaseDir); err != nil {
				return Err7zipExtractError
			}
		}
		if err := g.WriterFile(ctx, baseDir); err != nil {
			return err
		}
	case ".csv":
		f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		reader := csv.NewReader(f)
		defer func() {
			f.Close()
		}()
		var idx = 0
		var readerFileType = ReaderFileTypeUnknow
		for {
			idx++
			datas, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			if idx == 1 {
				if datas[0][0] == 0xef && datas[0][1] == 0xbb && datas[0][2] == 0xbf {
					datas[0] = datas[0][3:]
				}
				for k, v := range datas {
					v = strings.ToLower(strings.TrimSpace(v))
					datas[k] = v
				}
				headerStr := strings.Join(datas, ",")
				switch headerStr {
				case areaHeaderStr:
					readerFileType = ReaderFileTypeArea
				case geoHeaderStr:
					readerFileType = ReaderFileTypeGeo
					if err := g.areaPolygonUsecase.Truncate(ctx); err != nil {
						return err
					}
				default:
					return ErrUnsupportedSheetFile
				}
				continue
			}
			switch readerFileType {
			case ReaderFileTypeArea:
				if err := g.WriterByAreaData(ctx, datas); err != nil {
					return err
				}
			case ReaderFileTypeGeo:
				if err := g.WriterByGeoData(ctx, datas); err != nil {
					return err
				}
			default:
			}

		}
	default:
		return ErrUnsupportedFileType
	}
	return nil
}

func NewGithubRepoUsecase(
	repo GithubRepo,
	app *app.App,
	filesystem *filesystem.FileSystem,
	transaction transaction.Transaction,
	spliderUsecase *splider.SpliderUsecase,
	assetUsecase *asset.AssetUsecase,
	areaUsecase *area.AreaUsecase,
	areaPolygonUsecase *polygon.PolygonUsecase,
) *GithubUsecase {
	var account = Account{
		User: "xiangyuecn",
		Repo: "AreaCity-JsSpider-StatsGov",
	}
	return &GithubUsecase{
		account:            account,
		repo:               repo,
		app:                app,
		filesystem:         filesystem,
		transaction:        transaction,
		spliderUsecase:     spliderUsecase,
		assetUsecase:       assetUsecase,
		areaUsecase:        areaUsecase,
		areaPolygonUsecase: areaPolygonUsecase,
	}
}
