package data

import (
	"context"
	"fmt"
	"strings"
	"time"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/cnartlu/area-service/component/database"
	"github.com/cnartlu/area-service/internal/data/ent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerEnt struct {
	logger *zap.Logger
}

// DebugLog 实现ent的日志记录器方法
func (l loggerEnt) Log(level zapcore.Level, keyvals ...interface{}) {
	var data []zap.Field
	var msg string
	var length = len(keyvals)
	switch length {
	case 0:
		return
	case 1:
		msg = fmt.Sprint(keyvals[0])
	default:
		if length%2 == 0 {
			for i := 0; i < len(keyvals); i += 2 {
				data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
			}
		} else {
			for i := 1; i < len(keyvals); i += 2 {
				data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
			}
		}
	}
	// data = append(data, zap.StackSkip("stack", 4))
	l.logger.Log(level, msg, data...)
}

func (l loggerEnt) DebugLog(keyvals ...interface{}) {
	l.Log(zapcore.DebugLevel, keyvals...)
}

func (l loggerEnt) ErrorLog(keyvals ...interface{}) {
	l.Log(zapcore.ErrorLevel, keyvals...)
}

func newloggerEnt(l *zap.Logger) *loggerEnt {
	return &loggerEnt{
		logger: l.WithOptions(zap.AddCallerSkip(5)),
	}
}

type multiDriver struct {
	r, w dialect.Driver
}

var _ dialect.Driver = (*multiDriver)(nil)

func (d *multiDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return d.r.Query(ctx, query, args, v)
}

func (d *multiDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return d.w.Exec(ctx, query, args, v)
}

func (d *multiDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	return d.w.Tx(ctx)
}

func (d *multiDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	return d.w.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
}

func (d *multiDriver) Close() error {
	rerr := d.r.Close()
	werr := d.w.Close()
	if rerr != nil {
		return rerr
	}
	if werr != nil {
		return werr
	}
	return nil
}

func (d *multiDriver) Dialect() string {
	return d.w.Dialect()
}

func (d *Data) LoadEntDatabase(c *database.Config) (*ent.Client, func(), error) {
	driverName := c.GetDriver()
	if driverName == "" {
		driverName = "mysql"
	}
	driver, err := sql.Open(driverName, database.ParseDSN(c))
	if err != nil {
		return nil, nil, err
	}

	driver.DB()
	driver.DB().SetMaxOpenConns(100)
	driver.DB().SetMaxIdleConns(10)
	driver.DB().SetConnMaxLifetime(1 * time.Hour)
	driver.DB().SetConnMaxIdleTime(30 * time.Minute)

	if err := driver.DB().Ping(); err != nil {
		_ = driver.Close()
		return nil, nil, err
	}

	options := []entcache.Option{}

	switch strings.ToLower(c.GetCache()) {
	case "redis", "rds":
		if d.rds != nil {
			// options = append(options, entcache.Levels(entcache.NewRedis(d.rds)))
		}
		options = append(options, entcache.ContextLevel())
	case "lru":
		options = append(options, entcache.Levels(entcache.NewLRU(300)))
	case "context":
		fallthrough
	default:
		options = append(options, entcache.ContextLevel())
	}
	if len(options) > 0 {
		options = append(options, entcache.TTL(10*time.Minute))
		options = append(options, entcache.Hash(func(query string, args []any) (entcache.Key, error) {
			v, err := entcache.DefaultHash(query, args)
			if err != nil {
				return 0, err
			}
			return fmt.Sprintf("area.ent.%d", v.(uint64)), nil
		}))
	}
	logfunc := newloggerEnt(d.logger)
	cacheDriver := entcache.NewDriver(driver, options...)
	cacheDriver.Log = func(v ...any) {
		logfunc.ErrorLog(v...)
	}
	client := ent.NewClient(
		ent.Debug(),
		ent.Driver(cacheDriver),
		ent.Log(logfunc.DebugLog),
	)

	var cleanup = func() {
		if err := client.Close(); err != nil {
			d.logger.Error("[ent] close failed", zap.Error(err))
		}
	}

	return client, cleanup, nil
}
