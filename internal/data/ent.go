package data

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"

	"ariga.io/entcache"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/pkg/component/db"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	
)

type entConfig struct {
	l klog.Logger
}

func NewEnt(dsn db.Dsn, options ...db.HandlerFunc) (*ent.Client, func(), error) {
	driver, err := sql.Open(dsn.Dialect(), dsn.Source())
	if err != nil {
		return nil, nil, err
	}

	var c = &entConfig{}
	if len(options) > 0 {
		for _, fn := range options {
			fn := fn
			fn(c)
		}
	}

	cleanup := func() {
		// logger.Info("closing the ent resources")
		err = driver.Close()
		if err != nil {
			// logger.Error("close db resources failed", zap.Error(err))
		}
	}

	drv := entcache.NewDriver(
		driver,
		entcache.Hash(func(query string, args []any) (entcache.Key, error) {
			v, err := entcache.DefaultHash(query, args)
			if err != nil {
				return 0, err
			}
			return fmt.Sprintf("area.ent.%d", v.(uint64)), nil
		}),
		entcache.TTL(3600*time.Second),
		entcache.Levels(
			entcache.NewLRU(180),
			entcache.NewRedis(redis.NewClient(&redis.Options{
				Addr: "127.0.0.1:6379",
			})),
		),
	)

	entlog := &Logger{l: c.l}
	client := ent.NewClient(
		ent.Debug(),
		ent.Driver(drv),
		ent.Log(func(i ...interface{}) {
			entlog.DebugLog(i)
		}),
	)

	// err = client.Schema.Create(
	// 	entcache.Skip(context.Background()),
	// 	migrate.WithForeignKeys(false),
	// )
	// if err != nil {
	// 	return client, cleanup, err
	// }

	return client, cleanup, nil
}
