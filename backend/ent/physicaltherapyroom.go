// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/to63/app/ent/physicaltherapyroom"
)

// Physicaltherapyroom is the model entity for the Physicaltherapyroom schema.
type Physicaltherapyroom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Physicaltherapyroomname holds the value of the "physicaltherapyroomname" field.
	Physicaltherapyroomname string `json:"physicaltherapyroomname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PhysicaltherapyroomQuery when eager-loading is set.
	Edges PhysicaltherapyroomEdges `json:"edges"`
}

// PhysicaltherapyroomEdges holds the relations/edges for other nodes in the graph.
type PhysicaltherapyroomEdges struct {
	// Physicaltherapyrecord holds the value of the physicaltherapyrecord edge.
	Physicaltherapyrecord []*Physicaltherapyrecord
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PhysicaltherapyrecordOrErr returns the Physicaltherapyrecord value or an error if the edge
// was not loaded in eager-loading.
func (e PhysicaltherapyroomEdges) PhysicaltherapyrecordOrErr() ([]*Physicaltherapyrecord, error) {
	if e.loadedTypes[0] {
		return e.Physicaltherapyrecord, nil
	}
	return nil, &NotLoadedError{edge: "physicaltherapyrecord"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Physicaltherapyroom) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case physicaltherapyroom.FieldID:
			values[i] = &sql.NullInt64{}
		case physicaltherapyroom.FieldPhysicaltherapyroomname:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Physicaltherapyroom", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Physicaltherapyroom fields.
func (ph *Physicaltherapyroom) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case physicaltherapyroom.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ph.ID = int(value.Int64)
		case physicaltherapyroom.FieldPhysicaltherapyroomname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field physicaltherapyroomname", values[i])
			} else if value.Valid {
				ph.Physicaltherapyroomname = value.String
			}
		}
	}
	return nil
}

// QueryPhysicaltherapyrecord queries the "physicaltherapyrecord" edge of the Physicaltherapyroom entity.
func (ph *Physicaltherapyroom) QueryPhysicaltherapyrecord() *PhysicaltherapyrecordQuery {
	return (&PhysicaltherapyroomClient{config: ph.config}).QueryPhysicaltherapyrecord(ph)
}

// Update returns a builder for updating this Physicaltherapyroom.
// Note that you need to call Physicaltherapyroom.Unwrap() before calling this method if this Physicaltherapyroom
// was returned from a transaction, and the transaction was committed or rolled back.
func (ph *Physicaltherapyroom) Update() *PhysicaltherapyroomUpdateOne {
	return (&PhysicaltherapyroomClient{config: ph.config}).UpdateOne(ph)
}

// Unwrap unwraps the Physicaltherapyroom entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ph *Physicaltherapyroom) Unwrap() *Physicaltherapyroom {
	tx, ok := ph.config.driver.(*txDriver)
	if !ok {
		panic("ent: Physicaltherapyroom is not a transactional entity")
	}
	ph.config.driver = tx.drv
	return ph
}

// String implements the fmt.Stringer.
func (ph *Physicaltherapyroom) String() string {
	var builder strings.Builder
	builder.WriteString("Physicaltherapyroom(")
	builder.WriteString(fmt.Sprintf("id=%v", ph.ID))
	builder.WriteString(", physicaltherapyroomname=")
	builder.WriteString(ph.Physicaltherapyroomname)
	builder.WriteByte(')')
	return builder.String()
}

// Physicaltherapyrooms is a parsable slice of Physicaltherapyroom.
type Physicaltherapyrooms []*Physicaltherapyroom

func (ph Physicaltherapyrooms) config(cfg config) {
	for _i := range ph {
		ph[_i].config = cfg
	}
}
