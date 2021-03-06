// Code generated by entc, DO NOT EDIT.

package video

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/pongsakorn-maker/entgo-playground/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// VideoTitle applies equality check predicate on the "video_title" field. It's identical to VideoTitleEQ.
func VideoTitle(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleEQ applies the EQ predicate on the "video_title" field.
func VideoTitleEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleNEQ applies the NEQ predicate on the "video_title" field.
func VideoTitleNEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleIn applies the In predicate on the "video_title" field.
func VideoTitleIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVideoTitle), v...))
	})
}

// VideoTitleNotIn applies the NotIn predicate on the "video_title" field.
func VideoTitleNotIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVideoTitle), v...))
	})
}

// VideoTitleGT applies the GT predicate on the "video_title" field.
func VideoTitleGT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleGTE applies the GTE predicate on the "video_title" field.
func VideoTitleGTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleLT applies the LT predicate on the "video_title" field.
func VideoTitleLT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleLTE applies the LTE predicate on the "video_title" field.
func VideoTitleLTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleContains applies the Contains predicate on the "video_title" field.
func VideoTitleContains(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleHasPrefix applies the HasPrefix predicate on the "video_title" field.
func VideoTitleHasPrefix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleHasSuffix applies the HasSuffix predicate on the "video_title" field.
func VideoTitleHasSuffix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleEqualFold applies the EqualFold predicate on the "video_title" field.
func VideoTitleEqualFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVideoTitle), v))
	})
}

// VideoTitleContainsFold applies the ContainsFold predicate on the "video_title" field.
func VideoTitleContainsFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVideoTitle), v))
	})
}

// HasPlaylistVideos applies the HasEdge predicate on the "playlist_videos" edge.
func HasPlaylistVideos() predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlaylistVideosTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PlaylistVideosTable, PlaylistVideosColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaylistVideosWith applies the HasEdge predicate on the "playlist_videos" edge with a given conditions (other predicates).
func HasPlaylistVideosWith(preds ...predicate.PlaylistVideo) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlaylistVideosInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PlaylistVideosTable, PlaylistVideosColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Video) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Video) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
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
func Not(p predicate.Video) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		p(s.Not())
	})
}
