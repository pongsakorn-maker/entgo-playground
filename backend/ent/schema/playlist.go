package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Playlist holds the schema definition for the Playlist entity.
type Playlist struct {
	ent.Schema
}

// Fields of the Playlist.
func (Playlist) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Playlist_ID"),
	}
}

// Edges of the Playlist.
func (Playlist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("playlist_videos", PlaylistVideo.Type).StorageKey(edge.Column("playlist_id")),
		edge.From("owner", User.Type).Ref("playlists").Unique(),
	}
}
