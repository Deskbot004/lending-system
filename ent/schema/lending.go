package schema

import "entgo.io/ent"

// Lending holds the schema definition for the Lending entity.
type Lending struct {
	ent.Schema
}

// Fields of the Lending.
func (Lending) Fields() []ent.Field {
	return nil
}

// Edges of the Lending.
func (Lending) Edges() []ent.Edge {
	return nil
}
