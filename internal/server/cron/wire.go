//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cron

import (
	"github.com/cnartlu/area-service/internal/server/cron/job"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	job.ProviderSet,
	NewServer,
)
