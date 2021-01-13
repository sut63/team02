package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Personnel holds the schema definition for the Personnel entity.
type Personnel struct {
	ent.Schema
}

// Fields of the Personnel.
func (Personnel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("department").NotEmpty(),
		field.String("user").NotEmpty(),
		field.String("password").NotEmpty(),
	}
}

// Edges of the Personnel.
func (Personnel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("physicaltherapyrecord", Physicaltherapyrecord.Type).StorageKey(edge.Column("Personnel_id")),
		edge.To("Bonedisease", Bonedisease.Type).StorageKey(edge.Column("Personnel_id")),
		edge.To("Checksymptoms", Checksymptoms.Type).StorageKey(edge.Column("Personnel_id")),
		edge.To("Dentalappointment", Dentalappointment.Type).StorageKey(edge.Column("Personnel_id")),
		edge.To("Surgeryappointment", Surgeryappointment.Type).StorageKey(edge.Column("Personnel_id")),
	}
}
