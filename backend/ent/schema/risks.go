package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Risks holds the schema definition for the Risks entity.
type Risks struct {
	ent.Schema
}

// Fields of the Risks.
func (Risks) Fields() []ent.Field {
	return []ent.Field{
		field.String("Risks").NotEmpty(),
	}
}

// Edges of the Risks.
func (Risks) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Antenatalinformation", Antenatalinformation.Type).Unique(),
	}
}
