package schema

import "github.com/facebookincubator/ent"

// Bonedisease holds the schema definition for the Bonedisease entity.
type Bonedisease struct {
	ent.Schema
}

// Fields of the Bonedisease.
func (Bonedisease) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedTime").Default(time.Now),
	}
}

// Edges of the Bonedisease.
func (Bonedisease) Edges() []ent.Edge {
	return []ent.Edge{
		Edge.From("Remedy", Remedy.Type).Ref("Bonedisease").Unique(),
		Edge.From("Orthopedics", Orthopedics.Type).Ref("Boneddisease").Unique(),
		Edge.From("Patient", Patient.Type).Ref("Bonedisease").Unique(),
		Edge.From("Advice", Advice.Type).Ref("Bonedisease").Unique(),
	}
}
