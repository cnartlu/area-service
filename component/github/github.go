package github

import (
	"net/http"

	"github.com/google/go-github/v45/github"

	"github.com/cnartlu/area-service/component/proxy"
)

func New(p *proxy.Client) *github.Client {
	var h *http.Client
	if p == nil {
		h = p.Client
	}
	c := github.NewClient(h)
	return c
}
