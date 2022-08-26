package release

import (
	"context"
	"time"
)

type ManagerData struct {
	ID, ReleaseID uint64
	Owner, Repo   string
	Name, NodeID  string
	Content       string
	PublishedAt   time.Time
	Status        int
}

type Manager interface {
	Count(ctx context.Context, options ...Option) int
	// FindList 查找数据列表
	FindList(ctx context.Context, options ...Option) ([]*ManagerData, error)
	// FindOne 查找数据
	FindOne(ctx context.Context, options ...Option) (*ManagerData, error)
	// Save 新增或保存数据
	// Save(ctx context.Context, data *ManagerData) (*ManagerData, error)
	// // Remove 移除数据
	// Remove(ctx context.Context, options ...Option) error
}

type Management struct {
	manager Manager
}

func (m *Management) List(ctx context.Context) {
	m.manager.FindList(ctx)
}

func NewManaement(manager Manager) *Management {
	return &Management{
		manager: manager,
	}
}
