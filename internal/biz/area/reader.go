package area

type Inquirer interface {
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

type Query func(Inquirer)

func Offset(offset int) Query {
	return func(r Inquirer) {
		r.Offset(offset)
	}
}

func Limit(limit int) Query {
	return func(r Inquirer) {
		r.Limit(limit)
	}
}

func Order(order string) Query {
	return func(r Inquirer) {
		r.Order(order)
	}
}

func IDEQ(id uint64) Query {
	return func(r Inquirer) {
		r.IDEQ(id)
	}
}

func IDIn(ids ...uint64) Query {
	return func(r Inquirer) {
		r.IDIn(ids...)
	}
}

func ParentIDEQ(parentID uint64) Query {
	return func(r Inquirer) {
		r.ParentIDEQ(parentID)
	}
}

func RegionIDEQ(regionID string) Query {
	return func(r Inquirer) {
		r.RegionIDEQ(regionID)
	}
}

func LevelEQ(level int) Query {
	return func(r Inquirer) {
		r.LevelEQ(level)
	}
}

func TitleContains(keyword string) Query {
	return func(r Inquirer) {
		r.TitleContains(keyword)
	}
}
