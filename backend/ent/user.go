// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/pongsakorn-maker/entgo-playground/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "User_ID" field.
	UserID int `json:"User_ID,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Playlists holds the value of the playlists edge.
	Playlists []*Playlist
	// Videos holds the value of the videos edge.
	Videos []*Video
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PlaylistsOrErr returns the Playlists value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PlaylistsOrErr() ([]*Playlist, error) {
	if e.loadedTypes[0] {
		return e.Playlists, nil
	}
	return nil, &NotLoadedError{edge: "playlists"}
}

// VideosOrErr returns the Videos value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) VideosOrErr() ([]*Video, error) {
	if e.loadedTypes[1] {
		return e.Videos, nil
	}
	return nil, &NotLoadedError{edge: "videos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // User_ID
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field User_ID", values[0])
	} else if value.Valid {
		u.UserID = int(value.Int64)
	}
	return nil
}

// QueryPlaylists queries the playlists edge of the User.
func (u *User) QueryPlaylists() *PlaylistQuery {
	return (&UserClient{config: u.config}).QueryPlaylists(u)
}

// QueryVideos queries the videos edge of the User.
func (u *User) QueryVideos() *VideoQuery {
	return (&UserClient{config: u.config}).QueryVideos(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", User_ID=")
	builder.WriteString(fmt.Sprintf("%v", u.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
