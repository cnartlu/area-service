package release

import (
	"context"
	"testing"

	"github.com/cnartlu/area-service/internal/repository/area/release"
	"github.com/cnartlu/area-service/internal/repository/area/release/asset"
	"github.com/cnartlu/area-service/internal/service/area/release/asset/importer"
	"github.com/cnartlu/area-service/tests"
)

func Test_Import(t *testing.T) {
	c, cleanup, err := tests.Init()
	if err != nil {
		t.Error(err)
		return
	}
	t.Cleanup(cleanup)
	releaseRepo := release.NewRepository(
		c.Ent,
		c.Redis,
	)
	assetRepo := asset.NewRepository(
		c.Ent,
		c.Redis,
	)
	importService := importer.NewService(
		c.Logger,
		c.Config.Bootstrap.Application,
	)
	ctx := context.Background()
	service := NewService(c.Logger, releaseRepo, assetRepo, importService)
	areaRelease, err := service.FindOneWithLastAdded(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	err = service.Import(ctx, areaRelease)
	if err != nil {
		t.Error(err)
	}
}
