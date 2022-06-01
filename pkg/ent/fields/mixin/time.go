package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// nowTimeToUint64 当前时间转为uint64时间戳
func nowTimeToUint64() uint64 {
	return uint64(time.Now().Unix())
}

// DeleteTime 删除时间字段
type DeleteTime struct {
	mixin.Schema
}

func (DeleteTime) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("delete_time").Comment("删除时间").Default(0),
	}
}

// delete time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*DeleteTime)(nil)

// CreateTime 创建时间字段
type CreateTime struct {
	mixin.Schema
}

func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("create_time").Comment("创建时间").Immutable().DefaultFunc(nowTimeToUint64),
	}
}

// create time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*CreateTime)(nil)

// UpdateTime 更新时间字段
type UpdateTime struct {
	mixin.Schema
}

func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("update_time").Comment("更新时间").DefaultFunc(nowTimeToUint64).UpdateDefault(nowTimeToUint64),
	}
}

// update time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*UpdateTime)(nil)

// Time 时间戳字段
type Time struct {
	mixin.Schema
}

func (Time) Fields() []ent.Field {
	return append(
		CreateTime{}.Fields(),
		UpdateTime{}.Fields()...,
	)
}

// time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*Time)(nil)
