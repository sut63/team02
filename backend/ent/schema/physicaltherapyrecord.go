package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Physicaltherapyrecord holds the schema definition for the Physicaltherapyrecord entity.
type Physicaltherapyrecord struct {
	ent.Schema
}

// Fields of the Physicaltherapyrecord.
func (Physicaltherapyrecord) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedTime").Default(time.Now),
	}
}

// Edges of the Physicaltherapyrecord.
func (Physicaltherapyrecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("personnel", Personnel.Type).Ref("physicaltherapyrecord").Unique(),
		edge.From("patient", Patient.Type).Ref("physicaltherapyrecord").Unique(),
		edge.From("physicaltherapyroom", Physicaltherapyroom.Type).Ref("physicaltherapyrecord").Unique(),
		edge.From("status", Status.Type).Ref("physicaltherapyrecord").Unique(),
	}
}
