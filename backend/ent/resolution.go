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
	// ResolutionID holds the value of the "Resolution_ID" field.
	ResolutionID int `json:"Resolution_ID,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Resolution) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // Resolution_ID
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
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Resolution_ID", values[0])
	} else if value.Valid {
		r.ResolutionID = int(value.Int64)
	}
	return nil
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
	builder.WriteString(", Resolution_ID=")
	builder.WriteString(fmt.Sprintf("%v", r.ResolutionID))
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