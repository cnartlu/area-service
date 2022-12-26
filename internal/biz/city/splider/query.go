package splider

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
	SourceEQ(source string)
	OwnerEQ(owner string)
	RepoEQ(repo string)
	SourceIDEQ(sourceID uint64)
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

func SourceEQ(source string) Query {
	return func(r Inquirer) {
		r.SourceEQ(source)
	}
}

func OwnerEQ(owner string) Query {
	return func(r Inquirer) {
		r.OwnerEQ(owner)
	}
}

func RepoEQ(repo string) Query {
	return func(r Inquirer) {
		r.RepoEQ(repo)
	}
}

func SourceIDEQ(sourceID uint64) Query {
	return func(r Inquirer) {
		r.SourceIDEQ(sourceID)
	}
}
