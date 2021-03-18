// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/physicaltherapyrecord"
	"github.com/to63/app/ent/physicaltherapyroom"
	"github.com/to63/app/ent/status"
)

// Physicaltherapyrecord is the model entity for the Physicaltherapyrecord schema.
type Physicaltherapyrecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Appointtime holds the value of the "appointtime" field.
	Appointtime time.Time `json:"appointtime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PhysicaltherapyrecordQuery when eager-loading is set.
	Edges                      PhysicaltherapyrecordEdges `json:"edges"`
	_Patient_id                *int
	_Personnel_id              *int
	physicaltherapyroomname_id *int
	statusname_id              *int
}

// PhysicaltherapyrecordEdges holds the relations/edges for other nodes in the graph.
type PhysicaltherapyrecordEdges struct {
	// Personnel holds the value of the personnel edge.
	Personnel *Personnel
	// Patient holds the value of the patient edge.
	Patient *Patient
	// Physicaltherapyroom holds the value of the physicaltherapyroom edge.
	Physicaltherapyroom *Physicaltherapyroom
	// Status holds the value of the status edge.
	Status *Status
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PersonnelOrErr returns the Personnel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhysicaltherapyrecordEdges) PersonnelOrErr() (*Personnel, error) {
	if e.loadedTypes[0] {
		if e.Personnel == nil {
			// The edge personnel was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: personnel.Label}
		}
		return e.Personnel, nil
	}
	return nil, &NotLoadedError{edge: "personnel"}
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhysicaltherapyrecordEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[1] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// PhysicaltherapyroomOrErr returns the Physicaltherapyroom value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhysicaltherapyrecordEdges) PhysicaltherapyroomOrErr() (*Physicaltherapyroom, error) {
	if e.loadedTypes[2] {
		if e.Physicaltherapyroom == nil {
			// The edge physicaltherapyroom was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: physicaltherapyroom.Label}
		}
		return e.Physicaltherapyroom, nil
	}
	return nil, &NotLoadedError{edge: "physicaltherapyroom"}
}

// StatusOrErr returns the Status value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhysicaltherapyrecordEdges) StatusOrErr() (*Status, error) {
	if e.loadedTypes[3] {
		if e.Status == nil {
			// The edge status was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: status.Label}
		}
		return e.Status, nil
	}
	return nil, &NotLoadedError{edge: "status"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Physicaltherapyrecord) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case physicaltherapyrecord.FieldID, physicaltherapyrecord.FieldPrice:
			values[i] = &sql.NullInt64{}
		case physicaltherapyrecord.FieldNote:
			values[i] = &sql.NullString{}
		case physicaltherapyrecord.FieldAppointtime:
			values[i] = &sql.NullTime{}
		case physicaltherapyrecord.ForeignKeys[0]: // _Patient_id
			values[i] = &sql.NullInt64{}
		case physicaltherapyrecord.ForeignKeys[1]: // _Personnel_id
			values[i] = &sql.NullInt64{}
		case physicaltherapyrecord.ForeignKeys[2]: // physicaltherapyroomname_id
			values[i] = &sql.NullInt64{}
		case physicaltherapyrecord.ForeignKeys[3]: // statusname_id
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Physicaltherapyrecord", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Physicaltherapyrecord fields.
func (ph *Physicaltherapyrecord) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case physicaltherapyrecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ph.ID = int(value.Int64)
		case physicaltherapyrecord.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				ph.Price = int(value.Int64)
			}
		case physicaltherapyrecord.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				ph.Note = value.String
			}
		case physicaltherapyrecord.FieldAppointtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field appointtime", values[i])
			} else if value.Valid {
				ph.Appointtime = value.Time
			}
		case physicaltherapyrecord.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field _Patient_id", value)
			} else if value.Valid {
				ph._Patient_id = new(int)
				*ph._Patient_id = int(value.Int64)
			}
		case physicaltherapyrecord.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field _Personnel_id", value)
			} else if value.Valid {
				ph._Personnel_id = new(int)
				*ph._Personnel_id = int(value.Int64)
			}
		case physicaltherapyrecord.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field physicaltherapyroomname_id", value)
			} else if value.Valid {
				ph.physicaltherapyroomname_id = new(int)
				*ph.physicaltherapyroomname_id = int(value.Int64)
			}
		case physicaltherapyrecord.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field statusname_id", value)
			} else if value.Valid {
				ph.statusname_id = new(int)
				*ph.statusname_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPersonnel queries the "personnel" edge of the Physicaltherapyrecord entity.
func (ph *Physicaltherapyrecord) QueryPersonnel() *PersonnelQuery {
	return (&PhysicaltherapyrecordClient{config: ph.config}).QueryPersonnel(ph)
}

// QueryPatient queries the "patient" edge of the Physicaltherapyrecord entity.
func (ph *Physicaltherapyrecord) QueryPatient() *PatientQuery {
	return (&PhysicaltherapyrecordClient{config: ph.config}).QueryPatient(ph)
}

// QueryPhysicaltherapyroom queries the "physicaltherapyroom" edge of the Physicaltherapyrecord entity.
func (ph *Physicaltherapyrecord) QueryPhysicaltherapyroom() *PhysicaltherapyroomQuery {
	return (&PhysicaltherapyrecordClient{config: ph.config}).QueryPhysicaltherapyroom(ph)
}

// QueryStatus queries the "status" edge of the Physicaltherapyrecord entity.
func (ph *Physicaltherapyrecord) QueryStatus() *StatusQuery {
	return (&PhysicaltherapyrecordClient{config: ph.config}).QueryStatus(ph)
}

// Update returns a builder for updating this Physicaltherapyrecord.
// Note that you need to call Physicaltherapyrecord.Unwrap() before calling this method if this Physicaltherapyrecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (ph *Physicaltherapyrecord) Update() *PhysicaltherapyrecordUpdateOne {
	return (&PhysicaltherapyrecordClient{config: ph.config}).UpdateOne(ph)
}

// Unwrap unwraps the Physicaltherapyrecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ph *Physicaltherapyrecord) Unwrap() *Physicaltherapyrecord {
	tx, ok := ph.config.driver.(*txDriver)
	if !ok {
		panic("ent: Physicaltherapyrecord is not a transactional entity")
	}
	ph.config.driver = tx.drv
	return ph
}

// String implements the fmt.Stringer.
func (ph *Physicaltherapyrecord) String() string {
	var builder strings.Builder
	builder.WriteString("Physicaltherapyrecord(")
	builder.WriteString(fmt.Sprintf("id=%v", ph.ID))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", ph.Price))
	builder.WriteString(", note=")
	builder.WriteString(ph.Note)
	builder.WriteString(", appointtime=")
	builder.WriteString(ph.Appointtime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Physicaltherapyrecords is a parsable slice of Physicaltherapyrecord.
type Physicaltherapyrecords []*Physicaltherapyrecord

func (ph Physicaltherapyrecords) config(cfg config) {
	for _i := range ph {
		ph[_i].config = cfg
	}
}
