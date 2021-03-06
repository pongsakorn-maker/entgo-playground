// Code generated by entc, DO NOT EDIT.

package playlistvideo

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/pongsakorn-maker/entgo-playground/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
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
func IDGT(id int) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HasVideo applies the HasEdge predicate on the "video" edge.
func HasVideo() predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VideoTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideoWith applies the HasEdge predicate on the "video" edge with a given conditions (other predicates).
func HasVideoWith(preds ...predicate.Video) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VideoInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlaylists applies the HasEdge predicate on the "playlists" edge.
func HasPlaylists() predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlaylistsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlaylistsTable, PlaylistsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaylistsWith applies the HasEdge predicate on the "playlists" edge with a given conditions (other predicates).
func HasPlaylistsWith(preds ...predicate.Playlist) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlaylistsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlaylistsTable, PlaylistsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResolution applies the HasEdge predicate on the "resolution" edge.
func HasResolution() predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResolutionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResolutionTable, ResolutionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResolutionWith applies the HasEdge predicate on the "resolution" edge with a given conditions (other predicates).
func HasResolutionWith(preds ...predicate.Resolution) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResolutionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResolutionTable, ResolutionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.PlaylistVideo) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.PlaylistVideo) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
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
func Not(p predicate.PlaylistVideo) predicate.PlaylistVideo {
	return predicate.PlaylistVideo(func(s *sql.Selector) {
		p(s.Not())
	})
}
