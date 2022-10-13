package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type HandlerFunc func(*redis.Client)

func WithPing(client *redis.Client) {
	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(fmt.Errorf("redis ping connection error: %w", err))
	}
}
