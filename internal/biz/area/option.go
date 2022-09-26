package area

// pager 分页查询
type pager interface {
	Limit(int)
	Offset(int)
}

type OptionInterface interface {
	pager
	WithID(uint64)
	WithParentID(pid uint64)
	WithRegionID(regionID string)
	WithLevel(int)
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
		r.WithID(id)
	}
}

// WithParentIDEQ 父级ID等价
func WithParentID(pid uint64) Option {
	return func(r OptionInterface) {
		r.WithParentID(pid)
	}
}

// WithRegionID 区域标识
func WithRegionID(regionID string) Option {
	return func(r OptionInterface) {
		r.WithRegionID(regionID)
	}
}

// WithLevel 级别
func WithLevel(level int) Option {
	return func(r OptionInterface) {
		r.WithLevel(level)
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
