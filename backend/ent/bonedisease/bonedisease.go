// Code generated by entc, DO NOT EDIT.

package bonedisease

import (
	"time"
)

const (
	// Label holds the string label denoting the bonedisease type in the database.
	Label = "bonedisease"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddedTime holds the string denoting the addedtime field in the database.
	FieldAddedTime = "added_time"
	// FieldAdvice holds the string denoting the advice field in the database.
	FieldAdvice = "advice"

	// EdgeRemedy holds the string denoting the remedy edge name in mutations.
	EdgeRemedy = "remedy"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgePersonnel holds the string denoting the personnel edge name in mutations.
	EdgePersonnel = "personnel"

	// Table holds the table name of the bonedisease in the database.
	Table = "bonediseases"
	// RemedyTable is the table the holds the remedy relation/edge.
	RemedyTable = "bonediseases"
	// RemedyInverseTable is the table name for the Remedy entity.
	// It exists in this package in order to avoid circular dependency with the "remedy" package.
	RemedyInverseTable = "remedies"
	// RemedyColumn is the table column denoting the remedy relation/edge.
	RemedyColumn = "remedy_id"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "bonediseases"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "Patient_id"
	// PersonnelTable is the table the holds the personnel relation/edge.
	PersonnelTable = "bonediseases"
	// PersonnelInverseTable is the table name for the Personnel entity.
	// It exists in this package in order to avoid circular dependency with the "personnel" package.
	PersonnelInverseTable = "personnels"
	// PersonnelColumn is the table column denoting the personnel relation/edge.
	PersonnelColumn = "Personel_id"
)

// Columns holds all SQL columns for bonedisease fields.
var Columns = []string{
	FieldID,
	FieldAddedTime,
	FieldAdvice,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Bonedisease type.
var ForeignKeys = []string{
	"Patient_id",
	"Personel_id",
	"remedy_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAddedTime holds the default value on creation for the "addedTime" field.
	DefaultAddedTime func() time.Time
	// AdviceValidator is a validator for the "advice" field. It is called by the builders before save.
	AdviceValidator func(string) error
)
