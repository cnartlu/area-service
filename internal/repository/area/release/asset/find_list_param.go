package asset

type FindListParam struct {
	AreaReleaseID uint64
	Keyword       string
	Status        string
	Pagination    bool
	Offset        int
	Limit         int
}

func (f *FindListParam) SetAreaReleaseID(areaReleaseID uint64) *FindListParam {
	f.AreaReleaseID = areaReleaseID
	return f
}

func (f *FindListParam) SetKeyword(keyword string) *FindListParam {
	f.Keyword = keyword
	return f
}

func (f *FindListParam) SetStatus(status string) *FindListParam {
	f.Status = status
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

func NewFindListParam() *FindListParam {
	return &FindListParam{}
}
