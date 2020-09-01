package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("password"),
		field.Int("age").Positive(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("playlists", Playlist.Type).StorageKey(edge.Column("user_id")),
		edge.To("videos", Video.Type).StorageKey(edge.Column("user_id")),
	}
}
