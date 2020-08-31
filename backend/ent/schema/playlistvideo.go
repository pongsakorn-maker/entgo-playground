package schema

import "github.com/facebook/ent"

// PlaylistVideo holds the schema definition for the PlaylistVideo entity.
type PlaylistVideo struct {
	ent.Schema
}

// Fields of the PlaylistVideo.
func (PlaylistVideo) Fields() []ent.Field {
	return nil
}

// Edges of the PlaylistVideo.
func (PlaylistVideo) Edges() []ent.Edge {
	return nil
}
