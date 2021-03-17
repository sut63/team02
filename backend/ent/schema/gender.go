package schema

import (
        "github.com/facebook/ent"
        "github.com/facebook/ent/schema/field"
        "github.com/facebook/ent/schema/edge"
    )

// Gender holds the schema definition for the Gender entity.
type Gender struct {
    ent.Schema
}

// Fields of the Gender.
func (Gender) Fields() []ent.Field {
    return []ent.Field{
        field.String("gender").NotEmpty(),
    }
}

// Edges of the Gender.
func (Gender) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("Patient", Patient.Type).StorageKey(edge.Column("Gender")),
    }
}