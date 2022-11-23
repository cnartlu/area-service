package release

type OptionInterface interface {
	Limit(limit int)
	Offset(offset int)
	Order(order string)
	IDEQ(id uint64)
	IDIn(ids ...uint64)
	ReleaseIDEQ(id uint64)
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

func Order(order string) Option {
	return func(r OptionInterface) {
		r.Order(order)
	}
}

func IDEQ(id uint64) Option {
	return func(r OptionInterface) {
		r.IDEQ(id)
	}
}

func IDIn(ids ...uint64) Option {
	return func(r OptionInterface) {
		r.IDIn(ids...)
	}
}

func ReleaseIDEQ(releaseID uint64) Option {
	return func(r OptionInterface) {
		r.ReleaseIDEQ(releaseID)
	}
}
