package release

import (
	"context"
)

type Releasor interface {
	Count(ctx context.Context, options ...Option) int
	// FindList 查找数据列表
	FindList(ctx context.Context, options ...Option) ([]*Release, error)
	// FindOne 查找数据
	FindOne(ctx context.Context, options ...Option) (*Release, error)
	// Save 新增或保存数据
	Save(ctx context.Context, data *Release) (*Release, error)
	// Remove 移除数据
	Remove(ctx context.Context, options ...Option) error
}
