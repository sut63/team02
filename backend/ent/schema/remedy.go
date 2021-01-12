package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Remedy holds the schema definition for the Remedy entity.
type Remedy struct {
	ent.Schema
}

// Fields of the Remedy.
func (Remedy) Fields() []ent.Field {
	return []ent.Field{
		field.String("remedy").NotEmpty(),
	}
}

// Edges of the Remedy.
func (Remedy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Bonedisease", Bonedisease.Type).StorageKey(edge.Column("remedy_id")),
	}
}
