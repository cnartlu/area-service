package us3

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"strconv"

	"github.com/cnartlu/area-service/pkg/component/filesystem/storage"
)

// Us3 Ucloud对象存储
type Us3 struct {
	// PublicKey 公钥
	PublicKey string
	// PrivateKey 私钥
	PrivateKey string
	// Region 区域
	Region string
	// 存储桶名称
	BucketName string
	// ProjectID 项目
	ProjectId string
	// Type 类型，public公共读,private私有读
	Type string
}

func (u *Us3) Name() string {
	return "us3"
}

// PrefixFileList 获取Bucket中指定文件前缀的文件列表
// prefix 前缀，utf-8编码，默认为空字符串
// marker 标志字符串，utf-8编码，默认为空字符串)
// limit 文件列表数目，默认为20
func (u *Us3) PrefixFileList(ctx context.Context, prefix, marker string, limit int) (*PrefixFileList, error) {
	const path = "/"
	v := urlpkg.Values{}
	v.Set("list", "")
	if prefix != "" {
		v.Set("prefix", prefix)
	}
	if marker != "" {
		v.Set("marker", marker)
	}
	if limit > 0 {
		v.Set("limit", strconv.FormatInt(int64(limit), 10))
	}
	data := PrefixFileList{}
	middleware := func(c *storage.Context) {
		res := c.Next()
		bs, err := io.ReadAll(res.Body)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		err = json.Unmarshal(bs, &data)
		if err != nil {
			c.AbortWithError(err)
			return
		}
	}
	err := u.Request(ctx, http.MethodGet, fmt.Sprintf("/%s?%s", path, v.Encode()), nil, middleware)
	return &data, err
}

// Headile 查询文件基本信息
func (u *Us3) HeadFile(ctx context.Context, key string, middlewares ...storage.HandlerFunc) error {
	middleware := func(c *storage.Context) {
		res := c.Next()
		if res.ContentLength <= 0 {
			return
		}
		if err := newResponseError(res.Body); err != nil {
			c.AbortWithError(err)
		}
	}
	middlewares = append(middlewares, middleware)
	err := u.Request(ctx, http.MethodHead, fmt.Sprintf("/%s", key), nil, middlewares...)
	return err
}

// Headile 查询文件基本信息
func (u *Us3) DeleteFile(ctx context.Context, key string, middlewares ...storage.HandlerFunc) error {
	middleware := func(c *storage.Context) {
		res := c.Next()
		if res.ContentLength <= 0 {
			return
		}
		if err := newResponseError(res.Body); err != nil {
			c.AbortWithError(err)
		}
	}
	middlewares = append(middlewares, middleware)
	err := u.Request(ctx, http.MethodDelete, fmt.Sprintf("/%s", key), nil, middlewares...)
	return err
}

// Request 发起请求
// ctx context.Context 请求上下文
func (u *Us3) Request(ctx context.Context, method string, url string, body io.Reader, handlers ...storage.HandlerFunc) error {
	uri, err := urlpkg.Parse(url)
	if err != nil {
		return err
	}
	if uri.Scheme == "" {
		uri.Scheme = "https"
	}
	if uri.Host == "" {
		uri.Host = fmt.Sprintf("%s.%s.ufileos.com", u.BucketName, u.Region)
	}
	req, err := http.NewRequestWithContext(ctx, method, uri.String(), body)
	if err != nil {
		return err
	}
	// 初始化中间件
	c := storage.Context{
		Request: req,
	}
	if err := c.CheckAbortIndex(len(handlers)); err != nil {
		return fmt.Errorf("storage/us3: %w", err)
	}
	c.Use(handlers...)
	// 关闭请求
	var res *http.Response
	c.Use(func(c *storage.Context) {
		// 是否需要判断是否需要开启认证
		if c.EnableAuthorization() {
			c.Request.Header.Add("Authorization", Signature(u.PrivateKey, u.PublicKey, u.BucketName, c.Request))
		}
		client := http.DefaultClient
		res, err = client.Do(c.Request)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		// 设置请求头
		storage.SetResponse(res)(c)
	})
	// 关闭读取器
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	// 启动中间件
	return c.Run()
}
