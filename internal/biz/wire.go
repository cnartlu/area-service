//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package biz

import (
	"github.com/cnartlu/area-service/internal/biz/area"
	citygithub "github.com/cnartlu/area-service/internal/biz/city/github"
	citysplider "github.com/cnartlu/area-service/internal/biz/city/splider"
	cityspliderarea "github.com/cnartlu/area-service/internal/biz/city/splider/area"
	cityspliderasset "github.com/cnartlu/area-service/internal/biz/city/splider/asset"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	area.NewAreaUsecase,
	citygithub.NewGithubRepoUsecase,
	citysplider.NewSpliderUsecase,
	cityspliderarea.NewAreaUsecase,
	cityspliderasset.NewAssetUsecase,
)
