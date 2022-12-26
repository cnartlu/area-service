package github

import (
	"context"
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"

	zip7 "github.com/cnartlu/area-service/component/7zip"
	"github.com/cnartlu/area-service/errors"
)

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

func (g *GithubUsecase) ReaderDir(ctx context.Context, path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return errors.ErrorServerError("os read dir error").WithCause(err).WithMetadata(map[string]string{
			"path": path,
		})
	}
	for _, file := range files {
		var err error
		if file.IsDir() {
			err = g.ReaderDir(ctx, filepath.Join(path, file.Name()))
		} else {
			err = g.ReaderFile(ctx, filepath.Join(path, file.Name()))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GithubUsecase) ReaderFile(ctx context.Context, filename string) error {
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".7z":
		if !filepath.IsAbs(filename) {
			filename, _ = filepath.Abs(filename)
		}
		basename := filepath.Base(filename)
		b := md5.Sum([]byte(basename))
		basedir := filepath.Join(filepath.Dir(filename), hex.EncodeToString(b[0:md5.Size-1]))
		if err := zip7.Extract(filename, basedir); err != nil {
			return Err7zipExtractError.WithCause(err).WithMetadata(map[string]string{
				"filename": filename,
			})
		}
		if err := g.ReaderDir(ctx, basedir); err != nil {
			return err
		}
	case ".csv":
		f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return errors.ErrorServerError("open file error").WithCause(err).WithMetadata(map[string]string{
				"filename": filename,
			})
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
				// if err := g.WriterByGeoData(ctx, datas); err != nil {
				// 	return err
				// }
			default:
			}

		}
	default:
		return ErrUnsupportedFileType.WithMetadata(map[string]string{
			"filename": filename,
		})
	}
	return nil
}
