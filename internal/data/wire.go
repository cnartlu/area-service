//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package data

import (
	bizarea "github.com/cnartlu/area-service/internal/biz/area"
	bizarearealease "github.com/cnartlu/area-service/internal/biz/area/release"
	bizarearealeaseasset "github.com/cnartlu/area-service/internal/biz/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/area"
	"github.com/cnartlu/area-service/internal/data/area/release"
	"github.com/cnartlu/area-service/internal/data/area/release/asset"
	"github.com/cnartlu/area-service/internal/data/github"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	// 这里应该加入 db 存储系统
	github.NewXiangyuecnRepo,
	wire.Bind(new(bizarearealease.XiangyuecnRepository), new(*github.XiangyuecnRepo)),
	area.NewAreaRepo,
	wire.Bind(new(bizarea.Manager), new(*area.AreaRepo)),
	release.NewRepository,
	wire.Bind(new(bizarearealease.ManageRepo), new(*release.ReleaseRepo)),
	asset.NewAssetRepo,
	wire.Bind(new(bizarearealeaseasset.ManageRepo), new(*asset.AssetRepo)),
)
