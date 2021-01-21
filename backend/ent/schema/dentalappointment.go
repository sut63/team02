package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Dentalappointment holds the schema definition for the Dentalappointment entity.
type Dentalappointment struct {
	ent.Schema
}

// Fields of the Dentalappointment.
func (Dentalappointment) Fields() []ent.Field {
	return []ent.Field{
		field.Time("appointtime"),
		field.Int("amount").Min(1),
		field.Int("price").Min(1),
		field.String("note").MinLen(1).MaxLen(25),
	}
}

// Edges of the Dentalappointment.
func (Dentalappointment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Personnel", Personnel.Type).Ref("Dentalappointment").Unique(),
		edge.From("Patient", Patient.Type).Ref("Dentalappointment").Unique(),
		edge.From("Dentalkind", Dentalkind.Type).Ref("Dentalappointment").Unique(),
	}
}
