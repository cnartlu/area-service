package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cnartlu/area-service/pkg/ent/fields/mixin"
)

// AreaPolygon holds the schema definition for the AreaPolygon entity.
type AreaPolygon struct {
	ent.Schema
}

// Mixin of the schema.
func (AreaPolygon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.DeleteTime{},
		mixin.Time{},
		mixin.Point{},
	}
}

// Fields of the AreaPolygon.
func (AreaPolygon) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("ID"),
		field.Uint64("area_id").Comment("区域").Default(0),
		field.String("region_id").Comment("地区标识").Default("").MaxLen(20),
	}
}

// Edges of the AreaPolygon.
func (AreaPolygon) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the schema.
func (AreaPolygon) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("area_id"),
		index.Fields("region_id"),
	}
}

// Annotations of the schema.
func (AreaPolygon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "area_polygon",
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Options:   "ENGINE = INNODB",
		},
	}
}
