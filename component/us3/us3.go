package us3

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"
)

// Signature 计算签名
func Signature(privateKey, publicKey, bucketName string, req *http.Request) string {
	b := bytes.Buffer{}
	b.WriteString("UCloud ")
	b.WriteString(publicKey)
	b.WriteString(":")
	signBuffer := bytes.Buffer{}
	signBuffer.WriteString(req.Method)
	signBuffer.WriteString("\n")
	// Content-MD5
	signBuffer.WriteString(req.Header.Get("Content-MD5"))
	signBuffer.WriteString("\n")
	// Content-Type
	signBuffer.WriteString(req.Header.Get("Content-Type"))
	signBuffer.WriteString("\n")
	// Date
	req.Header.Set("Date", time.Now().String())
	signBuffer.WriteString(req.Header.Get("Date"))
	signBuffer.WriteString("\n")
	// Headers
	headerKeys := []string{}
	headerValues := map[string][]string{}
	headerBuffer := bytes.Buffer{}
	for k, v := range req.Header {
		k1 := strings.ToLower(k)
		if strings.HasSuffix(k1, "x-ucloud-") {
			if _, ok := headerValues[k1]; !ok {
				headerValues[k1] = []string{}
				headerKeys = append(headerKeys, k1)
			}
			headerValues[k1] = append(headerValues[k1], v...)
		}
	}
	if len(headerKeys) > 0 {
		sort.Strings(headerKeys)
		for _, k := range headerKeys {
			headerBuffer.WriteString(k)
			headerBuffer.WriteString(":")
			headerBuffer.WriteString(strings.Join(headerValues[k], ","))
			headerBuffer.WriteString("\n")
		}
	}
	signBuffer.Write(headerBuffer.Bytes())
	signBuffer.WriteString("/")
	signBuffer.WriteString(bucketName)
	if req.URL.Path == "" {
		signBuffer.WriteString("/")
	} else {
		signBuffer.WriteString(req.URL.Path)
	}
	h := hmac.New(sha1.New, []byte(privateKey))
	h.Write(signBuffer.Bytes())
	hexValue := h.Sum(nil)
	b.WriteString(base64.StdEncoding.EncodeToString(hexValue))

	signture := b.String()
	return signture
}

