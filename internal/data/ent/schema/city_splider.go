package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CitySplider holds the schema definition for the CitySplider entity.
type CitySplider struct {
	ent.Schema
}

// Fields of the CitySplider.
func (CitySplider) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.String("source").Default("github.com").Comment("账号"),
		field.String("owner").Default("xiangyuecn").Comment("账号"),
		field.String("repo").Default("AreaCity-JsSpider-StatsGov").Comment("包名"),
		field.Uint64("source_id").Default(0).Comment("来源标识"),
		field.String("title").Default("").Comment("标题"),
		field.Bool("draft").Default(false).Comment("是否草稿"),
		field.Bool("pre_release").Default(false).Comment("是否预发布"),
		field.Uint64("publishe_time").Default(0).Comment("发版时间"),
		field.Uint8("status").Default(0).Comment("状态"),
		field.Uint64("create_time").Default(0).Comment("创建时间"),
		field.Uint64("update_time").Default(0).Comment("更新时间"),
	}
}

// Edges of the CitySplider.
func (CitySplider) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the schema.
func (CitySplider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "city_splider",
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Options:   "ENGINE = INNODB",
		},
	}
}
