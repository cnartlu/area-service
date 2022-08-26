package area

import (
	"context"
)

type FrontData struct {
	ID       uint64
	Title    string
	Pinyin   string
	Ucfirst  string
	CityCode string
	ZipCode  string
}

type Fronter interface {
	Count(ctx context.Context, options ...Option) int
	// FindList 查找数据列表
	FindList(ctx context.Context, options ...Option) ([]*FrontData, error)
	// FindOne 查找数据
	FindOne(ctx context.Context, options ...Option) (*FrontData, error)
	// Save 新增或保存数据
	Save(ctx context.Context, data *FrontData) (*FrontData, error)
	// Remove 移除数据
	Remove(ctx context.Context, options ...Option) error
}
