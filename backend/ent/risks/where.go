// Code generated by entc, DO NOT EDIT.

package risks

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/to63/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Risks applies equality check predicate on the "Risks" field. It's identical to RisksEQ.
func Risks(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRisks), v))
	})
}

// RisksEQ applies the EQ predicate on the "Risks" field.
func RisksEQ(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRisks), v))
	})
}

// RisksNEQ applies the NEQ predicate on the "Risks" field.
func RisksNEQ(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRisks), v))
	})
}

// RisksIn applies the In predicate on the "Risks" field.
func RisksIn(vs ...string) predicate.Risks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Risks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRisks), v...))
	})
}

// RisksNotIn applies the NotIn predicate on the "Risks" field.
func RisksNotIn(vs ...string) predicate.Risks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Risks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRisks), v...))
	})
}

// RisksGT applies the GT predicate on the "Risks" field.
func RisksGT(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRisks), v))
	})
}

// RisksGTE applies the GTE predicate on the "Risks" field.
func RisksGTE(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRisks), v))
	})
}

// RisksLT applies the LT predicate on the "Risks" field.
func RisksLT(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRisks), v))
	})
}

// RisksLTE applies the LTE predicate on the "Risks" field.
func RisksLTE(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRisks), v))
	})
}

// RisksContains applies the Contains predicate on the "Risks" field.
func RisksContains(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRisks), v))
	})
}

// RisksHasPrefix applies the HasPrefix predicate on the "Risks" field.
func RisksHasPrefix(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRisks), v))
	})
}

// RisksHasSuffix applies the HasSuffix predicate on the "Risks" field.
func RisksHasSuffix(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRisks), v))
	})
}

// RisksEqualFold applies the EqualFold predicate on the "Risks" field.
func RisksEqualFold(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRisks), v))
	})
}

// RisksContainsFold applies the ContainsFold predicate on the "Risks" field.
func RisksContainsFold(v string) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRisks), v))
	})
}

// HasAntenatalinformation applies the HasEdge predicate on the "Antenatalinformation" edge.
func HasAntenatalinformation() predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AntenatalinformationTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AntenatalinformationTable, AntenatalinformationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAntenatalinformationWith applies the HasEdge predicate on the "Antenatalinformation" edge with a given conditions (other predicates).
func HasAntenatalinformationWith(preds ...predicate.Antenatalinformation) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AntenatalinformationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AntenatalinformationTable, AntenatalinformationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Risks) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Risks) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
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
func Not(p predicate.Risks) predicate.Risks {
	return predicate.Risks(func(s *sql.Selector) {
		p(s.Not())
	})
}
