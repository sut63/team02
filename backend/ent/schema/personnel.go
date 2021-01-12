package schema

import "github.com/facebookincubator/ent"

// Personnel holds the schema definition for the Personnel entity.
type Personnel struct {
	ent.Schema
}

// Fields of the Personnel.
func (Personnel) Fields() []ent.Field {
	return nil
}

// Edges of the Personnel.
func (Personnel) Edges() []ent.Edge {
	return nil
}
