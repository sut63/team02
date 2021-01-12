package schema

import  (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Bonedisease holds the schema definition for the Bonedisease entity.
type Bonedisease struct {
	ent.Schema
}

// Fields of the Bonedisease.
func (Bonedisease) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedTime").Default(time.Now),
		field.String("advice").NotEmpty(),
	}
}

// Edges of the Bonedisease.
func (Bonedisease) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("remedy", Remedy.Type).Ref("Bonedisease").Unique(),
		edge.From("patient", Patient.Type).Ref("Bonedisease").Unique(),
		edge.From("personnel", Personnel.Type).Ref("Bonedisease").Unique(),
	}
}
