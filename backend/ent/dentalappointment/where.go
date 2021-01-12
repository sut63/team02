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

// AppointTime applies equality check predicate on the "appoint_time" field. It's identical to AppointTimeEQ.
func AppointTime(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointTime), v))
	})
}

// AppointTimeEQ applies the EQ predicate on the "appoint_time" field.
func AppointTimeEQ(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppointTime), v))
	})
}

// AppointTimeNEQ applies the NEQ predicate on the "appoint_time" field.
func AppointTimeNEQ(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppointTime), v))
	})
}

// AppointTimeIn applies the In predicate on the "appoint_time" field.
func AppointTimeIn(vs ...time.Time) predicate.Dentalappointment {
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
		s.Where(sql.In(s.C(FieldAppointTime), v...))
	})
}

// AppointTimeNotIn applies the NotIn predicate on the "appoint_time" field.
func AppointTimeNotIn(vs ...time.Time) predicate.Dentalappointment {
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
		s.Where(sql.NotIn(s.C(FieldAppointTime), v...))
	})
}

// AppointTimeGT applies the GT predicate on the "appoint_time" field.
func AppointTimeGT(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppointTime), v))
	})
}

// AppointTimeGTE applies the GTE predicate on the "appoint_time" field.
func AppointTimeGTE(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppointTime), v))
	})
}

// AppointTimeLT applies the LT predicate on the "appoint_time" field.
func AppointTimeLT(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppointTime), v))
	})
}

// AppointTimeLTE applies the LTE predicate on the "appoint_time" field.
func AppointTimeLTE(v time.Time) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppointTime), v))
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

// HasDentaltype applies the HasEdge predicate on the "Dentaltype" edge.
func HasDentaltype() predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DentaltypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DentaltypeTable, DentaltypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDentaltypeWith applies the HasEdge predicate on the "Dentaltype" edge with a given conditions (other predicates).
func HasDentaltypeWith(preds ...predicate.Dentaltype) predicate.Dentalappointment {
	return predicate.Dentalappointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DentaltypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DentaltypeTable, DentaltypeColumn),
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
