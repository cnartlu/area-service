package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// New 创建 redis 客户端
// 当 Close 客户端时
func New(config *Config) (*redis.Client, func(), error) {
	if config == nil {
		return nil, nil, fmt.Errorf("component redis new error config is nil")
	}
	if config.Port < 1 {
		config.Port = 6379
	}
	option := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
	}
	if config.Username != "" {
		option.Username = config.Username
	}
	if config.Password != "" {
		option.Password = config.Password
	}
	if config.Db != 0 {
		option.DB = int(config.Db)
	}
	if config.MaxRetries != 0 {
		option.MaxRetries = int(config.MaxRetries)
	}
	if config.MinRetryBackoff != 0 {
		option.MinRetryBackoff = time.Duration(config.MinRetryBackoff) * time.Second
	}
	if config.MaxRetryBackoff != 0 {
		option.MaxRetryBackoff = time.Duration(config.MaxRetryBackoff) * time.Second
	}
	if config.GetDialTimeout() != nil && config.GetDialTimeout().IsValid() {
		option.DialTimeout = config.GetDialTimeout().AsDuration() * time.Second
	}
	if config.GetReadTimeout() != nil && config.GetReadTimeout().IsValid() {
		option.ReadTimeout = config.GetReadTimeout().AsDuration() * time.Second
	}
	if config.GetWriteTimeout() != nil && config.GetWriteTimeout().IsValid() {
		option.WriteTimeout = config.GetWriteTimeout().AsDuration() * time.Second
	}
	if config.PoolSize != 0 {
		option.PoolSize = int(config.PoolSize)
	}
	if config.MinIdleConns != 0 {
		option.MinIdleConns = int(config.MinIdleConns)
	}
	if config.MaxConnAge != 0 {
		option.MaxConnAge = time.Duration(config.MaxConnAge) * time.Second
	}
	if config.PoolTimeout != nil && config.PoolTimeout.IsValid() {
		option.PoolTimeout = config.PoolTimeout.AsDuration() * time.Second
	}
	if config.IdleTimeout != nil && config.IdleTimeout.IsValid() {
		option.IdleTimeout = config.IdleTimeout.AsDuration() * time.Second
	}
	if config.IdleCheckFrequency != 0 {
		option.IdleCheckFrequency = time.Duration(config.IdleCheckFrequency) * time.Second
	}

	client := redis.NewClient(option)

	ctx := context.Background()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, nil, fmt.Errorf("redis ping connection error: %w", err)
	}

	cleanup := func() {
		if err := client.Close(); err != nil {
			return
		}
		return
	}

	return client, cleanup, nil
}
