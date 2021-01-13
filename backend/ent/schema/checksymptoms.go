package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Checksymptoms holds the schema definition for the Checksymptoms entity.
type Checksymptoms struct {
	ent.Schema
}

// Fields of the Checksymptoms.
func (Checksymptoms) Fields() []ent.Field {
	return nil

}

// Edges of the Checksymptoms.
func (Checksymptoms) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("disease", Disease.Type).Ref("Checksymptoms").Unique(),
		edge.From("patient", Patient.Type).Ref("Checksymptoms").Unique(),
		edge.From("personnel", Personnel.Type).Ref("Checksymptoms").Unique(),
		edge.From("doctorordersheet", DoctorOrderSheet.Type).Ref("Checksymptoms").Unique(),
	}
}
