package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("birthday").NotEmpty(),
		field.String("gender").NotEmpty(),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("physicaltherapyrecord", Physicaltherapyrecord.Type).StorageKey(edge.Column("Patient_id")),
		edge.To("Bonedisease", Bonedisease.Type).StorageKey(edge.Column("Patient_id")),
		edge.To("Checksymptom", Checksymptom.Type).StorageKey(edge.Column("Patient_id")),
		edge.To("Dentalappointment", Dentalappointment.Type).StorageKey(edge.Column("Patient_id")),
		edge.To("Surgeryappointment", Surgeryappointment.Type).StorageKey(edge.Column("Patient_id")),
		edge.To("Antenatalinformation", Antenatalinformation.Type).StorageKey(edge.Column("Patient_id")),
	}
}
