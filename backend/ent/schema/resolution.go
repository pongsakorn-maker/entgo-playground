package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Resolution holds the schema definition for the Resolution entity.
type Resolution struct {
	ent.Schema
}

// Fields of the Resolution.
func (Resolution) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Resolution_ID"),
	}
}

// Edges of the Resolution.
func (Resolution) Edges() []ent.Edge {
	return nil
}
