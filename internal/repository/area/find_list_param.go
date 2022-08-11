package area

type FindListParam struct {
	Keyword  string
	ParentID uint64
}

func (f *FindListParam) SetKeyword(keyword string) *FindListParam {
	f.Keyword = keyword
	return f
}

func (f *FindListParam) SetParentID(parentID uint64) *FindListParam {
	f.ParentID = parentID
	return f
}
