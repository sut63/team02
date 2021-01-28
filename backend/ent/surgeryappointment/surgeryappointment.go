// Code generated by entc, DO NOT EDIT.

package surgeryappointment

const (
	// Label holds the string label denoting the surgeryappointment type in the database.
	Label = "surgeryappointment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppointTime holds the string denoting the appoint_time field in the database.
	FieldAppointTime = "appoint_time"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"

	// EdgePersonnel holds the string denoting the personnel edge name in mutations.
	EdgePersonnel = "Personnel"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "Patient"
	// EdgeSurgerytype holds the string denoting the surgerytype edge name in mutations.
	EdgeSurgerytype = "Surgerytype"

	// Table holds the table name of the surgeryappointment in the database.
	Table = "surgeryappointments"
	// PersonnelTable is the table the holds the Personnel relation/edge.
	PersonnelTable = "surgeryappointments"
	// PersonnelInverseTable is the table name for the Personnel entity.
	// It exists in this package in order to avoid circular dependency with the "personnel" package.
	PersonnelInverseTable = "personnels"
	// PersonnelColumn is the table column denoting the Personnel relation/edge.
	PersonnelColumn = "Personnel_id"
	// PatientTable is the table the holds the Patient relation/edge.
	PatientTable = "surgeryappointments"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the Patient relation/edge.
	PatientColumn = "Patient_id"
	// SurgerytypeTable is the table the holds the Surgerytype relation/edge.
	SurgerytypeTable = "surgeryappointments"
	// SurgerytypeInverseTable is the table name for the Surgerytype entity.
	// It exists in this package in order to avoid circular dependency with the "surgerytype" package.
	SurgerytypeInverseTable = "surgerytypes"
	// SurgerytypeColumn is the table column denoting the Surgerytype relation/edge.
	SurgerytypeColumn = "Surgerytype"
)

// Columns holds all SQL columns for surgeryappointment fields.
var Columns = []string{
	FieldID,
	FieldAppointTime,
	FieldPhone,
	FieldNote,
	FieldAge,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Surgeryappointment type.
var ForeignKeys = []string{
	"Patient_id",
	"Personnel_id",
	"Surgerytype",
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
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// NoteValidator is a validator for the "note" field. It is called by the builders before save.
	NoteValidator func(string) error
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
)
