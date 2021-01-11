package schema

import "github.com/facebookincubator/ent"

// Remedy holds the schema definition for the Remedy entity.
type Remedy struct {
	ent.Schema
}

// Fields of the Remedy.
func (Remedy) Fields() []ent.Field {
	return []ent.Field{
		Field.String("Remedy").NotEmpty(),
	}
}

// Edges of the Remedy.
func (Remedy) Edges() []ent.Edge {
	return []ent.Edge{
		Edge.To("Bonedisease", Bonedisease).StorageKey(edge.Column("remedy_id")),
	}
}
