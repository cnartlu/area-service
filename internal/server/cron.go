package server

import (
	"context"
	"fmt"

	"github.com/cnartlu/area-service/pkg/log"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type cronLogger struct {
	l *log.Logger
}

func (c *cronLogger) format(keyvals ...interface{}) []zap.Field {
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}
	return data
}

func (c *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	c.l.Info(msg, c.format(keysAndValues...)...)
}

func (c *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	p := []zap.Field{zap.Error(err)}
	p = append(p, c.format(keysAndValues...)...)
	c.l.Error(msg, p...)
}

type Cron struct {
	c *cron.Cron
}

func (c *Cron) Start(ctx context.Context) error {
	c.c.Start()
	return nil
}

func (c *Cron) Stop(ctx context.Context) error {
	c1 := c.c.Stop()
	return c1.Err()
}

func NewCronServer(logger *log.Logger) *Cron {
	l := &cronLogger{l: logger}
	server := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(l),
			cron.DelayIfStillRunning(l),
			cron.SkipIfStillRunning(l),
		),
	)
	return &Cron{
		c: server,
	}
}
