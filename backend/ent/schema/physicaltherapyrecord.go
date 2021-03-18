package schema

import (
	"regexp"
	"errors"

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
		field.Int("price").Min(1),
		field.String("note").
		Validate(func(s string) error {
			match, _ := regexp.MatchString("[ก-๙]", s)
			if !match {
				return errors.New("รูปแบบไม่ถูกต้อง")
			}
			return nil
		}),
		field.Time("appointtime"),
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
