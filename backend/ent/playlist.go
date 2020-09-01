// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/pongsakorn-maker/entgo-playground/ent/playlist"
	"github.com/pongsakorn-maker/entgo-playground/ent/user"
)

// Playlist is the model entity for the Playlist schema.
type Playlist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PlaylistID holds the value of the "playlist_id" field.
	PlaylistID int `json:"playlist_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlaylistQuery when eager-loading is set.
	Edges   PlaylistEdges `json:"edges"`
	user_id *int
}

// PlaylistEdges holds the relations/edges for other nodes in the graph.
type PlaylistEdges struct {
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
func (e PlaylistEdges) PlaylistVideosOrErr() ([]*PlaylistVideo, error) {
	if e.loadedTypes[0] {
		return e.PlaylistVideos, nil
	}
	return nil, &NotLoadedError{edge: "playlist_videos"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlaylistEdges) OwnerOrErr() (*User, error) {
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
func (*Playlist) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // playlist_id
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Playlist) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Playlist fields.
func (pl *Playlist) assignValues(values ...interface{}) error {
	if m, n := len(values), len(playlist.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pl.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field playlist_id", values[0])
	} else if value.Valid {
		pl.PlaylistID = int(value.Int64)
	}
	values = values[1:]
	if len(values) == len(playlist.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_id", value)
		} else if value.Valid {
			pl.user_id = new(int)
			*pl.user_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPlaylistVideos queries the playlist_videos edge of the Playlist.
func (pl *Playlist) QueryPlaylistVideos() *PlaylistVideoQuery {
	return (&PlaylistClient{config: pl.config}).QueryPlaylistVideos(pl)
}

// QueryOwner queries the owner edge of the Playlist.
func (pl *Playlist) QueryOwner() *UserQuery {
	return (&PlaylistClient{config: pl.config}).QueryOwner(pl)
}

// Update returns a builder for updating this Playlist.
// Note that, you need to call Playlist.Unwrap() before calling this method, if this Playlist
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Playlist) Update() *PlaylistUpdateOne {
	return (&PlaylistClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pl *Playlist) Unwrap() *Playlist {
	tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Playlist is not a transactional entity")
	}
	pl.config.driver = tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Playlist) String() string {
	var builder strings.Builder
	builder.WriteString("Playlist(")
	builder.WriteString(fmt.Sprintf("id=%v", pl.ID))
	builder.WriteString(", playlist_id=")
	builder.WriteString(fmt.Sprintf("%v", pl.PlaylistID))
	builder.WriteByte(')')
	return builder.String()
}

// Playlists is a parsable slice of Playlist.
type Playlists []*Playlist

func (pl Playlists) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}
