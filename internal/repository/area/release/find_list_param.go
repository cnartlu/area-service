package release

type FindListParam struct {
	Owner      string
	Repository string
	Keyword    string
	Pagination bool
	Offset     int
	Limit      int
}

func (f *FindListParam) SetKeyword(keyword string) *FindListParam {
	f.Keyword = keyword
	return f
}

func (f *FindListParam) SetPagination() *FindListParam {
	f.Pagination = true
	return f
}

func (f *FindListParam) SetOffset(offset int) *FindListParam {
	f.Offset = offset
	return f
}

func (f *FindListParam) SetLimit(limit int) *FindListParam {
	f.Limit = limit
	return f
}
