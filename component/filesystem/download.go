package filesystem

import (
	"context"
	"io"
	"net/url"
	"os"
	"path/filepath"
)

func (f *FileSystem) Download(ctx context.Context, downloadUrl, filename string) error {
	_, err := url.Parse(downloadUrl)
	if err != nil {
		return err
	}
	filePath := filepath.Dir(filename)
	if fstat, err1 := os.Stat(filePath); err1 != nil {
		if !os.IsNotExist(err1) {
			return err1
		}
		if err1 := os.MkdirAll(filePath, os.ModeDir); err1 != nil {
			return err1
		}
	} else if !fstat.IsDir() {
		return os.ErrInvalid
	}
	resp, err := f.httpClient.Get(downloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		out.Close()
		os.Remove(filename)
		return err
	}
	defer out.Close()
	return nil
}
