// Code generated by entc, DO NOT EDIT.

package status

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/to63/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Statusname applies equality check predicate on the "statusname" field. It's identical to StatusnameEQ.
func Statusname(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusname), v))
	})
}

// StatusnameEQ applies the EQ predicate on the "statusname" field.
func StatusnameEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusname), v))
	})
}

// StatusnameNEQ applies the NEQ predicate on the "statusname" field.
func StatusnameNEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusname), v))
	})
}

// StatusnameIn applies the In predicate on the "statusname" field.
func StatusnameIn(vs ...string) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatusname), v...))
	})
}

// StatusnameNotIn applies the NotIn predicate on the "statusname" field.
func StatusnameNotIn(vs ...string) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatusname), v...))
	})
}

// StatusnameGT applies the GT predicate on the "statusname" field.
func StatusnameGT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatusname), v))
	})
}

// StatusnameGTE applies the GTE predicate on the "statusname" field.
func StatusnameGTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatusname), v))
	})
}

// StatusnameLT applies the LT predicate on the "statusname" field.
func StatusnameLT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatusname), v))
	})
}

// StatusnameLTE applies the LTE predicate on the "statusname" field.
func StatusnameLTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatusname), v))
	})
}

// StatusnameContains applies the Contains predicate on the "statusname" field.
func StatusnameContains(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatusname), v))
	})
}

// StatusnameHasPrefix applies the HasPrefix predicate on the "statusname" field.
func StatusnameHasPrefix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatusname), v))
	})
}

// StatusnameHasSuffix applies the HasSuffix predicate on the "statusname" field.
func StatusnameHasSuffix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatusname), v))
	})
}

// StatusnameEqualFold applies the EqualFold predicate on the "statusname" field.
func StatusnameEqualFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatusname), v))
	})
}

// StatusnameContainsFold applies the ContainsFold predicate on the "statusname" field.
func StatusnameContainsFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatusname), v))
	})
}

// HasPhysicaltherapyrecord applies the HasEdge predicate on the "physicaltherapyrecord" edge.
func HasPhysicaltherapyrecord() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PhysicaltherapyrecordTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PhysicaltherapyrecordTable, PhysicaltherapyrecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPhysicaltherapyrecordWith applies the HasEdge predicate on the "physicaltherapyrecord" edge with a given conditions (other predicates).
func HasPhysicaltherapyrecordWith(preds ...predicate.Physicaltherapyrecord) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PhysicaltherapyrecordInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PhysicaltherapyrecordTable, PhysicaltherapyrecordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
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
func Not(p predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		p(s.Not())
	})
}
