package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Dentalkind holds the schema definition for the Dentalkind entity.
type Dentalkind struct {
	ent.Schema
}

// Fields of the Dentalkind.
func (Dentalkind) Fields() []ent.Field {
	return []ent.Field{
		field.String("kindname"),
	}
}

// Edges of the Dentalkind.
func (Dentalkind) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Dentalappointment", Dentalappointment.Type).StorageKey(edge.Column("kindname")),
	}
}
