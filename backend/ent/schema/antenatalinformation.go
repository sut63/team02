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
		field.String("gestationalage").NotEmpty(),
		field.Time("added_time"),
	}
}

// Edges of the Antenatalinformation.
func (Antenatalinformation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Personnel", Personnel.Type).Ref("Antenatalinformation").Unique(),
		edge.From("Patient", Patient.Type).Ref("Antenatalinformation").Unique(),
		edge.From("Pregnancystatusid", Pregnancystatus.Type).Ref("Antenatalinformation").Unique(),
		edge.From("Risksid", Risks.Type).Ref("Antenatalinformation").Unique(),
	}
}
