// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package comment

import (
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Comment {
	return predicate.Comment(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i], _ = strconv.Atoi(ids[i])
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i], _ = strconv.Atoi(ids[i])
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// UniqueInt applies equality check predicate on the "unique_int" field. It's identical to UniqueIntEQ.
func UniqueInt(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUniqueInt), v))
	},
	)
}

// UniqueFloat applies equality check predicate on the "unique_float" field. It's identical to UniqueFloatEQ.
func UniqueFloat(v float64) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUniqueFloat), v))
	},
	)
}

// NillableInt applies equality check predicate on the "nillable_int" field. It's identical to NillableIntEQ.
func NillableInt(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNillableInt), v))
	},
	)
}

// UniqueIntEQ applies the EQ predicate on the "unique_int" field.
func UniqueIntEQ(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUniqueInt), v))
	},
	)
}

// UniqueIntNEQ applies the NEQ predicate on the "unique_int" field.
func UniqueIntNEQ(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUniqueInt), v))
	},
	)
}

// UniqueIntIn applies the In predicate on the "unique_int" field.
func UniqueIntIn(vs ...int) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUniqueInt), v...))
	},
	)
}

// UniqueIntNotIn applies the NotIn predicate on the "unique_int" field.
func UniqueIntNotIn(vs ...int) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUniqueInt), v...))
	},
	)
}

// UniqueIntGT applies the GT predicate on the "unique_int" field.
func UniqueIntGT(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUniqueInt), v))
	},
	)
}

// UniqueIntGTE applies the GTE predicate on the "unique_int" field.
func UniqueIntGTE(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUniqueInt), v))
	},
	)
}

// UniqueIntLT applies the LT predicate on the "unique_int" field.
func UniqueIntLT(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUniqueInt), v))
	},
	)
}

// UniqueIntLTE applies the LTE predicate on the "unique_int" field.
func UniqueIntLTE(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUniqueInt), v))
	},
	)
}

// UniqueFloatEQ applies the EQ predicate on the "unique_float" field.
func UniqueFloatEQ(v float64) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUniqueFloat), v))
	},
	)
}

// UniqueFloatNEQ applies the NEQ predicate on the "unique_float" field.
func UniqueFloatNEQ(v float64) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUniqueFloat), v))
	},
	)
}

// UniqueFloatIn applies the In predicate on the "unique_float" field.
func UniqueFloatIn(vs ...float64) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUniqueFloat), v...))
	},
	)
}

// UniqueFloatNotIn applies the NotIn predicate on the "unique_float" field.
func UniqueFloatNotIn(vs ...float64) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUniqueFloat), v...))
	},
	)
}

// UniqueFloatGT applies the GT predicate on the "unique_float" field.
func UniqueFloatGT(v float64) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUniqueFloat), v))
	},
	)
}

// UniqueFloatGTE applies the GTE predicate on the "unique_float" field.
func UniqueFloatGTE(v float64) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUniqueFloat), v))
	},
	)
}

// UniqueFloatLT applies the LT predicate on the "unique_float" field.
func UniqueFloatLT(v float64) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUniqueFloat), v))
	},
	)
}

// UniqueFloatLTE applies the LTE predicate on the "unique_float" field.
func UniqueFloatLTE(v float64) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUniqueFloat), v))
	},
	)
}

// NillableIntEQ applies the EQ predicate on the "nillable_int" field.
func NillableIntEQ(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNillableInt), v))
	},
	)
}

// NillableIntNEQ applies the NEQ predicate on the "nillable_int" field.
func NillableIntNEQ(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNillableInt), v))
	},
	)
}

// NillableIntIn applies the In predicate on the "nillable_int" field.
func NillableIntIn(vs ...int) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNillableInt), v...))
	},
	)
}

// NillableIntNotIn applies the NotIn predicate on the "nillable_int" field.
func NillableIntNotIn(vs ...int) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNillableInt), v...))
	},
	)
}

// NillableIntGT applies the GT predicate on the "nillable_int" field.
func NillableIntGT(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNillableInt), v))
	},
	)
}

// NillableIntGTE applies the GTE predicate on the "nillable_int" field.
func NillableIntGTE(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNillableInt), v))
	},
	)
}

// NillableIntLT applies the LT predicate on the "nillable_int" field.
func NillableIntLT(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNillableInt), v))
	},
	)
}

// NillableIntLTE applies the LTE predicate on the "nillable_int" field.
func NillableIntLTE(v int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNillableInt), v))
	},
	)
}

// NillableIntIsNil applies the IsNil predicate on the "nillable_int" field.
func NillableIntIsNil() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNillableInt)))
	},
	)
}

// NillableIntNotNil applies the NotNil predicate on the "nillable_int" field.
func NillableIntNotNil() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNillableInt)))
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
