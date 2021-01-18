// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/to63/app/ent/risks"
)

// Risks is the model entity for the Risks schema.
type Risks struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Risks holds the value of the "Risks" field.
	Risks string `json:"Risks,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RisksQuery when eager-loading is set.
	Edges RisksEdges `json:"edges"`
}

// RisksEdges holds the relations/edges for other nodes in the graph.
type RisksEdges struct {
	// Antenatalinformation holds the value of the Antenatalinformation edge.
	Antenatalinformation []*Antenatalinformation
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AntenatalinformationOrErr returns the Antenatalinformation value or an error if the edge
// was not loaded in eager-loading.
func (e RisksEdges) AntenatalinformationOrErr() ([]*Antenatalinformation, error) {
	if e.loadedTypes[0] {
		return e.Antenatalinformation, nil
	}
	return nil, &NotLoadedError{edge: "Antenatalinformation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Risks) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case risks.FieldID:
			values[i] = &sql.NullInt64{}
		case risks.FieldRisks:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Risks", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Risks fields.
func (r *Risks) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case risks.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case risks.FieldRisks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Risks", values[i])
			} else if value.Valid {
				r.Risks = value.String
			}
		}
	}
	return nil
}

// QueryAntenatalinformation queries the "Antenatalinformation" edge of the Risks entity.
func (r *Risks) QueryAntenatalinformation() *AntenatalinformationQuery {
	return (&RisksClient{config: r.config}).QueryAntenatalinformation(r)
}

// Update returns a builder for updating this Risks.
// Note that you need to call Risks.Unwrap() before calling this method if this Risks
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Risks) Update() *RisksUpdateOne {
	return (&RisksClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Risks entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Risks) Unwrap() *Risks {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Risks is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Risks) String() string {
	var builder strings.Builder
	builder.WriteString("Risks(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", Risks=")
	builder.WriteString(r.Risks)
	builder.WriteByte(')')
	return builder.String()
}

// RisksSlice is a parsable slice of Risks.
type RisksSlice []*Risks

func (r RisksSlice) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
