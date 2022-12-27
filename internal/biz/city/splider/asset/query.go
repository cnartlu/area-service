package asset

type Inquirer interface {
	Limit(limit int)
	Offset(offset int)
	Order(order ...string)
	IDEQ(id int)
	IDIn(ids ...int)
	CitySpliderIDEQ(id int)
	SourceIDEQ(sourceID uint64)
	SourceIDIn(sourceIDs []uint64)
	StatusEQ(status Status)
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

func Order(order ...string) Query {
	return func(r Inquirer) {
		r.Order(order...)
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

func CitySpliderIDEQ(id int) Query {
	return func(r Inquirer) {
		r.CitySpliderIDEQ(id)
	}
}

func SourceIDEQ(sourceID uint64) Query {
	return func(r Inquirer) {
		r.SourceIDEQ(sourceID)
	}
}

func SourceIDIn(sourceIDs []uint64) Query {
	return func(r Inquirer) {
		r.SourceIDIn(sourceIDs)
	}
}

func StatusEQ(status Status) Query {
	return func(r Inquirer) {
		r.StatusEQ(status)
	}
}
