// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/pongsakorn-maker/entgo-playground/ent/user"
	"github.com/pongsakorn-maker/entgo-playground/ent/video"
)

// Video is the model entity for the Video schema.
type Video struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VideoTitle holds the value of the "video_title" field.
	VideoTitle string `json:"video_title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VideoQuery when eager-loading is set.
	Edges   VideoEdges `json:"edges"`
	user_id *int
}

// VideoEdges holds the relations/edges for other nodes in the graph.
type VideoEdges struct {
	// PlaylistVideos holds the value of the playlist_videos edge.
	PlaylistVideos []*PlaylistVideo
	// Owner holds the value of the owner edge.
	Owner *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PlaylistVideosOrErr returns the PlaylistVideos value or an error if the edge
// was not loaded in eager-loading.
func (e VideoEdges) PlaylistVideosOrErr() ([]*PlaylistVideo, error) {
	if e.loadedTypes[0] {
		return e.PlaylistVideos, nil
	}
	return nil, &NotLoadedError{edge: "playlist_videos"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VideoEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Video) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // video_title
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Video) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Video fields.
func (v *Video) assignValues(values ...interface{}) error {
	if m, n := len(values), len(video.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	v.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field video_title", values[0])
	} else if value.Valid {
		v.VideoTitle = value.String
	}
	values = values[1:]
	if len(values) == len(video.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_id", value)
		} else if value.Valid {
			v.user_id = new(int)
			*v.user_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPlaylistVideos queries the playlist_videos edge of the Video.
func (v *Video) QueryPlaylistVideos() *PlaylistVideoQuery {
	return (&VideoClient{config: v.config}).QueryPlaylistVideos(v)
}

// QueryOwner queries the owner edge of the Video.
func (v *Video) QueryOwner() *UserQuery {
	return (&VideoClient{config: v.config}).QueryOwner(v)
}

// Update returns a builder for updating this Video.
// Note that, you need to call Video.Unwrap() before calling this method, if this Video
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Video) Update() *VideoUpdateOne {
	return (&VideoClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (v *Video) Unwrap() *Video {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Video is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Video) String() string {
	var builder strings.Builder
	builder.WriteString("Video(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", video_title=")
	builder.WriteString(v.VideoTitle)
	builder.WriteByte(')')
	return builder.String()
}

// Videos is a parsable slice of Video.
type Videos []*Video

func (v Videos) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
