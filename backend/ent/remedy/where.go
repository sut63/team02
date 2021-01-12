// Code generated by entc, DO NOT EDIT.

package remedy

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/to63/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Remedy applies equality check predicate on the "remedy" field. It's identical to RemedyEQ.
func Remedy(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemedy), v))
	})
}

// RemedyEQ applies the EQ predicate on the "remedy" field.
func RemedyEQ(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemedy), v))
	})
}

// RemedyNEQ applies the NEQ predicate on the "remedy" field.
func RemedyNEQ(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemedy), v))
	})
}

// RemedyIn applies the In predicate on the "remedy" field.
func RemedyIn(vs ...string) predicate.Remedy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remedy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRemedy), v...))
	})
}

// RemedyNotIn applies the NotIn predicate on the "remedy" field.
func RemedyNotIn(vs ...string) predicate.Remedy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remedy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRemedy), v...))
	})
}

// RemedyGT applies the GT predicate on the "remedy" field.
func RemedyGT(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemedy), v))
	})
}

// RemedyGTE applies the GTE predicate on the "remedy" field.
func RemedyGTE(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemedy), v))
	})
}

// RemedyLT applies the LT predicate on the "remedy" field.
func RemedyLT(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemedy), v))
	})
}

// RemedyLTE applies the LTE predicate on the "remedy" field.
func RemedyLTE(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemedy), v))
	})
}

// RemedyContains applies the Contains predicate on the "remedy" field.
func RemedyContains(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemedy), v))
	})
}

// RemedyHasPrefix applies the HasPrefix predicate on the "remedy" field.
func RemedyHasPrefix(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemedy), v))
	})
}

// RemedyHasSuffix applies the HasSuffix predicate on the "remedy" field.
func RemedyHasSuffix(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemedy), v))
	})
}

// RemedyEqualFold applies the EqualFold predicate on the "remedy" field.
func RemedyEqualFold(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemedy), v))
	})
}

// RemedyContainsFold applies the ContainsFold predicate on the "remedy" field.
func RemedyContainsFold(v string) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemedy), v))
	})
}

// HasBonedisease applies the HasEdge predicate on the "Bonedisease" edge.
func HasBonedisease() predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BonediseaseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BonediseaseTable, BonediseaseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBonediseaseWith applies the HasEdge predicate on the "Bonedisease" edge with a given conditions (other predicates).
func HasBonediseaseWith(preds ...predicate.Bonedisease) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BonediseaseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BonediseaseTable, BonediseaseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Remedy) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Remedy) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
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
func Not(p predicate.Remedy) predicate.Remedy {
	return predicate.Remedy(func(s *sql.Selector) {
		p(s.Not())
	})
}
