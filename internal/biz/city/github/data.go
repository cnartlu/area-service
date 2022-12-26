package github

import (
	"github.com/cnartlu/area-service/internal/biz/city/splider"
	"github.com/cnartlu/area-service/internal/biz/city/splider/asset"
)

const Source = splider.SourceGithub

type Account struct {
	User string
	Repo string
}

type Github struct {
	*splider.Splider
	Assets []*asset.Asset
}
