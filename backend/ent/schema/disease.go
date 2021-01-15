package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Disease holds the schema definition for the Disease entity.
type Disease struct {
	ent.Schema
}

// Fields of the Disease.
func (Disease) Fields() []ent.Field {
	return []ent.Field{
		field.String("disease").NotEmpty(),
	}
}

// Edges of the Disease.
func (Disease) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Checksymptom", Checksymptom.Type).StorageKey(edge.Column("disease_id")),
	}
}
