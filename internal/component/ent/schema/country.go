package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cnartlu/area-service/pkg/ent/fields/mixin"
	"github.com/google/uuid"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Mixin of the schema.
func (Country) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.DeleteTime{},
		mixin.Time{},
	}
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("ID"),
		field.UUID("uuid", uuid.New()).Comment("UUID").Unique(),
		field.String("title").Default("").Comment("标题"),
		field.String("en_title").Default("").Comment("英文标题"),
		field.String("cn_title").Default("").Comment("中文标题"),
		field.Uint64("founding_time").Default(0).Comment("创立时间"),
		field.String("two_digit_code").Default("").Comment("ISO3166-1两位代码"),
		field.String("there_digit_code").Default("").Comment("ISO3166-1三位代码"),
		field.Uint("number_code").Default(0).Comment("ISO3166-1数值代码"),
		field.String("iso3166_2").Default("").Comment("ISO3166-2代码"),
		field.String("is_sovereignty").Default("").Comment("主权国"),
		field.String("note").Default("").Comment("备注"),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return nil
}

// Indexes of the schema.
func (Country) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid").Unique(),
	}
}

// Annotations of the schema.
func (Country) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "country",
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Options:   "ENGINE = INNODB",
		},
	}
}
