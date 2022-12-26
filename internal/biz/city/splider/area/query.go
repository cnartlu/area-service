package area

type OrderType int

const (
	OrderTypeCustome OrderType = iota
	OrderTypeAsc
	OrderTypeDesc
)

type Inquirer interface {
	Limit(limit int)
	Offset(offset int)
	Order(order ...string)
	IDEQ(id int)
	IDIn(ids ...int)
	RegionIDEQ(regionID string)
	LevelEQ(level int)
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
