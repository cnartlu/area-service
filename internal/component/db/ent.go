package db

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/pkg/log"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	// 通过MySQL驱动使用Opencensus​
	_ "github.com/go-sql-driver/mysql"
	// 使用pgx驱动PostgreSQL​
	_ "github.com/jackc/pgx/v4/stdlib"
)

type entLogger struct {
	*log.Logger
}

// DebugLog 实现ent的日志记录器方法
func (l *entLogger) DebugLog(keyvals ...interface{}) {
	length := len(keyvals)
	switch length {
	case 0:
	case 1:
		l.AddCallerSkip(1).Debug(fmt.Sprint(keyvals[0]))
	default:
		var (
			msg  string
			data []zap.Field
		)
		if length%2 == 0 {
			for i := 0; i < len(keyvals); i += 2 {
				data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
			}
		} else {
			for i := 1; i < len(keyvals); i += 2 {
				data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
			}
		}
		l.AddCallerSkip(1).Debug(msg, data...)
	}
}

// WithTx 数据库事务
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

// NewEnt 实例化数据库客户端
func NewEnt(config *Config, logger *log.Logger) (*ent.Client, func(), error) {
	if config == nil {
		return nil, func() {}, fmt.Errorf("db ent config is nil")
	}
	switch config.Driver {
	case dialect.SQLite:
		config.Driver = dialect.SQLite
	case dialect.MySQL:
		config.Driver = dialect.MySQL
	case dialect.Gremlin:
		config.Driver = dialect.Gremlin
	case dialect.Postgres:
		config.Driver = dialect.Postgres
	default:
		config.Driver = dialect.MySQL
	}

	driver, err := sql.Open(config.Driver, buildSource(config))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ent open fail")
	}
	// 默认配置
	driver.DB().SetMaxOpenConns(100)
	driver.DB().SetMaxIdleConns(10)
	// 连接的最大生命周期
	driver.DB().SetConnMaxLifetime(0)
	driver.DB().SetConnMaxIdleTime(time.Second * 60 * 60)

	cleanup := func() {
		logger.Info("[ent] db client stopping")
		err = driver.Close()
		if err != nil {
			logger.Error("[ent] db client stopping error", zap.Error(err))
		}
	}

	if err := driver.DB().Ping(); err != nil {
		defer cleanup()
		logger.Error("ping db resources failed", zap.Error(err))
		return nil, nil, err
	}

	// 创建缓存
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

	entlog := entLogger{Logger: logger}
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

func buildSource(c *Config) string {
	var (
		options string
		dsn     = ""
	)
	if c == nil {
		return dsn
	}
	if c.Source != "" {
		return c.Source
	}
	if c.Charset == "" {
		c.Charset = "utf8"
	}
	switch c.Driver {
	case dialect.Postgres:
		for k, v := range c.Params {
			options += k + "=" + v.String() + " "
		}
		dsn = "host=" + c.Hostname + " port=" + strconv.Itoa(int(c.Hostport)) + " user=" + c.Username + " password=" + c.Password + " dbname=" + c.Database + " " + options
	case dialect.MySQL:
		fallthrough
	default:
		for k, v := range c.Params {
			options += k + "=" + v.String() + "&"
		}
		options += "charset=" + c.Charset + "&"
		dsn = c.Username + ":" + c.Password + "@tcp(" + c.Hostname + ":" + strconv.Itoa(int(c.Hostport)) + ")/" + c.Database + "?" + options
	}

	return dsn
}
