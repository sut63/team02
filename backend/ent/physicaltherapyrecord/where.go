// Code generated by entc, DO NOT EDIT.

package physicaltherapyrecord

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/to63/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// Appointtime applies equality check predicate on the "appointtime" field. It's identical to AppointtimeEQ.
func Appointtime(v time.Time) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointtime), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.Physicaltherapyrecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int) predicate.Physicaltherapyrecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Physicaltherapyrecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Physicaltherapyrecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNote), v))
	})
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Physicaltherapyrecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNote), v...))
	})
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.Physicaltherapyrecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNote), v...))
	})
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNote), v))
	})
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNote), v))
	})
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNote), v))
	})
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNote), v))
	})
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNote), v))
	})
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNote), v))
	})
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNote), v))
	})
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNote), v))
	})
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNote), v))
	})
}

// AppointtimeEQ applies the EQ predicate on the "appointtime" field.
func AppointtimeEQ(v time.Time) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointtime), v))
	})
}

// AppointtimeNEQ applies the NEQ predicate on the "appointtime" field.
func AppointtimeNEQ(v time.Time) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppointtime), v))
	})
}

// AppointtimeIn applies the In predicate on the "appointtime" field.
func AppointtimeIn(vs ...time.Time) predicate.Physicaltherapyrecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppointtime), v...))
	})
}

// AppointtimeNotIn applies the NotIn predicate on the "appointtime" field.
func AppointtimeNotIn(vs ...time.Time) predicate.Physicaltherapyrecord {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppointtime), v...))
	})
}

// AppointtimeGT applies the GT predicate on the "appointtime" field.
func AppointtimeGT(v time.Time) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppointtime), v))
	})
}

// AppointtimeGTE applies the GTE predicate on the "appointtime" field.
func AppointtimeGTE(v time.Time) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppointtime), v))
	})
}

// AppointtimeLT applies the LT predicate on the "appointtime" field.
func AppointtimeLT(v time.Time) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppointtime), v))
	})
}

// AppointtimeLTE applies the LTE predicate on the "appointtime" field.
func AppointtimeLTE(v time.Time) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppointtime), v))
	})
}

// HasPersonnel applies the HasEdge predicate on the "personnel" edge.
func HasPersonnel() predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonnelTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonnelTable, PersonnelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonnelWith applies the HasEdge predicate on the "personnel" edge with a given conditions (other predicates).
func HasPersonnelWith(preds ...predicate.Personnel) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonnelInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonnelTable, PersonnelColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatient applies the HasEdge predicate on the "patient" edge.
func HasPatient() predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPhysicaltherapyroom applies the HasEdge predicate on the "physicaltherapyroom" edge.
func HasPhysicaltherapyroom() predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PhysicaltherapyroomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PhysicaltherapyroomTable, PhysicaltherapyroomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPhysicaltherapyroomWith applies the HasEdge predicate on the "physicaltherapyroom" edge with a given conditions (other predicates).
func HasPhysicaltherapyroomWith(preds ...predicate.Physicaltherapyroom) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PhysicaltherapyroomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PhysicaltherapyroomTable, PhysicaltherapyroomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatus applies the HasEdge predicate on the "status" edge.
func HasStatus() predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatusTable, StatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusWith applies the HasEdge predicate on the "status" edge with a given conditions (other predicates).
func HasStatusWith(preds ...predicate.Status) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatusTable, StatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Physicaltherapyrecord) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Physicaltherapyrecord) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Physicaltherapyrecord) predicate.Physicaltherapyrecord {
	return predicate.Physicaltherapyrecord(func(s *sql.Selector) {
		p(s.Not())
	})
}
