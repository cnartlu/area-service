package handler

import (
	"github.com/cnartlu/area-service/internal/transport/http/handler/v1/area"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	area.NewHandler,
)
