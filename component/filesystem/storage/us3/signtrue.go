package us3

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
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
