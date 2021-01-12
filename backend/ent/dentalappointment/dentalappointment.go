// Code generated by entc, DO NOT EDIT.

package dentalappointment

const (
	// Label holds the string label denoting the dentalappointment type in the database.
	Label = "dentalappointment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppointTime holds the string denoting the appoint_time field in the database.
	FieldAppointTime = "appoint_time"

	// EdgePersonnel holds the string denoting the personnel edge name in mutations.
	EdgePersonnel = "Personnel"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "Patient"
	// EdgeDentaltype holds the string denoting the dentaltype edge name in mutations.
	EdgeDentaltype = "Dentaltype"

	// Table holds the table name of the dentalappointment in the database.
	Table = "dentalappointments"
	// PersonnelTable is the table the holds the Personnel relation/edge.
	PersonnelTable = "dentalappointments"
	// PersonnelInverseTable is the table name for the Personnel entity.
	// It exists in this package in order to avoid circular dependency with the "personnel" package.
	PersonnelInverseTable = "personnels"
	// PersonnelColumn is the table column denoting the Personnel relation/edge.
	PersonnelColumn = "Personnel_id"
	// PatientTable is the table the holds the Patient relation/edge.
	PatientTable = "dentalappointments"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the Patient relation/edge.
	PatientColumn = "Patient_id"
	// DentaltypeTable is the table the holds the Dentaltype relation/edge.
	DentaltypeTable = "dentalappointments"
	// DentaltypeInverseTable is the table name for the Dentaltype entity.
	// It exists in this package in order to avoid circular dependency with the "dentaltype" package.
	DentaltypeInverseTable = "dentaltypes"
	// DentaltypeColumn is the table column denoting the Dentaltype relation/edge.
	DentaltypeColumn = "typename"
)

// Columns holds all SQL columns for dentalappointment fields.
var Columns = []string{
	FieldID,
	FieldAppointTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Dentalappointment type.
var ForeignKeys = []string{
	"typename",
	"Patient_id",
	"Personnel_id",
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
