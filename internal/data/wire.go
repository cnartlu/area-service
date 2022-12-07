//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package data

import (
	bizarea "github.com/cnartlu/area-service/internal/biz/area"
	bizarearealease "github.com/cnartlu/area-service/internal/biz/area/release"
	bizarearealeaseasset "github.com/cnartlu/area-service/internal/biz/area/release/asset"
	bizgithub "github.com/cnartlu/area-service/internal/biz/github"
	biztransaction "github.com/cnartlu/area-service/internal/biz/transaction"
	"github.com/cnartlu/area-service/internal/data/area"
	"github.com/cnartlu/area-service/internal/data/area/release"
	"github.com/cnartlu/area-service/internal/data/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/github"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	data.NewData,
	wire.Bind(new(biztransaction.Transaction), new(*data.Data)),
	github.NewXiangyuecnRepo,
	wire.Bind(new(bizgithub.XiangyuecnRepository), new(*github.XiangyuecnRepo)),
	area.NewAreaRepo,
	wire.Bind(new(bizarea.AreaRepo), new(*area.AreaRepo)),
	release.NewRepository,
	wire.Bind(new(bizarearealease.ReleaseRepo), new(*release.ReleaseRepo)),
	asset.NewAssetRepo,
	wire.Bind(new(bizarearealeaseasset.AssetRepo), new(*asset.AssetRepo)),
)
