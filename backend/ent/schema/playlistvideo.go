package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// PlaylistVideo holds the schema definition for the PlaylistVideo entity.
type PlaylistVideo struct {
	ent.Schema
}

// Fields of the PlaylistVideo.
func (PlaylistVideo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("playlistVideo_id"),
	}
}

// Edges of the PlaylistVideo.
func (PlaylistVideo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).Ref("playlist_videos").Unique(),
		edge.From("playlists", Playlist.Type).Ref("playlist_videos").Unique(),
		edge.From("resolution", Resolution.Type).Ref("playlist_videos").Unique(),
	}
}
