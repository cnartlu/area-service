package proxy

import (
	"net/http"
	"net/url"
)

// New 实例化代理客户端
// http 代理 http://127.0.0.1:7890
// socket5 socket5://127.0.0.1:8000
func New(proxyUrl string) *http.Client {
	var client = http.Client{}
	if proxyUrl != "" {
		uri, err := url.Parse(proxyUrl)
		if err != nil {
			panic(err)
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(uri),
		}
	}
	return &client
}
