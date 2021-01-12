package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Dentaltype holds the schema definition for the Dentaltype entity.
type Dentaltype struct {
	ent.Schema
}

// Fields of the Dentaltype.
func (Dentaltype) Fields() []ent.Field {
	return []ent.Field{
		field.String("typename"),
	}
}

// Edges of the Dentaltype.
func (Dentaltype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Dentalappointment", Dentalappointment.Type).StorageKey(edge.Column("typename")),
	}
}