type Error struct {
	RetCode int    `json:"RetCode"`
	ErrMsg  string `json:"ErrMsg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("err code %d, message: %s", e.RetCode, e.ErrMsg)
}

// DataSetItem 文件数据项
type DataSetItem struct {
	// 文件所属Bucket名称
	BucketName string
	// 文件名称,utf-8编码
	FileName string
	// 文件hash值
	Hash string
	// 文件mimetype
	MimeType string
	// 文件大小
	Size int64
	// 文件创建时间
	CreateTime int64
	// 文件修改时间
	ModifyTime int64
	// 文件存储类型，分别是标准、低频、归档，对应有效值：STANDARD, IA, ARCHIVE
	StorageClass string
}

// PrefixFileList 前缀文件列表
type PrefixFileList struct {
	// Bucket的名称
	BucketName string
	// Bucket的ID
	BucketId string
	// 下一个标志字符串，utf-8编码
	NextMarker string
	DataSet    []DataSetItem
}

// InitiateMultipartUploadReply 分配上传返回值
type InitiateMultipartUploadReply struct {
	// 本次分片上传的上传Id
	UploadId string
	// 分片的块大小
	BlkSize int
	// 上传文件所属Bucket的名称
	Bucket string
	// 上传文件在Bucket中的Key名称
	Key string
}

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

// PrefixFileList 获取Bucket中指定文件前缀的文件列表
// prefix 前缀，utf-8编码，默认为空字符串
// marker 标志字符串，utf-8编码，默认为空字符串)
// limit 文件列表数目，默认为20
// func (u *Us3) PrefixFileList(ctx context.Context, prefix, marker string, limit int) (*PrefixFileList, error) {
// 	const path = "/"
// 	v := urlpkg.Values{}
// 	v.Set("list", "")
// 	if prefix != "" {
// 		v.Set("prefix", prefix)
// 	}
// 	if marker != "" {
// 		v.Set("marker", marker)
// 	}
// 	if limit > 0 {
// 		v.Set("limit", strconv.FormatInt(int64(limit), 10))
// 	}
// 	data := PrefixFileList{}
// 	middleware := func(c *storage.Context) {
// 		res := c.Next()
// 		bs, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		err = json.Unmarshal(bs, &data)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 	}
// 	err := u.Request(ctx, http.MethodGet, fmt.Sprintf("/%s?%s", path, v.Encode()), nil, middleware)
// 	return &data, err
// }

// // Headile 查询文件基本信息
// func (u *Us3) HeadFile(ctx context.Context, key string, middlewares ...storage.HandlerFunc) error {
// 	middleware := func(c *storage.Context) {
// 		res := c.Next()
// 		if res.ContentLength <= 0 {
// 			return
// 		}
// 		if err := newResponseError(res.Body); err != nil {
// 			c.AbortWithError(err)
// 		}
// 	}
// 	middlewares = append(middlewares, middleware)
// 	err := u.Request(ctx, http.MethodHead, fmt.Sprintf("/%s", key), nil, middlewares...)
// 	return err
// }

// // PutFile 获取Bucket中指定文件前缀的文件列表
// // key 文件存储键
// // body 上传的文件内容
// // storage 文件存储类型，分别是标准、低频、归档，对应有效值：STANDARD, IA, ARCHIVE
// // metas US3中规定所有以X-Ufile-Meta-为前缀的参数视为用户自定义元数据（User Meta），比如x-ufile-meta-location。一个文件可以有多个类似的参数，但所有的User Meta总大小不能超过8KB。这些User Meta信息会在GetFile或者HeadFile的时候在HTTP头部中返回。
// func (u *Us3) PutFile(ctx context.Context, key string, body io.Reader, middlewares ...storage.HandlerFunc) error {
// 	middleware := func(c *storage.Context) {
// 		res := c.Next()
// 		if res.ContentLength <= 0 {
// 			return
// 		}
// 		if err := newResponseError(res.Body); err != nil {
// 			c.AbortWithError(err)
// 		}
// 	}
// 	middlewares = append(middlewares, middleware)
// 	err := u.Request(ctx, http.MethodPut, fmt.Sprintf("/%s", key), body, middlewares...)
// 	return err
// }

// // PutFile 获取Bucket中指定文件前缀的文件列表
// // key 文件存储键
// // body 上传的文件内容
// // metas US3中规定所有以X-Ufile-Meta-为前缀的参数视为用户自定义元数据（User Meta），比如x-ufile-meta-location。一个文件可以有多个类似的参数，但所有的User Meta总大小不能超过8KB。这些User Meta信息会在GetFile或者HeadFile的时候在HTTP头部中返回。
// func (u *Us3) PostFile(ctx context.Context, key string, body io.Reader, middlewares ...storage.HandlerFunc) error {
// 	buffer := &bytes.Buffer{}
// 	writer := multipart.NewWriter(buffer)
// 	writer.WriteField("FileName", key)
// 	// 设置传入的文件
// 	p, err := writer.CreateFormFile("file", filepath.Base(key))
// 	if err != nil {
// 		return err
// 	}
// 	if _, err := io.Copy(p, body); err != nil {
// 		return err
// 	}
// 	middleware := func(c *storage.Context) {
// 		c.Authorization(false)
// 		writer.WriteField("Authorization", Signature(u.PrivateKey, u.PublicKey, u.BucketName, c.Request))
// 		// 关闭写入器
// 		if err := writer.Close(); err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		c.Request.Header.Set("Content-Type", writer.FormDataContentType())
// 		c.Request.ContentLength = int64(buffer.Len())
// 		res := c.Next()
// 		if res.ContentLength <= 0 {
// 			return
// 		}
// 		if err := newResponseError(res.Body); err != nil {
// 			c.AbortWithError(err)
// 		}
// 	}
// 	// 发起请求
// 	err = u.Request(ctx, http.MethodPost, "/", buffer, middleware)
// 	return err
// }

// // UploadHit 说明：先判断待上传文件的hash值，如果US3中可以查到此文件，则不必再传文件本身。
// func (u *Us3) UploadHit(ctx context.Context, key, hash string, size int, body io.Reader, middlewares ...storage.HandlerFunc) error {
// 	middleware := func(c *storage.Context) {
// 		res := c.Next()
// 		if res.ContentLength <= 0 {
// 			return
// 		}
// 		if err := newResponseError(res.Body); err != nil {
// 			c.AbortWithError(err)
// 		}
// 	}
// 	middlewares = append(middlewares, middleware)
// 	err := u.Request(ctx, http.MethodPost, fmt.Sprintf("/uploadhit?Hash=%s&FileName=%s&FileSize=%d", hash, key, size), body, middlewares...)
// 	return err
// }

// // InitiateMultipartUpload 初始化分片上传
// func (u *Us3) InitiateMultipartUpload(ctx context.Context, key string, middlewares ...storage.HandlerFunc) (result *InitiateMultipartUploadReply, err error) {
// 	middleware := func(c *storage.Context) {
// 		res := c.Next()
// 		bs, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		err = json.Unmarshal(bs, result)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 	}
// 	middlewares = append(middlewares, middleware)
// 	err = u.Request(ctx, http.MethodPost, fmt.Sprintf("/%s?uploads", key), nil, middlewares...)
// 	return
// }

// // UploadPart 上传文件分片
// func (u *Us3) UploadPart(ctx context.Context, key, uploadID string, partNumber int, middlewares ...storage.HandlerFunc) (number int, err error) {
// 	middleware := func(c *storage.Context) {
// 		res := c.Next()
// 		bs, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		var result struct {
// 			// 本次分片上传的分片号码
// 			PartNumber int
// 		}
// 		err = json.Unmarshal(bs, &result)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		number = result.PartNumber
// 	}
// 	middlewares = append(middlewares, middleware)
// 	err = u.Request(ctx, http.MethodPut, fmt.Sprintf("/%s?uploadId=%s&partNumber=%d", key, uploadID, partNumber), nil, middlewares...)
// 	return
// }

// // FinishMultipartUpload 完成上传文件分片
// func (u *Us3) FinishMultipartUpload(ctx context.Context, key, uploadID, newKey string, body io.Reader, middlewares ...storage.HandlerFunc) (number int, err error) {
// 	middleware := func(c *storage.Context) {
// 		res := c.Next()
// 		bs, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		var result struct {
// 			// 本次分片上传的分片号码
// 			PartNumber int
// 		}
// 		err = json.Unmarshal(bs, &result)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		number = result.PartNumber
// 	}
// 	middlewares = append(middlewares, middleware)
// 	err = u.Request(ctx, http.MethodPost, fmt.Sprintf("/%s?uploadId=%s&newKey=%s", key, uploadID, newKey), body, middlewares...)
// 	return
// }

// // AbortMultipartUpload 放弃上传文件分片
// func (u *Us3) AbortMultipartUpload(ctx context.Context, key, uploadID string, middlewares ...storage.HandlerFunc) (number int, err error) {
// 	middleware := func(c *storage.Context) {
// 		res := c.Next()
// 		bs, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		var result struct {
// 			// 本次分片上传的分片号码
// 			PartNumber int
// 		}
// 		err = json.Unmarshal(bs, &result)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		number = result.PartNumber
// 	}
// 	middlewares = append(middlewares, middleware)
// 	err = u.Request(ctx, http.MethodDelete, fmt.Sprintf("/%s?uploadId=%s", key, uploadID), nil, middlewares...)
// 	return
// }

// // Headile 查询文件基本信息
// func (u *Us3) DeleteFile(ctx context.Context, key string, middlewares ...storage.HandlerFunc) error {
// 	middleware := func(c *storage.Context) {
// 		res := c.Next()
// 		if res.ContentLength <= 0 {
// 			return
// 		}
// 		if err := newResponseError(res.Body); err != nil {
// 			c.AbortWithError(err)
// 		}
// 	}
// 	middlewares = append(middlewares, middleware)
// 	err := u.Request(ctx, http.MethodDelete, fmt.Sprintf("/%s", key), nil, middlewares...)
// 	return err
// }

// // Request 发起请求
// // ctx context.Context 请求上下文
// func (u *Us3) Request(ctx context.Context, method string, url string, body io.Reader, handlers ...storage.HandlerFunc) error {
// 	uri, err := urlpkg.Parse(url)
// 	if err != nil {
// 		return err
// 	}
// 	if uri.Scheme == "" {
// 		uri.Scheme = "https"
// 	}
// 	if uri.Host == "" {
// 		uri.Host = fmt.Sprintf("%s.%s.ufileos.com", u.BucketName, u.Region)
// 	}
// 	req, err := http.NewRequestWithContext(ctx, method, uri.String(), body)
// 	if err != nil {
// 		return err
// 	}
// 	// 初始化中间件
// 	c := storage.Context{
// 		Request: req,
// 	}
// 	if err := c.CheckAbortIndex(len(handlers)); err != nil {
// 		return fmt.Errorf("storage/us3: %w", err)
// 	}
// 	c.Use(handlers...)
// 	// 关闭请求
// 	var res *http.Response
// 	c.Use(func(c *storage.Context) {
// 		// 是否需要判断是否需要开启认证
// 		if c.EnableAuthorization() {
// 			c.Request.Header.Add("Authorization", Signature(u.PrivateKey, u.PublicKey, u.BucketName, c.Request))
// 		}
// 		client := http.DefaultClient
// 		res, err = client.Do(c.Request)
// 		if err != nil {
// 			c.AbortWithError(err)
// 			return
// 		}
// 		// 设置请求头
// 		storage.SetResponse(res)(c)
// 	})
// 	// 关闭读取器
// 	defer func() {
// 		if res != nil {
// 			res.Body.Close()
// 		}
// 	}()
// 	// 启动中间件
// 	return c.Run()
// }
