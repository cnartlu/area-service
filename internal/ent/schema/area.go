package schema

import "entgo.io/ent"

// Area holds the schema definition for the Area entity.
type Area struct {
	ent.Schema
}

// Fields of the Area.
func (Area) Fields() []ent.Field {
	return nil
}

// Edges of the Area.
func (Area) Edges() []ent.Edge {
	return nil
}
