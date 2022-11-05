package redis

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cnartlu/area-service/component/log"
	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
)

// New 创建 redis 客户端
// 当 Close 客户端时
func New(config *Config, logger *log.Logger) (*redis.Client, func(), error) {
	addr := bytes.Buffer{}
	if config.GetHost() != "" {
		addr.WriteString(config.GetHost())
	}
	if config.GetPort() < 0 {
		addr.WriteString(":")
		addr.Write([]byte(strconv.FormatInt(config.GetPort(), 10)))
	}
	option := &redis.Options{
		Addr: addr.String(),
	}
	if config.GetUsername() != "" {
		option.Username = config.GetUsername()
	}
	if config.GetPassword() != "" {
		option.Password = config.GetPassword()
	}
	if config.GetDb() != 0 {
		option.DB = int(config.GetDb())
	}
	if config.GetMaxRetries() != 0 {
		option.MaxRetries = int(config.GetMaxRetries())
	}
	if config.GetMinRetryBackoff() != 0 {
		option.MinRetryBackoff = time.Duration(config.GetMinRetryBackoff()) * time.Second
	}
	if config.GetMaxRetryBackoff() != 0 {
		option.MaxRetryBackoff = time.Duration(config.GetMaxRetryBackoff()) * time.Second
	}
	if config.GetDialTimeout().IsValid() {
		option.DialTimeout = config.GetDialTimeout().AsDuration() * time.Second
	}
	if config.GetReadTimeout().IsValid() {
		option.ReadTimeout = config.GetReadTimeout().AsDuration() * time.Second
	}
	if config.GetWriteTimeout().IsValid() {
		option.WriteTimeout = config.GetWriteTimeout().AsDuration() * time.Second
	}
	if config.GetPoolSize() != 0 {
		option.PoolSize = int(config.GetPoolSize())
	}
	if config.GetMinIdleConns() != 0 {
		option.MinIdleConns = int(config.GetMinIdleConns())
	}
	if config.GetMaxConnAge() != 0 {
		option.MaxConnAge = time.Duration(config.GetMaxConnAge()) * time.Second
	}
	if config.GetPoolTimeout().IsValid() {
		option.PoolTimeout = config.GetPoolTimeout().AsDuration() * time.Second
	}
	if config.GetIdleTimeout().IsValid() {
		option.IdleTimeout = config.GetIdleTimeout().AsDuration() * time.Second
	}
	if config.GetIdleCheckFrequency() != 0 {
		option.IdleCheckFrequency = time.Duration(config.GetIdleCheckFrequency()) * time.Second
	}

	client := redis.NewClient(option)

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, func() {}, fmt.Errorf("redis ping connection error: %w", err)
	}

	cleanup := func() {
		if err := client.Close(); err != nil {
			// 记录关闭的错误日志
			logger.Error("", zap.Error(err))
		}
	}

	return client, cleanup, nil
}
