// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/to63/app/ent/personnel"
)

// Personnel is the model entity for the Personnel schema.
type Personnel struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Department holds the value of the "department" field.
	Department string `json:"department,omitempty"`
	// User holds the value of the "user" field.
	User string `json:"user,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonnelQuery when eager-loading is set.
	Edges PersonnelEdges `json:"edges"`
}

// PersonnelEdges holds the relations/edges for other nodes in the graph.
type PersonnelEdges struct {
	// Physicaltherapyrecord holds the value of the physicaltherapyrecord edge.
	Physicaltherapyrecord []*Physicaltherapyrecord
	// Bonedisease holds the value of the Bonedisease edge.
	Bonedisease []*Bonedisease
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PhysicaltherapyrecordOrErr returns the Physicaltherapyrecord value or an error if the edge
// was not loaded in eager-loading.
func (e PersonnelEdges) PhysicaltherapyrecordOrErr() ([]*Physicaltherapyrecord, error) {
	if e.loadedTypes[0] {
		return e.Physicaltherapyrecord, nil
	}
	return nil, &NotLoadedError{edge: "physicaltherapyrecord"}
}

// BonediseaseOrErr returns the Bonedisease value or an error if the edge
// was not loaded in eager-loading.
func (e PersonnelEdges) BonediseaseOrErr() ([]*Bonedisease, error) {
	if e.loadedTypes[1] {
		return e.Bonedisease, nil
	}
	return nil, &NotLoadedError{edge: "Bonedisease"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Personnel) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case personnel.FieldID:
			values[i] = &sql.NullInt64{}
		case personnel.FieldName, personnel.FieldDepartment, personnel.FieldUser, personnel.FieldPassword:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Personnel", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Personnel fields.
func (pe *Personnel) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case personnel.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case personnel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case personnel.FieldDepartment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field department", values[i])
			} else if value.Valid {
				pe.Department = value.String
			}
		case personnel.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				pe.User = value.String
			}
		case personnel.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				pe.Password = value.String
			}
		}
	}
	return nil
}

// QueryPhysicaltherapyrecord queries the "physicaltherapyrecord" edge of the Personnel entity.
func (pe *Personnel) QueryPhysicaltherapyrecord() *PhysicaltherapyrecordQuery {
	return (&PersonnelClient{config: pe.config}).QueryPhysicaltherapyrecord(pe)
}

// QueryBonedisease queries the "Bonedisease" edge of the Personnel entity.
func (pe *Personnel) QueryBonedisease() *BonediseaseQuery {
	return (&PersonnelClient{config: pe.config}).QueryBonedisease(pe)
}

// Update returns a builder for updating this Personnel.
// Note that you need to call Personnel.Unwrap() before calling this method if this Personnel
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Personnel) Update() *PersonnelUpdateOne {
	return (&PersonnelClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Personnel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Personnel) Unwrap() *Personnel {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Personnel is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Personnel) String() string {
	var builder strings.Builder
	builder.WriteString("Personnel(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", department=")
	builder.WriteString(pe.Department)
	builder.WriteString(", user=")
	builder.WriteString(pe.User)
	builder.WriteString(", password=")
	builder.WriteString(pe.Password)
	builder.WriteByte(')')
	return builder.String()
}

// Personnels is a parsable slice of Personnel.
type Personnels []*Personnel

func (pe Personnels) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
