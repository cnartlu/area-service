package component

import (
	"github.com/cnartlu/area-service/pkg/component/casbin"
	"github.com/cnartlu/area-service/pkg/component/discovery"
	"github.com/cnartlu/area-service/pkg/component/redis"
	"github.com/cnartlu/area-service/pkg/component/trace"
	"github.com/cnartlu/area-service/pkg/component/uid"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.NewSet(redis.New),
	wire.NewSet(trace.New),
	wire.NewSet(discovery.New),
	wire.NewSet(casbin.New),
	wire.NewSet(wire.Bind(new(uid.Generator), new(*uid.Uid)), uid.New),
)
