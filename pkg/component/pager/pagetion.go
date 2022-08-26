package pager

import "math"

const (
	// defaultPageSize 未定义页面大小的时候的默认大小
	defaultPageSize int = 20
)

type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
	FirstPage  int `json:"firstPage"`
	LastPage   int `json:"lastPage"`
	PrevPage   int `json:"prevPage"`
	NextPage   int `json:"nextPage"`
}

// GetPage 获取页码
func (p Pagination) GetPage() int {
	if p.Page < 1 {
		return 1
	}
	return p.Page
}

// GetPageSize 获取页码大小
func (p Pagination) GetPageSize() int {
	if p.PageSize < 1 {
		return defaultPageSize
	}
	return p.PageSize
}

// GetTotalCount 获取数据量总数
func (p Pagination) GetTotalCount() int {
	return p.TotalCount
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

func (p *Pagination) SetTotalCount(totalCount int) *Pagination {
	p.TotalCount = totalCount
	return p
}

func (p *Pagination) SetPage(page int) *Pagination {
	p.Page = page
	return p
}

func (p *Pagination) SetPageSize(pageSize int) *Pagination {
	p.PageSize = pageSize
	return p
}

func NewPagination(page int, pageSize int, totalCount int) *Pagination {
	return &Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
	}
}
