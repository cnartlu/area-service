package cron

import (
	"context"
	"fmt"
	"time"

	job "github.com/cnartlu/area-service/internal/server/cron/job"
	"github.com/cnartlu/area-service/pkg/log"
	v3cron "github.com/robfig/cron/v3"
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

type Server struct {
	c      *v3cron.Cron
	logger *log.Logger
	// 任务的IDS
	entryIDs map[string][]v3cron.EntryID
	// 计划
	daily *job.Daily
}

func (c *Server) AppendEntryID(name string, id v3cron.EntryID) *Server {
	if id != 0 {
		c.entryIDs[name] = append(c.entryIDs[name], id)
	}
	return c
}

func (c *Server) OnceEntryID(name string, id v3cron.EntryID) *Server {
	if id != 0 {
		if len(c.entryIDs[name]) > 0 {
			for _, id := range c.entryIDs[name] {
				c.c.Remove(id)
			}
			c.entryIDs[name] = []v3cron.EntryID{}
		}
		c.entryIDs[name] = append(c.entryIDs[name], id)
	}
	return c
}

func (c *Server) Start(ctx context.Context) error {
	var (
		id  v3cron.EntryID
		err error
	)
	// 增加任务
	if id, err = c.c.AddJob("@daily", c.daily); err != nil {
		c.logger.Warn("add cron job failed", zap.Error(err))
	}
	c.OnceEntryID("daily", id)
	// 增加任务
	if id, err = c.c.AddJob("*/1 * * * * *", c.daily); err != nil {
		c.logger.Warn("add cron job failed", zap.Error(err))
	}
	c.OnceEntryID("daily1", id)

	c.c.Start()
	return nil
}

func (c *Server) Stop(ctx context.Context) error {
	c.logger.Info("[CRON] server stopping")
	c1 := c.c.Stop()
	number := 0
	for {
		select {
		case <-c1.Done():
			return c1.Err()
		default:
			number++
			switch number {
			case 1:
				time.Sleep(time.Millisecond * 100)
			case 2:
				time.Sleep(time.Millisecond * 500)
			case 3:
				time.Sleep(time.Second * 1)
			case 4:
				time.Sleep(time.Second * 3)
			// case 5:
			// 	time.Sleep(time.Second * 5)
			default:
				return fmt.Errorf("cron stop timeout")
			}
		}
	}
}

func NewServer(
	logger *log.Logger,
	daily *job.Daily,
) *Server {
	l := &cronLogger{l: logger.AddCallerSkip(1)}
	server := v3cron.New(v3cron.WithSeconds(), v3cron.WithLogger(l))
	return &Server{
		c:        server,
		logger:   logger,
		entryIDs: make(map[string][]v3cron.EntryID),
		daily:    daily,
	}
}
