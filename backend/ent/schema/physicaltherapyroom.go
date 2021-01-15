package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Physicaltherapyroom holds the schema definition for the Physicaltherapyroom entity.
type Physicaltherapyroom struct {
	ent.Schema
}

// Fields of the Physicaltherapyroom.
func (Physicaltherapyroom) Fields() []ent.Field {
	return []ent.Field{
		field.String("physicaltherapyroomname").NotEmpty(),
	}
}

// Edges of the Physicaltherapyroom.
func (Physicaltherapyroom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("physicaltherapyrecord", Physicaltherapyrecord.Type).Unique(),
	}
}
