package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cnartlu/area-service/pkg/ent/fields/mixin"
	"github.com/cnartlu/area-service/pkg/ent/validates"
)

// Area holds the schema definition for the Area entity.
type Area struct {
	ent.Schema
}

// Mixin of the schema.
func (Area) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.DeleteTime{},
		mixin.Time{},
		mixin.Point{},
	}
}

// Fields of the Area.
func (Area) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("ID"),
		field.Uint64("parent_id").Comment("父级").Default(0),
		field.String("region_id").Comment("区域标识").Default("").Validate(validates.MaxRuneCount(20)),
		field.String("parent_list").Comment("族谱").Default("").Validate(validates.MaxRuneCount(255)),
		field.String("title").Comment("标题").Default("").Validate(validates.MaxRuneCount(36)),
		field.String("pinyin").Comment("拼音").Default("").Validate(validates.MaxRuneCount(128)),
		field.String("ucfirst").Comment("大写首字母").Default("").MaxLen(1).SchemaType(map[string]string{
			dialect.MySQL: "CHAR(1)",
		}),
		field.String("geohash").Comment("GeoHash算法值").Default("").Validate(validates.MaxRuneCount(12)).SchemaType(map[string]string{
			dialect.MySQL: "CHAR(12)",
		}),
		field.String("geo_gs2").Comment("GoogleGeo2算法").Default("").Validate(validates.MaxRuneCount(64)),
		field.Uint64("geo_gs2_id").Comment("点ID").Default(0),
		field.Uint32("geo_gs2_level").Comment("点级别").Default(0),
		field.Uint32("geo_gs2_face").Comment("面级别").Default(0),
		field.String("city_code").Comment("	城市编码").Default("").Validate(validates.MaxRuneCount(12)),
		field.String("zip_code").Comment("邮编编码").Default("").Validate(validates.MaxRuneCount(12)),
		field.Uint8("level").Comment("等级").Default(1).SchemaType(map[string]string{
			dialect.MySQL: "SMALLINT(3)",
		}),
		field.Uint32("children_number").Comment("子节点个数").Default(0),
	}
}

// Edges of the Area.
func (Area) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("childrens", Area.Type).Comment("子级").From("parent").Field("parent_id").Required().Unique(),
	}
}

// Indexes of the schema.
func (Area) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("parent_id"),
		index.Fields("region_id"),
	}
}

// Annotations of the schema.
func (Area) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "area",
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Options:   "ENGINE = INNODB",
		},
	}
}
