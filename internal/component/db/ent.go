package db

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/pkg/component/log"
	"go.uber.org/zap"
)

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
		return nil, func() {}, nil
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
		logger.Error("", zap.Error(err))
		return nil, nil, err
	}
	// 默认配置
	driver.DB().SetMaxOpenConns(100)
	driver.DB().SetMaxIdleConns(10)
	// 连接的最大生命周期
	driver.DB().SetConnMaxLifetime(0)
	driver.DB().SetConnMaxIdleTime(time.Second * 60 * 60)

	// if config.MaxIdleConn > 0 {
	// 	driver.DB().SetMaxIdleConns(config.MaxIdleConn)
	// }
	// if config.MaxOpenConn > 0 {
	// 	driver.DB().SetMaxOpenConns(config.MaxOpenConn)
	// }
	// if config.ConnMaxIdleTime > 0 {
	// 	driver.DB().SetConnMaxIdleTime(time.Duration(config.ConnMaxIdleTime) * time.Second)
	// }
	// if config.ConnMaxLifeTime > 0 {
	// 	driver.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)
	// }

	cleanup := func() {
		logger.Info("closing the ent resources")

		err = driver.Close()
		if err != nil {
			logger.Error("close db resources failed", zap.Error(err))
		}
	}

	if err := driver.DB().Ping(); err != nil {
		defer cleanup()
		logger.Error("", zap.Error(err))
		return nil, nil, err
	}

	client := ent.NewClient(
		ent.Driver(driver),
		ent.Log(func(i ...interface{}) {
			logger.DebugLog(i)
		}),
	)

	// err = client.Schema.Create(
	// 	context.Background(),
	// 	migrate.WithForeignKeys(false),
	// )
	// if err != nil {
	// 	hLogger.Errorf("failed to creat schema resources: %v", err)
	// 	return nil, nil, err
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
