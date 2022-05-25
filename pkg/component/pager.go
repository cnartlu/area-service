package component

import "math"

var (
	defaultPageSize int = 20
)

// PagerFetcherInput 分页获取器
type PagerFetcherInput interface {
	GetPage() int
	GetPageSize() int
	GetTotalCount() int
}

// PagerFetcher 分页获取器
type PagerFetcher interface {
	GetLimit() int
	GetOffset() int
	GetTotalPage() int
}

// PagerSetter 分页设置器
type PagerSetter interface {
	SetTotalCount(totalCount int)
	SetPage(page int)
	SetPageSize(pageSize int)
}

// 分页
type Pager interface {
	PagerFetcher
	PagerSetter
}

type Pagination struct {
	page       int
	pageSize   int
	totalCount int
}

// GetPage 获取页码
func (p Pagination) GetPage() int {
	if p.page < 1 {
		return 1
	}
	return p.page
}

// GetPageSize 获取页码大小
func (p Pagination) GetPageSize() int {
	if p.pageSize < 1 {
		return defaultPageSize
	}
	return p.pageSize
}

// GetTotalCount 获取数据量总数
func (p Pagination) GetTotalCount() int {
	return p.totalCount
}

// GetTotalPage 获取最大分页的页数
func (p Pagination) GetTotalPage() int {
	if p.GetTotalCount() < 1 {
		return 0
	}
	return int(math.Ceil(float64(p.GetTotalCount()) / float64(p.GetPageSize())))
}

// GetOffset 获取数据的偏移量
func (p Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

// GetLimit 获取限制长度
func (p Pagination) GetLimit() int {
	return p.GetPageSize()
}

func NewPagination(page int, pageSize int, totalCount int) *Pagination {
	return &Pagination{
		page:       page,
		pageSize:   pageSize,
		totalCount: totalCount,
	}
}
