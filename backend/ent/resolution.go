// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/pongsakorn-maker/entgo-playground/ent/resolution"
)

// Resolution is the model entity for the Resolution schema.
type Resolution struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Resolution holds the value of the "resolution" field.
	Resolution string `json:"resolution,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResolutionQuery when eager-loading is set.
	Edges ResolutionEdges `json:"edges"`
}

// ResolutionEdges holds the relations/edges for other nodes in the graph.
type ResolutionEdges struct {
	// PlaylistVideos holds the value of the playlist_videos edge.
	PlaylistVideos []*PlaylistVideo
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PlaylistVideosOrErr returns the PlaylistVideos value or an error if the edge
// was not loaded in eager-loading.
func (e ResolutionEdges) PlaylistVideosOrErr() ([]*PlaylistVideo, error) {
	if e.loadedTypes[0] {
		return e.PlaylistVideos, nil
	}
	return nil, &NotLoadedError{edge: "playlist_videos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Resolution) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // resolution
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Resolution fields.
func (r *Resolution) assignValues(values ...interface{}) error {
	if m, n := len(values), len(resolution.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field resolution", values[0])
	} else if value.Valid {
		r.Resolution = value.String
	}
	return nil
}

// QueryPlaylistVideos queries the playlist_videos edge of the Resolution.
func (r *Resolution) QueryPlaylistVideos() *PlaylistVideoQuery {
	return (&ResolutionClient{config: r.config}).QueryPlaylistVideos(r)
}

// Update returns a builder for updating this Resolution.
// Note that, you need to call Resolution.Unwrap() before calling this method, if this Resolution
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Resolution) Update() *ResolutionUpdateOne {
	return (&ResolutionClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Resolution) Unwrap() *Resolution {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Resolution is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Resolution) String() string {
	var builder strings.Builder
	builder.WriteString("Resolution(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", resolution=")
	builder.WriteString(r.Resolution)
	builder.WriteByte(')')
	return builder.String()
}

// Resolutions is a parsable slice of Resolution.
type Resolutions []*Resolution

func (r Resolutions) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
