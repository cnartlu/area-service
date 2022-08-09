package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cnartlu/area-service/pkg/ent/fields/mixin"
	"github.com/cnartlu/area-service/pkg/ent/validates"
)

// AreaReleaseAsset holds the schema definition for the AreaReleaseAsset entity.
type AreaReleaseAsset struct {
	ent.Schema
}

// Mixin of the schema.
func (AreaReleaseAsset) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the AreaReleaseAsset.
func (AreaReleaseAsset) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("ID"),
		field.Uint64("area_release_id").Comment("区域发版").Default(0),
		field.Uint64("asset_id").Comment("资源ID").Default(0),
		field.String("asset_name").Comment("资源名称").Default("").Validate(validates.MaxRuneCount(255)),
		field.String("asset_label").Comment("资源标签").Default("").Validate(validates.MaxRuneCount(255)),
		field.String("asset_state").Comment("资源状态").Default("").Validate(validates.MaxRuneCount(64)),
		field.String("file_path").Comment("文件路径").Default("").Validate(validates.MaxRuneCount(255)),
		field.Uint("file_size").Comment("文件大小").Default(0),
		field.String("download_url").Comment("下载地址").Default("").Validate(validates.MaxRuneCount(255)),
		field.Enum("status").
			NamedValues(
				"WaitLoaded", "0",
				"FinishedLoaded", "1",
			).
			Default("0").
			Comment("状态"),
	}
}

// Edges of the AreaReleaseAsset.
func (AreaReleaseAsset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("release", AreaRelease.Type).Ref("assets").Unique(),
	}
}

// Indexes of the schema.
func (AreaReleaseAsset) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("area_release_id"),
	}
}

// Annotations of the schema.
func (AreaReleaseAsset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "area_release_asset",
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Options:   "ENGINE = INNODB",
		},
	}
}
