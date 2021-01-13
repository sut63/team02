package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Pregnancystatus holds the schema definition for the Pregnancystatus entity.
type Pregnancystatus struct {
	ent.Schema
}

// Fields of the Pregnancystatus.
func (Pregnancystatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("Pregnancystatus").NotEmpty(),
	}
}

// Edges of the Pregnancystatus.
func (Pregnancystatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Antenatalinformation" , Antenatalinformation.Type).Unique(),
	}
}
