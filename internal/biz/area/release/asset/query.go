package asset

// Inquirer 查询记录
type Inquirer interface {
	Limit(limit int)
	Offset(offset int)
	Order(order string)
	IDEQ(id uint64)
	IDIn(ids ...uint64)
	AreaReleaseIDEQ(id uint64)
	StatusEQ(status Status)
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

func AreaReleaseIDEQ(id uint64) Query {
	return func(r Inquirer) {
		r.AreaReleaseIDEQ(id)
	}
}

func StatusEQ(status Status) Query {
	return func(r Inquirer) {
		r.StatusEQ(status)
	}
}
