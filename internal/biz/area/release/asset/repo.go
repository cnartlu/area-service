package asset

import (
	"context"
)

type Asseter interface {
	Count(ctx context.Context, options ...Option) int
	// FindList 查找数据列表
	FindList(ctx context.Context, options ...Option) ([]*Asset, error)
	// FindOne 查找数据
	FindOne(ctx context.Context, options ...Option) (*Asset, error)
	// Save 新增或保存数据
	Save(ctx context.Context, data *Asset) (*Asset, error)
	// Remove 移除数据
	Remove(ctx context.Context, options ...Option) error
}
