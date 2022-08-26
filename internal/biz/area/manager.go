package area

import (
	"context"
)

type ManagerData struct {
	ID       uint64
	Title    string
	Pinyin   string
	Ucfirst  string
	CityCode string
	ZipCode  string
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
