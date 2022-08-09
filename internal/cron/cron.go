package cron

import (
	"context"

	"github.com/cnartlu/area-service/pkg/component/log"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/cron/jobs"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	logger *log.Logger
	rdb    *redis.Client
	db     *ent.Client
	server *cron.Cron

	syncArea *jobs.SyncArea
}

// Start cron 服务启动
func (c *Cron) Start() (err error) {
	// TODO 编写 cron 任务
	// 每 5 秒钟运行一次
	// if _, err = c.server.AddJob("*/5 * * * * *", c.syncArea); err != nil {
	// 	return err
	// }
	// 每天 00:00 运行一次
	if _, err = c.server.AddJob("@daily", c.syncArea); err != nil {
		return err
	}
	// 每 1 小时 30 分 10 秒运行一次
	// if _, err = c.server.AddJob("@every 12h00m00s", c.exampleJob); err != nil {
	// 	return err
	// }

	// 启动 cron 服务
	c.server.Start()

	c.logger.Info("cron server started")
	return nil
}

// Stop cron 服务关闭
func (c *Cron) Stop(ctx context.Context) (err error) {
	c.server.Stop()

	c.logger.Info("cron server has been stop")
	return nil
}

// New 实例化cron对象
func New(
	logger *log.Logger,
	rdb *redis.Client,
	db *ent.Client,
	// .... 此处开始注入job
	sa *jobs.SyncArea,
) (*Cron, error) {
	server := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(cron.PrintfLogger(logger)),
			cron.DelayIfStillRunning(cron.PrintfLogger(logger)),
		),
	)

	return &Cron{
		logger:   logger,
		rdb:      rdb,
		db:       db,
		server:   server,
		syncArea: sa,
	}, nil
}
