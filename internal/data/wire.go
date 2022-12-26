//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package data

import (
	bizarea "github.com/cnartlu/area-service/internal/biz/area"
	bizcitygithub "github.com/cnartlu/area-service/internal/biz/city/github"
	bizcitysplider "github.com/cnartlu/area-service/internal/biz/city/splider"
	bizcityspliderarea "github.com/cnartlu/area-service/internal/biz/city/splider/area"
	bizcityspliderasset "github.com/cnartlu/area-service/internal/biz/city/splider/asset"
	biztransaction "github.com/cnartlu/area-service/internal/biz/transaction"
	"github.com/cnartlu/area-service/internal/data/area"
	citysplider "github.com/cnartlu/area-service/internal/data/city/splider"
	cityspliderarea "github.com/cnartlu/area-service/internal/data/city/splider/area"
	cityspliderasset "github.com/cnartlu/area-service/internal/data/city/splider/asset"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/github"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	data.NewData,
	wire.Bind(new(biztransaction.Transaction), new(*data.Data)),
	github.NewGithubRepo,
	wire.Bind(new(bizcitygithub.GithubRepo), new(*github.GithubRepo)),
	area.NewAreaRepo,
	wire.Bind(new(bizarea.AreaRepo), new(*area.AreaRepo)),
	citysplider.NewSpliderRepo,
	wire.Bind(new(bizcitysplider.SpliderRepo), new(*citysplider.SpliderRepo)),
	cityspliderasset.NewAssetRepo,
	wire.Bind(new(bizcityspliderasset.AssetRepo), new(*cityspliderasset.AssetRepo)),
	cityspliderarea.NewAreaRepo,
	wire.Bind(new(bizcityspliderarea.AreaRepo), new(*cityspliderarea.AreaRepo)),
)
