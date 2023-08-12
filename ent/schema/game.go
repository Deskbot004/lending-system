package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
		field.String("type").
			Default("unknown"),
		field.String("ou").
			Default("unknown"),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return nil
}
