package db

import (
	"context"
	"strconv"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/component/ent/migrate"
	"github.com/go-kratos/kratos/v2/log"
)

func NewEnt(config *Config_DB, logger log.Logger) (*ent.Client, func(), error) {
	driver, err := sql.Open(config.Driver, buildSource(config))
	if err != nil {
		return nil, nil, err
	}

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

	hLogger := log.NewHelper(logger)
	cleanup := func() {
		hLogger.Info("closing the ent resources")

		err = driver.Close()
		if err != nil {
			hLogger.Error(err)
		}
	}

	client := ent.NewClient(
		ent.Driver(driver),
		ent.Log(func(i ...interface{}) {
			hLogger.Debug(i)
		}),
	)

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false),
	); err != nil {
		hLogger.Errorf("failed to creat schema resources: %v", err)
		return nil, nil, err
	}

	return client, cleanup, nil
}

func buildSource(c *Config_DB) string {
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
