package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Area holds the schema definition for the Area entity.
type Area struct {
	ent.Schema
}

// Fields of the Area.
func (Area) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("parent_id").Default(0).Comment("父级"),
		field.String("region_id").Default("").Comment("区域标识"),
		field.String("parent_list").Default("").Comment("族谱"),
		field.String("title").Default("").Comment("标题"),
		field.String("pinyin").Default("").Comment("拼音"),
		field.String("ucfirst").Default("").MaxLen(1).SchemaType(map[string]string{dialect.MySQL: "CHAR(1)"}).Comment("大写首字母"),
		field.Float("lat").Default(0).SchemaType(map[string]string{dialect.MySQL: "DECIMAL(10,7)"}).Comment("纬度"),
		field.Float("lng").Default(0).SchemaType(map[string]string{dialect.MySQL: "DECIMAL(10,7)"}).Comment("纬度"),
		field.String("geohash").Default("").SchemaType(map[string]string{dialect.MySQL: "CHAR(12)"}).Comment("GeoHash算法值"),
		field.String("geo_gs2").Default("").Comment("GoogleGeo2算法"),
		field.Uint64("geo_gs2_id").Default(0).Comment("点ID"),
		field.Uint32("geo_gs2_level").Default(0).Comment("点级别"),
		field.Uint32("geo_gs2_face").Default(0).Comment("面级别"),
		field.String("city_code").Default("").Comment("	城市编码"),
		field.String("zip_code").Default("").Comment("邮编编码"),
		field.Uint8("level").Default(1).SchemaType(map[string]string{dialect.MySQL: "SMALLINT(3)"}).Comment("等级"),
		field.Uint32("children_number").Default(0).Comment("子节点个数"),
		field.Uint64("delete_time").Default(0).Comment("删除时间"),
		field.Uint64("create_time").Default(0).Comment("创建时间"),
		field.Uint64("update_time").Default(0).Comment("更新时间"),
	}
}

// Edges of the Area.
func (Area) Edges() []ent.Edge {
	return []ent.Edge{}
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
