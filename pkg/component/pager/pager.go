package pager

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
