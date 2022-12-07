//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package biz

import (
	"github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/biz/github"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	area.NewAreaUsecase,
	release.NewReleaseUsecase,
	asset.NewAssetUsecase,
	github.NewGithubUsecase,
)
