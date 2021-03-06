// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/to63/app/ent/checksymptom"
	"github.com/to63/app/ent/disease"
	"github.com/to63/app/ent/doctorordersheet"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
)

// Checksymptom is the model entity for the Checksymptom schema.
type Checksymptom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Identitycard holds the value of the "Identitycard" field.
	Identitycard string `json:"Identitycard,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChecksymptomQuery when eager-loading is set.
	Edges                ChecksymptomEdges `json:"edges"`
	disease_id           *int
	_Doctorordersheet_id *int
	_Patient_id          *int
	_Personnel_id        *int
}

// ChecksymptomEdges holds the relations/edges for other nodes in the graph.
type ChecksymptomEdges struct {
	// Patient holds the value of the patient edge.
	Patient *Patient
	// Personnel holds the value of the personnel edge.
	Personnel *Personnel
	// Doctorordersheet holds the value of the doctorordersheet edge.
	Doctorordersheet *Doctorordersheet
	// Disease holds the value of the disease edge.
	Disease *Disease
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChecksymptomEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[0] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// PersonnelOrErr returns the Personnel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChecksymptomEdges) PersonnelOrErr() (*Personnel, error) {
	if e.loadedTypes[1] {
		if e.Personnel == nil {
			// The edge personnel was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: personnel.Label}
		}
		return e.Personnel, nil
	}
	return nil, &NotLoadedError{edge: "personnel"}
}

// DoctorordersheetOrErr returns the Doctorordersheet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChecksymptomEdges) DoctorordersheetOrErr() (*Doctorordersheet, error) {
	if e.loadedTypes[2] {
		if e.Doctorordersheet == nil {
			// The edge doctorordersheet was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doctorordersheet.Label}
		}
		return e.Doctorordersheet, nil
	}
	return nil, &NotLoadedError{edge: "doctorordersheet"}
}

// DiseaseOrErr returns the Disease value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChecksymptomEdges) DiseaseOrErr() (*Disease, error) {
	if e.loadedTypes[3] {
		if e.Disease == nil {
			// The edge disease was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: disease.Label}
		}
		return e.Disease, nil
	}
	return nil, &NotLoadedError{edge: "disease"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Checksymptom) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case checksymptom.FieldID:
			values[i] = &sql.NullInt64{}
		case checksymptom.FieldNote, checksymptom.FieldIdentitycard, checksymptom.FieldPhone:
			values[i] = &sql.NullString{}
		case checksymptom.FieldDate:
			values[i] = &sql.NullTime{}
		case checksymptom.ForeignKeys[0]: // disease_id
			values[i] = &sql.NullInt64{}
		case checksymptom.ForeignKeys[1]: // _Doctorordersheet_id
			values[i] = &sql.NullInt64{}
		case checksymptom.ForeignKeys[2]: // _Patient_id
			values[i] = &sql.NullInt64{}
		case checksymptom.ForeignKeys[3]: // _Personnel_id
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Checksymptom", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Checksymptom fields.
func (c *Checksymptom) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case checksymptom.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case checksymptom.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				c.Date = value.Time
			}
		case checksymptom.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				c.Note = value.String
			}
		case checksymptom.FieldIdentitycard:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Identitycard", values[i])
			} else if value.Valid {
				c.Identitycard = value.String
			}
		case checksymptom.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case checksymptom.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field disease_id", value)
			} else if value.Valid {
				c.disease_id = new(int)
				*c.disease_id = int(value.Int64)
			}
		case checksymptom.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field _Doctorordersheet_id", value)
			} else if value.Valid {
				c._Doctorordersheet_id = new(int)
				*c._Doctorordersheet_id = int(value.Int64)
			}
		case checksymptom.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field _Patient_id", value)
			} else if value.Valid {
				c._Patient_id = new(int)
				*c._Patient_id = int(value.Int64)
			}
		case checksymptom.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field _Personnel_id", value)
			} else if value.Valid {
				c._Personnel_id = new(int)
				*c._Personnel_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPatient queries the "patient" edge of the Checksymptom entity.
func (c *Checksymptom) QueryPatient() *PatientQuery {
	return (&ChecksymptomClient{config: c.config}).QueryPatient(c)
}

// QueryPersonnel queries the "personnel" edge of the Checksymptom entity.
func (c *Checksymptom) QueryPersonnel() *PersonnelQuery {
	return (&ChecksymptomClient{config: c.config}).QueryPersonnel(c)
}

// QueryDoctorordersheet queries the "doctorordersheet" edge of the Checksymptom entity.
func (c *Checksymptom) QueryDoctorordersheet() *DoctorordersheetQuery {
	return (&ChecksymptomClient{config: c.config}).QueryDoctorordersheet(c)
}

// QueryDisease queries the "disease" edge of the Checksymptom entity.
func (c *Checksymptom) QueryDisease() *DiseaseQuery {
	return (&ChecksymptomClient{config: c.config}).QueryDisease(c)
}

// Update returns a builder for updating this Checksymptom.
// Note that you need to call Checksymptom.Unwrap() before calling this method if this Checksymptom
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Checksymptom) Update() *ChecksymptomUpdateOne {
	return (&ChecksymptomClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Checksymptom entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Checksymptom) Unwrap() *Checksymptom {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Checksymptom is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Checksymptom) String() string {
	var builder strings.Builder
	builder.WriteString("Checksymptom(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", date=")
	builder.WriteString(c.Date.Format(time.ANSIC))
	builder.WriteString(", note=")
	builder.WriteString(c.Note)
	builder.WriteString(", Identitycard=")
	builder.WriteString(c.Identitycard)
	builder.WriteString(", phone=")
	builder.WriteString(c.Phone)
	builder.WriteByte(')')
	return builder.String()
}

// Checksymptoms is a parsable slice of Checksymptom.
type Checksymptoms []*Checksymptom

func (c Checksymptoms) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
