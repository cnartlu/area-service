package us3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/cnartlu/area-service/pkg/component/filesystem/storage"
)

// PutFile 获取Bucket中指定文件前缀的文件列表
// key 文件存储键
// body 上传的文件内容
// storage 文件存储类型，分别是标准、低频、归档，对应有效值：STANDARD, IA, ARCHIVE
// metas US3中规定所有以X-Ufile-Meta-为前缀的参数视为用户自定义元数据（User Meta），比如x-ufile-meta-location。一个文件可以有多个类似的参数，但所有的User Meta总大小不能超过8KB。这些User Meta信息会在GetFile或者HeadFile的时候在HTTP头部中返回。
func (u *Us3) PutFile(ctx context.Context, key string, body io.Reader, middlewares ...storage.HandlerFunc) error {
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
	err := u.Request(ctx, http.MethodPut, fmt.Sprintf("/%s", key), body, middlewares...)
	return err
}

// PutFile 获取Bucket中指定文件前缀的文件列表
// key 文件存储键
// body 上传的文件内容
// metas US3中规定所有以X-Ufile-Meta-为前缀的参数视为用户自定义元数据（User Meta），比如x-ufile-meta-location。一个文件可以有多个类似的参数，但所有的User Meta总大小不能超过8KB。这些User Meta信息会在GetFile或者HeadFile的时候在HTTP头部中返回。
func (u *Us3) PostFile(ctx context.Context, key string, body io.Reader, middlewares ...storage.HandlerFunc) error {
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	writer.WriteField("FileName", key)
	// 设置传入的文件
	p, err := writer.CreateFormFile("file", filepath.Base(key))
	if err != nil {
		return err
	}
	if _, err := io.Copy(p, body); err != nil {
		return err
	}
	middleware := func(c *storage.Context) {
		c.Authorization(false)
		writer.WriteField("Authorization", Signature(u.PrivateKey, u.PublicKey, u.BucketName, c.Request))
		// 关闭写入器
		if err := writer.Close(); err != nil {
			c.AbortWithError(err)
			return
		}
		c.Request.Header.Set("Content-Type", writer.FormDataContentType())
		c.Request.ContentLength = int64(buffer.Len())
		res := c.Next()
		if res.ContentLength <= 0 {
			return
		}
		if err := newResponseError(res.Body); err != nil {
			c.AbortWithError(err)
		}
	}
	// 发起请求
	err = u.Request(ctx, http.MethodPost, "/", buffer, middleware)
	return err
}

// UploadHit 说明：先判断待上传文件的hash值，如果US3中可以查到此文件，则不必再传文件本身。
func (u *Us3) UploadHit(ctx context.Context, key, hash string, size int, body io.Reader, middlewares ...storage.HandlerFunc) error {
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
	err := u.Request(ctx, http.MethodPost, fmt.Sprintf("/uploadhit?Hash=%s&FileName=%s&FileSize=%d", hash, key, size), body, middlewares...)
	return err
}

// InitiateMultipartUpload 初始化分片上传
func (u *Us3) InitiateMultipartUpload(ctx context.Context, key string, middlewares ...storage.HandlerFunc) (result *InitiateMultipartUploadReply, err error) {
	middleware := func(c *storage.Context) {
		res := c.Next()
		bs, err := io.ReadAll(res.Body)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		err = json.Unmarshal(bs, result)
		if err != nil {
			c.AbortWithError(err)
			return
		}
	}
	middlewares = append(middlewares, middleware)
	err = u.Request(ctx, http.MethodPost, fmt.Sprintf("/%s?uploads", key), nil, middlewares...)
	return
}

// UploadPart 上传文件分片
func (u *Us3) UploadPart(ctx context.Context, key, uploadID string, partNumber int, middlewares ...storage.HandlerFunc) (number int, err error) {
	middleware := func(c *storage.Context) {
		res := c.Next()
		bs, err := io.ReadAll(res.Body)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		var result struct {
			// 本次分片上传的分片号码
			PartNumber int
		}
		err = json.Unmarshal(bs, &result)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		number = result.PartNumber
	}
	middlewares = append(middlewares, middleware)
	err = u.Request(ctx, http.MethodPut, fmt.Sprintf("/%s?uploadId=%s&partNumber=%d", key, uploadID, partNumber), nil, middlewares...)
	return
}

// FinishMultipartUpload 完成上传文件分片
func (u *Us3) FinishMultipartUpload(ctx context.Context, key, uploadID, newKey string, body io.Reader, middlewares ...storage.HandlerFunc) (number int, err error) {
	middleware := func(c *storage.Context) {
		res := c.Next()
		bs, err := io.ReadAll(res.Body)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		var result struct {
			// 本次分片上传的分片号码
			PartNumber int
		}
		err = json.Unmarshal(bs, &result)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		number = result.PartNumber
	}
	middlewares = append(middlewares, middleware)
	err = u.Request(ctx, http.MethodPost, fmt.Sprintf("/%s?uploadId=%s&newKey=%s", key, uploadID, newKey), body, middlewares...)
	return
}

// AbortMultipartUpload 放弃上传文件分片
func (u *Us3) AbortMultipartUpload(ctx context.Context, key, uploadID string, middlewares ...storage.HandlerFunc) (number int, err error) {
	middleware := func(c *storage.Context) {
		res := c.Next()
		bs, err := io.ReadAll(res.Body)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		var result struct {
			// 本次分片上传的分片号码
			PartNumber int
		}
		err = json.Unmarshal(bs, &result)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		number = result.PartNumber
	}
	middlewares = append(middlewares, middleware)
	err = u.Request(ctx, http.MethodDelete, fmt.Sprintf("/%s?uploadId=%s", key, uploadID), nil, middlewares...)
	return
}

// Upload 上传文件
func (u *Us3) Upload(ctx context.Context, key string, data io.Reader, middlewares ...storage.HandlerFunc) error {
	return u.PostFile(ctx, key, data, middlewares...)
}
