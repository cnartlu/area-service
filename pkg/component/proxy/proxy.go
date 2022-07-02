package proxy

import (
	"net/http"
	"net/url"
)

func New(config *Config) *http.Client {
	var client = http.Client{}
	if config != nil {
		if config.GetUrl() != "" {
			//"http://127.0.0.1:7890"
			uri, err := url.Parse(config.GetUrl())
			if err != nil {
				panic(err)
			}
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(uri),
			}
		}
	}
	return &client
}
