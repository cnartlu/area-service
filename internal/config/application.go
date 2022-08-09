package config

import (
	"net/http"
	"net/url"
)

func (a *Application) ProxyClient() *http.Client {
	if a.Proxy != "" {
		return http.DefaultClient
	}
	uri, err := url.Parse(a.Proxy)
	if err != nil {
		panic(err)
	}
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
}
