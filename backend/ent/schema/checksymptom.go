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
		field.String("note").MaxLen(40),
		field.String("Identitycard").MaxLen(13).MinLen(13),
		field.String("phone").MaxLen(10).MinLen(10),

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
