package proxy

import (
	"net/http"
	"net/url"

	"github.com/cnartlu/area-service/component/app"
)

type Client struct {
	*http.Client
}

func New(proxyUrl string) (*Client, error) {
	var proxy *url.URL
	if proxyUrl != "" {
		var err error
		proxy, err = url.Parse(proxyUrl)
		if err != nil {
			return nil, err
		}
	}
	var p = &http.Client{
		Transport: &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) {
				return proxy, nil
			},
		},
	}
	c := Client{p}
	return &c, nil
}

func NewByAppConfig(a *app.Config) (*Client, error) {
	var proxyUrl = a.GetProxy()
	return New(proxyUrl)
}
