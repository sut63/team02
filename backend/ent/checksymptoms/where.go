// Code generated by entc, DO NOT EDIT.

package checksymptoms

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/to63/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Checksymptoms {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checksymptoms(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Checksymptoms {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Checksymptoms(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// HasDisease applies the HasEdge predicate on the "disease" edge.
func HasDisease() predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DiseaseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DiseaseTable, DiseaseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDiseaseWith applies the HasEdge predicate on the "disease" edge with a given conditions (other predicates).
func HasDiseaseWith(preds ...predicate.Disease) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DiseaseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DiseaseTable, DiseaseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatient applies the HasEdge predicate on the "patient" edge.
func HasPatient() predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
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

// HasPersonnel applies the HasEdge predicate on the "personnel" edge.
func HasPersonnel() predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonnelTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonnelTable, PersonnelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonnelWith applies the HasEdge predicate on the "personnel" edge with a given conditions (other predicates).
func HasPersonnelWith(preds ...predicate.Personnel) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
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

// HasDoctorordersheet applies the HasEdge predicate on the "doctorordersheet" edge.
func HasDoctorordersheet() predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorordersheetTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorordersheetTable, DoctorordersheetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorordersheetWith applies the HasEdge predicate on the "doctorordersheet" edge with a given conditions (other predicates).
func HasDoctorordersheetWith(preds ...predicate.DoctorOrderSheet) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorordersheetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorordersheetTable, DoctorordersheetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Checksymptoms) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Checksymptoms) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
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
func Not(p predicate.Checksymptoms) predicate.Checksymptoms {
	return predicate.Checksymptoms(func(s *sql.Selector) {
		p(s.Not())
	})
}
