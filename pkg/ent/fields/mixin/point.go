package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Lng 纬度
type Lng struct {
	mixin.Schema
}

func (Lng) Fields() []ent.Field {
	return []ent.Field{
		field.Float("lng").Comment("经度").Default(0).SchemaType(map[string]string{
			dialect.MySQL: "DECIMAL(10,7)",
		}),
	}
}

// Lat 经度
type Lat struct {
	mixin.Schema
}

func (Lat) Fields() []ent.Field {
	return []ent.Field{
		field.Float("lat").Comment("纬度").Default(0).SchemaType(map[string]string{
			dialect.MySQL: "DECIMAL(10,7)",
		}),
	}
}

// Point 点位坐标
type Point struct {
	mixin.Schema
}

func (Point) Fields() []ent.Field {
	return append(
		Lng{}.Fields(),
		Lat{}.Fields()...,
	)
}
