package schema

import (
	"errors"
	"regexp"

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
		field.String("idnumber").MaxLen(13).MinLen(13),
		field.Int("age").Min(1),

		field.String("telephone").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[0]{1}[869]{1}[0-9]{8}$", s)
			if !match {
				return errors.New("ใส่เบอร์โทรไม่ถูกต้อง")
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
