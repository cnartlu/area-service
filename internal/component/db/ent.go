package db

import (
	"context"
	"fmt"
	"time"

	"ariga.io/entcache"
	"entgo.io/ent/dialect/sql"
	"github.com/cnartlu/area-service/component/db"
	"github.com/cnartlu/area-service/component/log"
	"github.com/cnartlu/area-service/component/redis"
	"github.com/cnartlu/area-service/internal/data/ent"
	"go.uber.org/zap"
)

func NewEnt(c *db.Config, l *log.Logger) (*ent.Client, func(), error) {
	dsn := db.NewSource(c)
	driver, err := sql.Open(dsn.Dialect(), dsn.Source())
	if err != nil {
		return nil, nil, err
	}
	if err := driver.DB().Ping(); err != nil {
		_ = driver.Close()
		return nil, nil, err
	}
	options := []entcache.Option{
		entcache.Hash(func(query string, args []any) (entcache.Key, error) {
			v, err := entcache.DefaultHash(query, args)
			if err != nil {
				return 0, err
			}
			return fmt.Sprintf("area.ent.%d", v.(uint64)), nil
		}),
	}
	options = append(options, entcache.TTL(3600*time.Second))
	var cleanups = []func(){}

	if c.GetCache() != nil {
		var rdc = c.GetCacheRedis()
		if c.GetCacheAppComponentName() != "" {
			// 加载配置项
		}
		if rdc != nil {
			rds, cleanup1, err := redis.New(rdc, l)
			if err != nil {
				_ = driver.Close()
				return nil, nil, err
			}
			options = append(options, entcache.Levels(entcache.NewLRU(0), entcache.NewRedis(rds)))
			cleanups = append(cleanups, cleanup1)
		}
	}

	ctx := context.Background()
	entcache.NewContext(ctx)

	client := ent.NewClient(
		ent.Debug(),
		ent.Driver(entcache.NewDriver(driver, options...)),
		ent.Log(NewLogFunc(l)),
	)

	cleanups = append(cleanups, func() {
		if err := client.Close(); err != nil {
			l.Error("[ent] client close failed", zap.Error(err))
		}
	})

	var cleanup = func() {
		for _, fn := range cleanups {
			fn()
		}
	}

	return client, cleanup, nil
}

// 默认配置
// driver.DB().SetMaxOpenConns(100)
// driver.DB().SetMaxIdleConns(10)
// // 连接的最大生命周期
// driver.DB().SetConnMaxLifetime(0)
// driver.DB().SetConnMaxIdleTime(time.Second * 60 * 60)
