package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AreaPolygon holds the schema definition for the AreaPolygon entity.
type AreaPolygon struct {
	ent.Schema
}

// Fields of the AreaPolygon.
func (AreaPolygon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("area_id").Default(0).Comment("区域"),
		field.String("region_id").Default("").Comment("地区标识"),
		field.Float("lat").Default(0).SchemaType(map[string]string{dialect.MySQL: "DECIMAL(10,7)"}).Comment("纬度"),
		field.Float("lng").Default(0).SchemaType(map[string]string{dialect.MySQL: "DECIMAL(10,7)"}).Comment("纬度"),
		field.Uint64("delete_time").Default(0).Comment("删除时间"),
		field.Uint64("create_time").Default(0).Comment("创建时间"),
		field.Uint64("update_time").Default(0).Comment("更新时间"),
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
