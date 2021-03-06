// Code generated by entc, DO NOT EDIT.

package dentalappointment

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/to63/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Appointtime applies equality check predicate on the "appointtime" field. It's identical to AppointtimeEQ.
func Appointtime(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointtime), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// AppointtimeEQ applies the EQ predicate on the "appointtime" field.
func AppointtimeEQ(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointtime), v))
	})
}

// AppointtimeNEQ applies the NEQ predicate on the "appointtime" field.
func AppointtimeNEQ(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppointtime), v))
	})
}

// AppointtimeIn applies the In predicate on the "appointtime" field.
func AppointtimeIn(vs ...time.Time) predicate.Dentalappointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentalappointment(func(s *sql.Selector) {
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
func AppointtimeNotIn(vs ...time.Time) predicate.Dentalappointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentalappointment(func(s *sql.Selector) {
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
func AppointtimeGT(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppointtime), v))
	})
}

// AppointtimeGTE applies the GTE predicate on the "appointtime" field.
func AppointtimeGTE(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppointtime), v))
	})
}

// AppointtimeLT applies the LT predicate on the "appointtime" field.
func AppointtimeLT(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppointtime), v))
	})
}

// AppointtimeLTE applies the LTE predicate on the "appointtime" field.
func AppointtimeLTE(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppointtime), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int) predicate.Dentalappointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentalappointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int) predicate.Dentalappointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentalappointment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int) predicate.Dentalappointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentalappointment(func(s *sql.Selector) {
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
func PriceNotIn(vs ...int) predicate.Dentalappointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentalappointment(func(s *sql.Selector) {
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
func PriceGT(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNote), v))
	})
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Dentalappointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentalappointment(func(s *sql.Selector) {
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
func NoteNotIn(vs ...string) predicate.Dentalappointment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dentalappointment(func(s *sql.Selector) {
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
func NoteGT(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNote), v))
	})
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNote), v))
	})
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNote), v))
	})
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNote), v))
	})
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNote), v))
	})
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNote), v))
	})
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNote), v))
	})
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNote), v))
	})
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNote), v))
	})
}

// HasPersonnel applies the HasEdge predicate on the "Personnel" edge.
func HasPersonnel() predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonnelTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonnelTable, PersonnelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonnelWith applies the HasEdge predicate on the "Personnel" edge with a given conditions (other predicates).
func HasPersonnelWith(preds ...predicate.Personnel) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
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

// HasPatient applies the HasEdge predicate on the "Patient" edge.
func HasPatient() predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "Patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
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

// HasDentalkind applies the HasEdge predicate on the "Dentalkind" edge.
func HasDentalkind() predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DentalkindTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DentalkindTable, DentalkindColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDentalkindWith applies the HasEdge predicate on the "Dentalkind" edge with a given conditions (other predicates).
func HasDentalkindWith(preds ...predicate.Dentalkind) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DentalkindInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DentalkindTable, DentalkindColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Dentalappointment) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Dentalappointment) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
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
func Not(p predicate.Dentalappointment) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		p(s.Not())
	})
}
