package handler

import (
	"github.com/cnartlu/area-service/internal/command/handler/greet"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	greet.NewHandler,
)
