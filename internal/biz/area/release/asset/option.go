package asset

// pager 分页查询
type pager interface {
	Limit(int)
	Offset(int)
}

// order 排序
type order interface {
}

// 查询条件
type querier interface {
	IDEQ(uint64)
	// WithReiginIDAndLevel(string, uint8)
}

type OptionInterface interface {
	pager
	order
	querier
}

type Option func(OptionInterface)

func Offset(offset int) Option {
	return func(r OptionInterface) {
		r.Offset(offset)
	}
}

func Limit(limit int) Option {
	return func(r OptionInterface) {
		r.Limit(limit)
	}
}

func WithID(id uint64) Option {
	return func(r OptionInterface) {
		i, ok := r.(interface{ WithID(uint64) })
		if ok {
			i.WithID(id)
		}
	}
}

func WithAreaReleaseID(areaReleaseID uint64) Option {
	return func(r OptionInterface) {
		i, ok := r.(interface{ WithAreaReleaseID(uint64) })
		if ok {
			i.WithAreaReleaseID(areaReleaseID)
		}
	}
}
