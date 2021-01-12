// Code generated by entc, DO NOT EDIT.

package doctorordersheet

import (
	"time"
)

const (
	// Label holds the string label denoting the doctorordersheet type in the database.
	Label = "doctor_order_sheet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"

	// EdgeChecksymptoms holds the string denoting the checksymptoms edge name in mutations.
	EdgeChecksymptoms = "Checksymptoms"

	// Table holds the table name of the doctorordersheet in the database.
	Table = "doctor_order_sheets"
	// ChecksymptomsTable is the table the holds the Checksymptoms relation/edge.
	ChecksymptomsTable = "checksymptoms"
	// ChecksymptomsInverseTable is the table name for the Checksymptoms entity.
	// It exists in this package in order to avoid circular dependency with the "checksymptoms" package.
	ChecksymptomsInverseTable = "checksymptoms"
	// ChecksymptomsColumn is the table column denoting the Checksymptoms relation/edge.
	ChecksymptomsColumn = "DoctorOrderSheet_id"
)

// Columns holds all SQL columns for doctorordersheet fields.
var Columns = []string{
	FieldID,
	FieldDate,
	FieldTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDate holds the default value on creation for the "date" field.
	DefaultDate func() time.Time
	// TimeValidator is a validator for the "time" field. It is called by the builders before save.
	TimeValidator func(string) error
)
