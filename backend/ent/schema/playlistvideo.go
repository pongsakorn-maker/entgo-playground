package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// PlaylistVideo holds the schema definition for the PlaylistVideo entity.
type PlaylistVideo struct {
	ent.Schema
}

// Fields of the PlaylistVideo.
func (PlaylistVideo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("PlaylistVideo_ID"),
	}
}

// Edges of the PlaylistVideo.
func (PlaylistVideo) Edges() []ent.Edge {
	return nil
}
