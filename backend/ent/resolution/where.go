// Code generated by entc, DO NOT EDIT.

package resolution

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/pongsakorn-maker/entgo-playground/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ResolutionID applies equality check predicate on the "Resolution_ID" field. It's identical to ResolutionIDEQ.
func ResolutionID(v int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResolutionID), v))
	})
}

// ResolutionIDEQ applies the EQ predicate on the "Resolution_ID" field.
func ResolutionIDEQ(v int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResolutionID), v))
	})
}

// ResolutionIDNEQ applies the NEQ predicate on the "Resolution_ID" field.
func ResolutionIDNEQ(v int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResolutionID), v))
	})
}

// ResolutionIDIn applies the In predicate on the "Resolution_ID" field.
func ResolutionIDIn(vs ...int) predicate.Resolution {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resolution(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResolutionID), v...))
	})
}

// ResolutionIDNotIn applies the NotIn predicate on the "Resolution_ID" field.
func ResolutionIDNotIn(vs ...int) predicate.Resolution {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Resolution(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResolutionID), v...))
	})
}

// ResolutionIDGT applies the GT predicate on the "Resolution_ID" field.
func ResolutionIDGT(v int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResolutionID), v))
	})
}

// ResolutionIDGTE applies the GTE predicate on the "Resolution_ID" field.
func ResolutionIDGTE(v int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResolutionID), v))
	})
}

// ResolutionIDLT applies the LT predicate on the "Resolution_ID" field.
func ResolutionIDLT(v int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResolutionID), v))
	})
}

// ResolutionIDLTE applies the LTE predicate on the "Resolution_ID" field.
func ResolutionIDLTE(v int) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResolutionID), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Resolution) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Resolution) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
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
func Not(p predicate.Resolution) predicate.Resolution {
	return predicate.Resolution(func(s *sql.Selector) {
		p(s.Not())
	})
}