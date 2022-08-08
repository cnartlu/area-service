package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/cnartlu/area-service/pkg/ent/fields/mixin"

	"github.com/cnartlu/area-service/pkg/ent/validates"
)

const (
	Status1 = 0
	Status2 = 1
)

// AreaRelease holds the schema definition for the AreaRelease entity.
type AreaRelease struct {
	ent.Schema
}

// Mixin of the schema.
func (AreaRelease) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the AreaRelease.
func (AreaRelease) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("ID"),
		field.String("owner").Default("").Comment("账号").Validate(validates.MaxRuneCount(128)),
		field.String("repo").Default("").Comment("包名").Validate(validates.MaxRuneCount(128)),
		field.Uint64("release_id").Default(0).Comment("发版标识"),
		field.String("release_name").Default("").Comment("发版名称").Validate(validates.MaxRuneCount(128)),
		field.String("release_node_id").Default("").Comment("发版节点").Validate(validates.MaxRuneCount(128)),
		field.Uint64("release_published_at").Default(0).Comment("发版时间"),
		field.String("release_content").Default("").Comment("发版内容").SchemaType(map[string]string{
			dialect.MySQL: "mediumtext",
		}).Annotations(entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		}),
		field.Enum("status").
			NamedValues(
				"WaitLoaded", "0",
				"FinishLoaded", "1",
			).Default("0").
			Comment("状态"),
	}
}

// Edges of the AreaRelease.
func (AreaRelease) Edges() []ent.Edge {
	return nil
}

// Indexes of the schema.
func (AreaRelease) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations of the schema.
func (AreaRelease) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "area_release",
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Options:   "ENGINE = INNODB",
		},
	}
}
