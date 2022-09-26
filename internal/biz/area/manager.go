package area

import (
	"context"
)

type FindListParam struct {
	// ParentID 父级ID
	ParentID uint64
	// RegionID 父级区域ID
	RegionID string
	// Level 区域级别
	Level int
	// Keyword 搜索关键字
	Keyword string
	// Order 排序
	Order string
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

// List 列表更新
func (m *ManagerUsecase) List(ctx context.Context, params FindListParam) ([]*Area, error) {
	options := []Option{}
	if params.RegionID != "" {
		options = append(options, WithRegionID(params.RegionID))
		if params.Level > 0 {
			options = append(options, WithLevel(params.Level))
		}
		parent, err := m.manager.FindOne(ctx, options...)
		if err != nil {
			return nil, err
		}
		params.ParentID = parent.ID
		options = []Option{}
	}
	options = append(options, WithParentID(params.ParentID))
	if params.Keyword != "" {
		options = append(options, WithKeywordContains(params.Keyword))
	}

	return m.manager.FindList(ctx, options...)
}

func (m *ManagerUsecase) CascadeList(ctx context.Context, options ...Option) ([]*CascadeArea, error) {
	// 查找数据
	return nil, nil
}

// ViewWithIDEQ 查询ID值等价
func (m *ManagerUsecase) ViewWithIDEQ(ctx context.Context, id uint64) (*Area, error) {
	return m.manager.FindOne(ctx, WithID(id))
}

func (m *ManagerUsecase) ViewWithRegionID(ctx context.Context, regionID string, level int) (*Area, error) {
	options := []Option{WithRegionID(regionID)}
	if level > 0 {
		options = append(options, WithLevel(level))
	}
	return m.manager.FindOne(ctx, options...)
}

// DeleteWithID 删除值
func (m *ManagerUsecase) DeleteWithID(ctx context.Context, id uint64) error {
	return m.manager.Remove(ctx, WithID(id))
}

func NewManagerUsecase(manager Manager) *ManagerUsecase {
	return &ManagerUsecase{
		manager: manager,
	}
}
