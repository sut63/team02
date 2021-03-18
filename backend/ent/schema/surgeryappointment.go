package schema

import (
	"errors"
	"regexp"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Surgeryappointment holds the schema definition for the Surgeryappointment entity.
type Surgeryappointment struct {
	ent.Schema
}

// Fields of the Surgeryappointment.
func (Surgeryappointment) Fields() []ent.Field {
	return []ent.Field{
		field.Time("appoint_time"),
		field.String("phone").NotEmpty().Validate(func(s string) error {
			match, _ := regexp.MatchString("^[0]{1}[896]{1}[0-9]{8}$", s)
			if !match {
				return errors.New("รูปแบบเลขทะเบียนรถไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("note").NotEmpty().Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙]+$", s)
			if !match {
				return errors.New("รูปแบบเลขทะเบียนรถไม่ถูกต้อง")
			}
			return nil
		}),
		field.Int("cost").Range(0, 1000000),
	}
}

// Edges of the Surgeryappointment.
func (Surgeryappointment) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("Personnel", Personnel.Type).Ref("Surgeryappointment").Unique(),
		edge.From("Patient", Patient.Type).Ref("Surgeryappointment").Unique(),
		edge.From("Surgerytype", Surgerytype.Type).Ref("Surgeryappointment").Unique(),
	}
}
