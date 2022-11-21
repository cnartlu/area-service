package area

type OptionInterface interface {
	Limit(limit int)
	Offset(offset int)
	Order(order string)
	IDEQ(id uint64)
	IDIn(ids ...uint64)
	ParentIDEQ(parentID uint64)
	RegionIDEQ(regionID string)
	LevelEQ(level int)
	TitleContains(keyword string)
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

func ParentIDEQ(parentID uint64) Option {
	return func(r OptionInterface) {
		r.ParentIDEQ(parentID)
	}
}

func RegionIDEQ(regionID string) Option {
	return func(r OptionInterface) {
		r.RegionIDEQ(regionID)
	}
}

func LevelEQ(level int) Option {
	return func(r OptionInterface) {
		r.LevelEQ(level)
	}
}

func TitleContains(keyword string) Option {
	return func(r OptionInterface) {
		r.TitleContains(keyword)
	}
}
