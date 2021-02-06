package schema

import (
	"time"
	"regexp"
	"errors"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Bonedisease holds the schema definition for the Bonedisease entity.
type Bonedisease struct {
	ent.Schema
}

// Fields of the Bonedisease.
func (Bonedisease) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedTime").Default(time.Now),
		field.String("advice").
			Validate(func(s string) error {
				match, _ := regexp.MatchString("[ก-๘]", s)
				if !match {
					return errors.New("รูปแบบไม่ถูกต้อง")
				}
				return nil
			}),
		field.String("tel").NotEmpty().MinLen(10).MaxLen(10).
			Validate(func (s string) error {
				match, _ := regexp.MatchString("^[0]{1}[8926]{1}[0-9]{8}$", s)
				if !match {
					return errors.New("รูปแบบเบอร์ไม่ถูกต้อง")
				}
				return nil
			}),
		field.String("identificationCard").NotEmpty().MinLen(13).MaxLen(13).
			Validate(func (s string) error {
				match, _ := regexp.MatchString("^[0-9]{13}$", s)
				if !match {
					return errors.New("รูปแบบเลขบัตรประชาชนไม่ถูกต้องไม่ถูกต้อง")
				}
				return nil
			}),
	}
}

// Edges of the Bonedisease.
func (Bonedisease) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("remedy", Remedy.Type).Ref("Bonedisease").Unique(),
		edge.From("patient", Patient.Type).Ref("Bonedisease").Unique(),
		edge.From("personnel", Personnel.Type).Ref("Bonedisease").Unique(),
	}
}
