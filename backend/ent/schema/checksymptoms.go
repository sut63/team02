package schema

import (
	"time"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Checksymptoms holds the schema definition for the Checksymptoms entity.
type Checksymptoms struct {
	ent.Schema
}

// Fields of the Checksymptoms.
func (Checksymptoms) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date").Default(time.Now),
	}
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
