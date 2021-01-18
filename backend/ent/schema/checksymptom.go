package schema

import  (
	
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Checksymptom holds the schema definition for the Checksymptom entity.
type Checksymptom struct {
	ent.Schema
}

// Fields of the Checksymptom.
func (Checksymptom) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
		field.String("times").NotEmpty(),
		field.String("note").NotEmpty(),

	}
}

// Edges of the Checksymptom.
func (Checksymptom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("patient", Patient.Type).Ref("Checksymptom").Unique(),
		edge.From("personnel", Personnel.Type).Ref("Checksymptom").Unique(),
		edge.From("doctorordersheet", Doctorordersheet.Type).Ref("Checksymptom").Unique(),
		edge.From("disease", Disease.Type).Ref("Checksymptom").Unique(),
		
	}
}
