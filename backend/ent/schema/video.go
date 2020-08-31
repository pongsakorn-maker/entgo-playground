package schema

import "github.com/facebook/ent"

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return nil
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return nil
}
