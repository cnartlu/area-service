package release

// Inquirer 查询人
type Inquirer interface {
	Limit(limit int)
	Offset(offset int)
	Order(order string)
	IDEQ(id uint64)
	IDIn(ids ...uint64)
	ReleaseIDEQ(id uint64)
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

func ReleaseIDEQ(releaseID uint64) Query {
	return func(r Inquirer) {
		r.ReleaseIDEQ(releaseID)
	}
}
