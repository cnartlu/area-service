package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CitySpliderAreaPolygon holds the schema definition for the CitySpliderAreaPolygon entity.
type CitySpliderAreaPolygon struct {
	ent.Schema
}

// Fields of the CitySpliderAreaPolygon.
func (CitySpliderAreaPolygon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.String("region_id").Default("").Comment("地区标识"),
		field.Float("lng").Default(0).SchemaType(map[string]string{dialect.MySQL: "DECIMAL(11,8)"}).Comment("经度"),
		field.Float("lat").Default(0).SchemaType(map[string]string{dialect.MySQL: "DECIMAL(11,8)"}).Comment("纬度"),
	}
}

// Edges of the CitySpliderAreaPolygon.
func (CitySpliderAreaPolygon) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the schema.
func (CitySpliderAreaPolygon) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("region_id"),
	}
}

// Annotations of the schema.
func (CitySpliderAreaPolygon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "city_splider_area_polygon",
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Options:   "ENGINE = INNODB",
		},
	}
}
