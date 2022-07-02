package cron

import (
	"context"

	"github.com/cnartlu/area-service/pkg/component/log"

	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/cnartlu/area-service/internal/cron/job"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	logger *log.Logger
	rdb    *redis.Client
	db     *db.DB
	server *cron.Cron

	githubJob *job.Github
}

// Start cron 服务启动
func (c *Cron) Start() (err error) {
	// TODO 编写 cron 任务
	// if _, err = c.server.AddFunc("*/5 * * * * *", func() {}); err != nil { // 每 5 秒钟运行一次
	// 	return err
	// }
	// if _, err = c.server.AddJob("@daily", c.exampleJob); err != nil { // 每天 00:00 运行一次
	// 	return err
	// }
	// if _, err = c.server.AddJob("@every 1h30m10s", c.exampleJob); err != nil { // 每 1 小时 30 分 10 秒运行一次
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
	db *db.DB,
	// .... 此处开始注入job
	githubJob *job.Github,
) (*Cron, error) {
	server := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(cron.PrintfLogger(logger)),
			cron.DelayIfStillRunning(cron.PrintfLogger(logger)),
		),
	)

	return &Cron{
		logger:    logger,
		rdb:       rdb,
		db:        db,
		server:    server,
		githubJob: githubJob,
	}, nil
}
