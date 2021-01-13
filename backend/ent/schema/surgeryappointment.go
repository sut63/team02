package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Surgeryappointment holds the schema definition for the Surgeryappointment entity.
type Surgeryappointment struct {
	ent.Schema
}

// Fields of the Surgeryappointment.
func (Surgeryappointment) Fields() []ent.Field {
	return []ent.Field{
		field.Time("appoint_time"),
	}
}

// Edges of the Surgeryappointment.
func (Surgeryappointment) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("Personnel", Personnel.Type).Ref("Surgeryappointment").Unique(),
		edge.From("Patient", Patient.Type).Ref("Surgeryappointment").Unique(),
	}
}
