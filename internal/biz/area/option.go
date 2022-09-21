package area

// pager 分页查询
type pager interface {
	Limit(int)
	Offset(int)
}

// order 排序
type order interface {
}

type OptionInterface interface {
	pager
	order
	WithID(uint64)
}

// Option 请求条件接口
type Option func(OptionInterface)

// Offset 偏移量
func Offset(offset int) Option {
	return func(r OptionInterface) {
		r.Offset(offset)
	}
}

// Limit 限制查询数量
func Limit(limit int) Option {
	return func(r OptionInterface) {
		r.Limit(limit)
	}
}

// WithID 查询具体ID
func WithID(id uint64) Option {
	return func(r OptionInterface) {
		i, ok := r.(interface{ WithID(uint64) })
		if ok {
			i.WithID(id)
		}
	}
}

// WithParentIDEQ 父级ID等价
func WithParentID(id uint64) Option {
	return func(r OptionInterface) {
		i, ok := r.(interface{ WithParentID(uint64) })
		if ok {
			i.WithParentID(id)
		}
	}
}

// WithReiginIDAndLevel 查询区域ID和级别
func WithReiginIDAndLevel(regionID string, level uint8) Option {
	return func(r OptionInterface) {
		i, ok := r.(interface{ WithReiginIDAndLevel(string, uint8) })
		if ok {
			i.WithReiginIDAndLevel(regionID, level)
		}
	}
}

// WithKeywordContains 搜索关键字
func WithKeywordContains(keyword string) Option {
	return func(r OptionInterface) {
		i, ok := r.(interface{ WithKeywordContains(string) })
		if ok {
			i.WithKeywordContains(keyword)
		}
	}
}
