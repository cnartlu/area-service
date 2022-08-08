package utils

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)

// IsFile 是否是文件路径
func IsFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return !f.IsDir()
}

type HttpClient struct {
	client *http.Client
}

var httpClient = &HttpClient{
	client: http.DefaultClient,
}

func (h *HttpClient) Download(ctx context.Context, downloadUrl, filename string) error {
	_, err := url.Parse(downloadUrl)
	if err != nil {
		return err
	}
	resp, err := h.client.Get(downloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		defer os.Remove(filename)
		return err
	}
	return nil
}

func Download(ctx context.Context, url, filename string) error {
	return httpClient.Download(ctx, url, filename)
}
