package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Video_ID"),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return nil
}
