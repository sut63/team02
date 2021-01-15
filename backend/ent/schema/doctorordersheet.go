package schema

import  (

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)
// Doctorordersheet holds the schema definition for the Doctorordersheet entity.
type Doctorordersheet struct {
	ent.Schema
}

// Fields of the DoctorOrderSheet.
func (Doctorordersheet) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").NotEmpty(),
		
	}
}

// Edges of the Doctorordersheet.
func (Doctorordersheet) Edges() []ent.Edge {
	return []ent.Edge{
	edge.To("Checksymptom", Checksymptom.Type).StorageKey(edge.Column("Doctorordersheet_id")),
	}
}
