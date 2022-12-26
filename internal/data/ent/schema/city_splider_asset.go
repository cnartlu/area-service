package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CitySpliderAsset holds the schema definition for the CitySpliderAsset entity.
type CitySpliderAsset struct {
	ent.Schema
}

// Fields of the CitySpliderAsset.
func (CitySpliderAsset) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("city_splider_id").Comment("城市蜘蛛").Default(0),
		field.Uint64("source_id").Comment("资源ID").Default(0),
		field.String("file_title").Comment("文件标题").Default(""),
		field.String("file_path").Comment("文件路径").Default(""),
		field.Uint("file_size").Comment("文件大小").Default(0),
		field.Uint8("status").Default(0).Comment("状态"),
		field.Uint64("create_time").Default(0).Comment("创建时间"),
		field.Uint64("update_time").Default(0).Comment("更新时间"),
	}
}

// Edges of the CitySpliderAsset.
func (CitySpliderAsset) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the schema.
func (CitySpliderAsset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "city_splider_asset",
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Options:   "ENGINE = INNODB",
		},
	}
}
