package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CitySpliderArea holds the schema definition for the CitySpliderArea entity.
type CitySpliderArea struct {
	ent.Schema
}

// Fields of the CitySpliderArea.
func (CitySpliderArea) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("parent_id").Comment("父级").Default(0),
		field.String("parent_region_id").Comment("父级区域").Default(""),
		field.String("region_id").Comment("区域标识").Default(""),
		field.String("title").Comment("标题").Default(""),
		field.Uint8("level").Comment("等级").Default(1).SchemaType(map[string]string{dialect.MySQL: "TINYINT(1)"}),
		field.Float("lng").Default(0).SchemaType(map[string]string{dialect.MySQL: "DECIMAL(11,8)"}).Comment("经度"),
		field.Float("lat").Default(0).SchemaType(map[string]string{dialect.MySQL: "DECIMAL(11,8)"}).Comment("纬度"),
	}
}

// Edges of the CitySpliderArea.
func (CitySpliderArea) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the schema.
func (CitySpliderArea) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "city_splider_area",
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Options:   "ENGINE = INNODB",
		},
	}
}
