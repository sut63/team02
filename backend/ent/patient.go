// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/to63/app/ent/patient"
)

// Patient is the model entity for the Patient schema.
type Patient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday string `json:"birthday,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientQuery when eager-loading is set.
	Edges PatientEdges `json:"edges"`
}

// PatientEdges holds the relations/edges for other nodes in the graph.
type PatientEdges struct {
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
func (e PatientEdges) PhysicaltherapyrecordOrErr() ([]*Physicaltherapyrecord, error) {
	if e.loadedTypes[0] {
		return e.Physicaltherapyrecord, nil
	}
	return nil, &NotLoadedError{edge: "physicaltherapyrecord"}
}

// BonediseaseOrErr returns the Bonedisease value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) BonediseaseOrErr() ([]*Bonedisease, error) {
	if e.loadedTypes[1] {
		return e.Bonedisease, nil
	}
	return nil, &NotLoadedError{edge: "Bonedisease"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patient) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case patient.FieldID:
			values[i] = &sql.NullInt64{}
		case patient.FieldName, patient.FieldBirthday, patient.FieldGender:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Patient", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patient fields.
func (pa *Patient) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case patient.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case patient.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pa.Name = value.String
			}
		case patient.FieldBirthday:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				pa.Birthday = value.String
			}
		case patient.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				pa.Gender = value.String
			}
		}
	}
	return nil
}

// QueryPhysicaltherapyrecord queries the "physicaltherapyrecord" edge of the Patient entity.
func (pa *Patient) QueryPhysicaltherapyrecord() *PhysicaltherapyrecordQuery {
	return (&PatientClient{config: pa.config}).QueryPhysicaltherapyrecord(pa)
}

// QueryBonedisease queries the "Bonedisease" edge of the Patient entity.
func (pa *Patient) QueryBonedisease() *BonediseaseQuery {
	return (&PatientClient{config: pa.config}).QueryBonedisease(pa)
}

// Update returns a builder for updating this Patient.
// Note that you need to call Patient.Unwrap() before calling this method if this Patient
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patient) Update() *PatientUpdateOne {
	return (&PatientClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Patient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Patient) Unwrap() *Patient {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patient is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patient) String() string {
	var builder strings.Builder
	builder.WriteString("Patient(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", name=")
	builder.WriteString(pa.Name)
	builder.WriteString(", birthday=")
	builder.WriteString(pa.Birthday)
	builder.WriteString(", gender=")
	builder.WriteString(pa.Gender)
	builder.WriteByte(')')
	return builder.String()
}

// Patients is a parsable slice of Patient.
type Patients []*Patient

func (pa Patients) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
