package component

type Pager interface{}

type Pagination struct {
	page      int
	pageSize  int
	totalSize int
}

func NewPagination(page int, pageSize int, totalSize int) *Pagination {
	return &Pagination{
		page:      page,
		pageSize:  pageSize,
		totalSize: totalSize,
	}
}
