package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Antenatalinformation holds the schema definition for the Antenatalinformation entity.
type Antenatalinformation struct {
	ent.Schema
}

// Fields of the Antenatalinformation.
func (Antenatalinformation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("gestationalage").Positive(),
		field.Time("time"),
	}
}

// Edges of the Antenatalinformation.
func (Antenatalinformation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Personnel", Personnel.Type).Ref("Antenatalinformation").Unique(),
		edge.From("Patient", Patient.Type).Ref("Antenatalinformation").Unique(),
		edge.From("Pregnancystatus", Pregnancystatus.Type).Ref("Antenatalinformation").Unique(),
		edge.From("Risks", Risks.Type).Ref("Antenatalinformation").Unique(),
	}
}
