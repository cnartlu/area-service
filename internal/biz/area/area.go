package area

import (
	"context"
	"sync"
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

type AreaRepo interface {
	Count(ctx context.Context, options ...Query) int
	// FindList 查找数据列表
	FindList(ctx context.Context, options ...Query) ([]*Area, error)
	// FindOne 查找数据
	FindOne(ctx context.Context, options ...Query) (*Area, error)
	// Save 新增或保存数据
	Save(ctx context.Context, data *Area) (*Area, error)
	// Remove 移除数据
	Remove(ctx context.Context, options ...Query) error
}

type AreaUsecase struct {
	repo AreaRepo
}

// List 列表更新
func (m *AreaUsecase) List(ctx context.Context, params FindListParam) ([]*Area, error) {
	options := []Query{}
	if params.RegionID != "" {
		options = append(options, RegionIDEQ(params.RegionID))
		if params.Level > 0 {
			options = append(options, LevelEQ(params.Level))
		}
		parent, err := m.repo.FindOne(ctx, options...)
		if err != nil {
			return nil, err
		}
		params.ParentID = parent.ID
		options = []Query{}
	}
	options = append(options, ParentIDEQ(params.ParentID))
	if params.Keyword != "" {
		options = append(options, TitleContains(params.Keyword))
	}
	if params.Order == "" {
		params.Order = "-id"
	}
	options = append(options, Order(params.Order))
	return m.repo.FindList(ctx, options...)
}

// CascadeList 级联列表
// parentID 父级ID，为0时表示顶级
// maxDeep 最大获取深度，小于等于0标识不限制，一直往下取
func (m *AreaUsecase) CascadeList(ctx context.Context, parentID uint64, maxDeep int) ([]*CascadeArea, error) {
	var handlerFunc func(parentID uint64, deep int) ([]*CascadeArea, error)
	var cancelCtx, cancelFun = context.WithCancel(ctx)
	var cerr error
	var nolimitDeep bool = maxDeep <= 0
	handlerFunc = func(parentID uint64, deep int) ([]*CascadeArea, error) {
		var g = &sync.WaitGroup{}
		results, err := m.repo.FindList(cancelCtx, ParentIDEQ(parentID))
		if err != nil {
			cancelFun()
			return nil, err
		}
		items := make([]*CascadeArea, len(results))
		for k, result := range results {
			result := result
			item := &CascadeArea{
				ID:             result.ID,
				RegionID:       result.RegionID,
				Title:          result.Title,
				Ucfirst:        result.Ucfirst,
				Pinyin:         result.Pinyin,
				Level:          result.Level,
				ChildrenNumber: 0,
				Items:          make([]*CascadeArea, 0),
			}
			items[k] = item
			// 深度达到
			if nolimitDeep || deep < maxDeep {
				g.Add(1)
				go func(r *CascadeArea) {
					defer g.Done()
					items, _ := handlerFunc(result.ID, deep+1)
					if err != nil {
						cerr = err
						cancelFun()
						return
					}
					r.Items = items
					r.ChildrenNumber = len(items)
				}(item)
			}
		}
		g.Wait()
		if cerr != nil {
			return nil, cerr
		}
		return items, nil
	}
	results, err := handlerFunc(parentID, 1)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// FindOne 查询ID值等价
func (m *AreaUsecase) FindOne(ctx context.Context, id uint64) (*Area, error) {
	return m.repo.FindOne(ctx, IDEQ(id))
}

func (m *AreaUsecase) FindByRegionID(ctx context.Context, regionID string, level int) (*Area, error) {
	options := []Query{RegionIDEQ(regionID)}
	if level > 0 {
		options = append(options, LevelEQ(level))
	}
	options = append(options, Order("region_id"))
	return m.repo.FindOne(ctx, options...)
}

// Delete 删除值
func (m *AreaUsecase) Delete(ctx context.Context, ids ...uint64) error {
	return m.repo.Remove(ctx, IDIn(ids...))
}

func NewAreaUsecase(repo AreaRepo) *AreaUsecase {
	return &AreaUsecase{
		repo: repo,
	}
}
