package schema

import  (
	
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)
// DoctorOrderSheet holds the schema definition for the DoctorOrderSheet entity.
type DoctorOrderSheet struct {
	ent.Schema
}

// Fields of the DoctorOrderSheet.
func (DoctorOrderSheet) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").NotEmpty(),
		field.String("time").NotEmpty(),
		field.String("note").NotEmpty(),
	}
}

// Edges of the DoctorOrderSheet.
func (DoctorOrderSheet) Edges() []ent.Edge {
	return []ent.Edge{
	edge.To("Checksymptoms", Checksymptoms.Type).StorageKey(edge.Column("DoctorOrderSheet_id")),
	}
}
