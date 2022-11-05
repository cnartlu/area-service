package component

import (
	"github.com/cnartlu/area-service/component/discovery"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/component/redis"
	"github.com/cnartlu/area-service/component/trace"
	"github.com/cnartlu/area-service/component/uid"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.NewSet(log.New),
	wire.NewSet(redis.New),
	wire.NewSet(trace.New),
	wire.NewSet(discovery.New),
	wire.NewSet(wire.Bind(new(uid.Generator), new(*uid.Uid)), uid.New),
)
