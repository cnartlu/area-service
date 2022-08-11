package handler

import (
	"github.com/cnartlu/area-service/internal/command/handler/config"
	"github.com/cnartlu/area-service/internal/command/handler/greet"
	"github.com/cnartlu/area-service/internal/command/handler/sync"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	greet.NewHandler,
	sync.NewHandler,
	config.NewHandler,
)
