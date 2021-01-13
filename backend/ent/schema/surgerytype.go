package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Surgerytype holds the schema definition for the Surgerytype entity.
type Surgerytype struct {
	ent.Schema
}

// Fields of the Surgerytype.
func (Surgerytype) Fields() []ent.Field {
	return []ent.Field{
		field.String("typename").NotEmpty(),
	}
}

// Edges of the Surgerytype.
func (Surgerytype) Edges() []ent.Edge {
	return nil
}
