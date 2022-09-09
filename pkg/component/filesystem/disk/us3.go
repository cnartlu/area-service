package disk

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	urlpkg "net/url"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ErrUs3 struct {
	// 返回状态码，为 0 则为成功返回，非 0 为失败
	RetCode int `json:"RetCode"`
	// 返回错误消息，当 RetCode 非 0 时提供详细的描述信息
	ErrMsg    string `json:"ErrMsg"`
	SessionId string `json:"SessionId"`
}

func (e ErrUs3) Error() string {
	// request is `%s` action `%s` request failed,
	return fmt.Sprintf("sessionID %s, err code %d, message: %s", e.SessionId, e.RetCode, e.ErrMsg)
}

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
type PrefixFileList struct {
	// Bucket的名称
	BucketName string
	// Bucket的ID
	BucketId string
	// 下一个标志字符串，utf-8编码
	NextMarker string
	DataSet    []DataSetItem
}

type CreateBucketResult struct {
	ErrUs3
	// 已创建Bucket的名称
	BucketName string
	// 已创建Bucket的ID
	BucketId string
}

type Us3Storage struct {
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

	httpClient *http.Client
}

// PrefixFileList 获取Bucket中指定文件前缀的文件列表
// prefix 前缀，utf-8编码，默认为空字符串
// marker 标志字符串，utf-8编码，默认为空字符串)
// limit 文件列表数目，默认为20
func (u *Us3Storage) PrefixFileList(ctx context.Context, prefix, marker string, limit int) (*PrefixFileList, error) {
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
	middleware := func(c *Context) {
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
	err := u.request(ctx, http.MethodGet, fmt.Sprintf("/%s?%s", path, v.Encode()), nil, middleware)
	return &data, err
}

// PutFile 获取Bucket中指定文件前缀的文件列表
// key 文件存储键
// body 上传的文件内容
// storage 文件存储类型，分别是标准、低频、归档，对应有效值：STANDARD, IA, ARCHIVE
// metas US3中规定所有以X-Ufile-Meta-为前缀的参数视为用户自定义元数据（User Meta），比如x-ufile-meta-location。一个文件可以有多个类似的参数，但所有的User Meta总大小不能超过8KB。这些User Meta信息会在GetFile或者HeadFile的时候在HTTP头部中返回。
func (u *Us3Storage) PutFile(ctx context.Context, key string, body io.Reader, storage string, middlewares ...HandlerFunc) error {
	middleware := func(c *Context) {
		if storage != "" {
			storage = strings.ToUpper(strings.TrimSpace(storage))
		}
		switch storage {
		case "STANDARD", "IA", "ARCHIVE":
		default:
			storage = "STANDARD"
		}
		c.Request.Header.Set("X-Ufile-Storage-Class", storage)
		res := c.Next()
		if res.ContentLength <= 0 {
			return
		}
		bs, err := io.ReadAll(res.Body)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		data := ErrUs3{}
		sessionId := res.Header.Get("X-SessionId")
		err = json.Unmarshal(bs, &data)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		data.SessionId = sessionId
		c.Error(data)
	}
	middlewares = append(middlewares, middleware)
	err := u.request(ctx, http.MethodPut, fmt.Sprintf("/%s", key), body, middlewares...)
	return err
}

// PutFile 获取Bucket中指定文件前缀的文件列表
// key 文件存储键
// body 上传的文件内容
// storage 文件存储类型，分别是标准、低频、归档，对应有效值：STANDARD, IA, ARCHIVE
// metas US3中规定所有以X-Ufile-Meta-为前缀的参数视为用户自定义元数据（User Meta），比如x-ufile-meta-location。一个文件可以有多个类似的参数，但所有的User Meta总大小不能超过8KB。这些User Meta信息会在GetFile或者HeadFile的时候在HTTP头部中返回。
func (u *Us3Storage) PostFile(ctx context.Context, key string, body io.Reader, storage string) error {
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
	middleware := func(c *Context) {
		if storage != "" {
			storage = strings.ToUpper(strings.TrimSpace(storage))
		}
		switch storage {
		case "STANDARD", "IA", "ARCHIVE":
		default:
			storage = "STANDARD"
		}
		writer.WriteField("Authorization", u.Signature(c.Request))
		// 关闭写入器
		if err := writer.Close(); err != nil {
			c.AbortWithError(err)
			return
		}
		c.Request.Header.Set("X-Ufile-Storage-Class", storage)
		c.Request.Header.Set("Content-Type", writer.FormDataContentType())
		c.Request.ContentLength = int64(buffer.Len())
		c.WithAuthorization(false)
		res := c.Next()
		if res.ContentLength <= 0 {
			return
		}
		bs, err := io.ReadAll(res.Body)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		if len(bs) == 0 {
			return
		}
		data := ErrUs3{}
		sessionId := res.Header.Get("X-SessionId")
		err = json.Unmarshal(bs, &data)
		if err != nil {
			c.AbortWithError(err)
			return
		}
		data.SessionId = sessionId
		c.Error(data)
	}
	// 发起请求
	err = u.request(ctx, http.MethodPost, "/", buffer, middleware)
	return err
}

// Signature 计算签名
func (u *Us3Storage) Signature(req *http.Request) string {
	b := bytes.Buffer{}
	b.WriteString("UCloud ")
	b.WriteString(u.PublicKey)
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
	signBuffer.WriteString(u.BucketName)
	if req.URL.Path == "" {
		signBuffer.WriteString("/")
	} else {
		signBuffer.WriteString(req.URL.Path)
	}

	h := hmac.New(sha1.New, []byte(u.PrivateKey))
	h.Write(signBuffer.Bytes())
	hexValue := h.Sum(nil)
	b.WriteString(base64.StdEncoding.EncodeToString(hexValue))

	signture := b.String()
	req.Header.Add("Authorization", signture)
	return signture
}

// request 发起请求
func (u *Us3Storage) request(ctx context.Context, method string, url string, body io.Reader, handlers ...HandlerFunc) error {
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
	handlerLen := len(handlers)
	c := &Context{
		Request:           req,
		response:          nil,
		withAuthorization: true,
		handlers:          HandlersChain{},
		index:             0,
		Errors:            []error{},
	}
	if int8(handlerLen) >= abortIndex {
		return fmt.Errorf("disk/us3: too many handlers")
	}
	if handlerLen > 0 {
		for _, fn := range handlers {
			if fn == nil {
				continue
			}
			c.handlers = append(c.handlers, fn)
		}
	}
	if u.httpClient == nil {
		u.httpClient = http.DefaultClient
	}
	c.handlers = append(c.handlers, func(c *Context) {
		if c.withAuthorization {
			// 签名补全
			if req.Header.Get("Authorization") == "" {
				u.Signature(req)
			}
		}
		c.response, err = u.httpClient.Do(req)
		if err != nil {
			c.AbortWithError(err)
			return
		}
	})
	// 关闭读取器
	defer func() {
		if c.response != nil {
			c.response.Body.Close()
		}
	}()
	// 启动中间件
	if len(c.handlers) > 0 {
		c.handlers[0](c)
		_ = c.Next()
	}
	if len(c.Errors) > 0 {
		return c.Errors
	}
	return nil
}
