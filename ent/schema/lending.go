package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Lending holds the schema definition for the Lending entity.
type Lending struct {
	ent.Schema
}

// Fields of the Lending.
func (Lending) Fields() []ent.Field {
	return []ent.Field{
		field.String("date").
			Default("unknown"),
		field.String("notes").
			Default("unknown"),
	}
}

// Edges of the Lending.
func (Lending) Edges() []ent.Edge {
	return nil
}
