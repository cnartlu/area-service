package area

type Inquirer interface {
	Limit(limit int)
	Offset(offset int)
	Order(orders ...string)
	IDEQ(id int)
	IDIn(ids ...int)
	ParentIDEQ(parentID int)
	RegionIDEQ(regionID string)
	LevelEQ(level int)
	TitleContains(keyword string)
}

type Query func(Inquirer)

func Cache(ttl int) Query {
	return func(r Inquirer) {
		if o, ok := r.(interface{ Cache(ttl int) }); ok {
			o.Cache(ttl)
		}
	}
}

func Offset(offset int) Query {
	return func(r Inquirer) {
		if o, ok := r.(interface{ Offset(offset int) }); ok {
			o.Offset(offset)
		}
	}
}

func Limit(limit int) Query {
	return func(r Inquirer) {
		if o, ok := r.(interface{ Limit(limit int) }); ok {
			o.Limit(limit)
		}
	}
}

func Order(orders ...string) Query {
	return func(r Inquirer) {
		r.Order(orders...)
	}
}

func IDEQ(id int) Query {
	return func(r Inquirer) {
		r.IDEQ(id)
	}
}

func IDIn(ids ...int) Query {
	return func(r Inquirer) {
		r.IDIn(ids...)
	}
}

func ParentIDEQ(parentID int) Query {
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
