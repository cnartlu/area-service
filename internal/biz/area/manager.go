package area

import (
	"context"
)

type FindListParam struct {
	ParentID uint64
	Keyword  string
}

type Manager interface {
	Count(ctx context.Context, options ...Option) int
	// FindList 查找数据列表
	FindList(ctx context.Context, options ...Option) ([]*Area, error)
	// FindOne 查找数据
	FindOne(ctx context.Context, options ...Option) (*Area, error)
	// Save 新增或保存数据
	Save(ctx context.Context, data *Area) (*Area, error)
	// Remove 移除数据
	Remove(ctx context.Context, options ...Option) error
}

type ManagerUsecase struct {
	manager Manager
}

func (m *ManagerUsecase) List(ctx context.Context, params FindListParam) ([]*Area, error) {
	return m.manager.FindList(ctx)
}

func (m *ManagerUsecase) CascadeList(ctx context.Context, options ...Option) ([]*Area, error) {
	return nil, nil
}

func (m *ManagerUsecase) Delete(ctx context.Context, id uint64) error {
	return nil
}

func NewManagerUsecase(manager Manager) *ManagerUsecase {
	return &ManagerUsecase{
		manager: manager,
	}
}
