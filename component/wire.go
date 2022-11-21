package component

import (
	"github.com/cnartlu/area-service/component/app"
	"github.com/cnartlu/area-service/component/config"
	"github.com/cnartlu/area-service/component/discovery"
	"github.com/cnartlu/area-service/component/github"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/component/proxy"
	"github.com/cnartlu/area-service/component/redis"
	"github.com/cnartlu/area-service/component/trace"
	"github.com/cnartlu/area-service/component/uid"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.NewSet(config.NewKratos),
	wire.NewSet(app.New),
	wire.NewSet(proxy.NewByAppConfig),
	wire.NewSet(log.New),
	wire.NewSet(redis.New),
	wire.NewSet(github.New),
	wire.NewSet(trace.New),
	wire.NewSet(discovery.New),
	wire.NewSet(wire.Bind(new(uid.Generator), new(*uid.Uid)), uid.New),
)
