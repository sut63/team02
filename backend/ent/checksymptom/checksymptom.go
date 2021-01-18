// Code generated by entc, DO NOT EDIT.

package checksymptom

const (
	// Label holds the string label denoting the checksymptom type in the database.
	Label = "checksymptom"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldTimes holds the string denoting the times field in the database.
	FieldTimes = "times"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"

	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgePersonnel holds the string denoting the personnel edge name in mutations.
	EdgePersonnel = "personnel"
	// EdgeDoctorordersheet holds the string denoting the doctorordersheet edge name in mutations.
	EdgeDoctorordersheet = "doctorordersheet"
	// EdgeDisease holds the string denoting the disease edge name in mutations.
	EdgeDisease = "disease"

	// Table holds the table name of the checksymptom in the database.
	Table = "checksymptoms"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "checksymptoms"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "Patient_id"
	// PersonnelTable is the table the holds the personnel relation/edge.
	PersonnelTable = "checksymptoms"
	// PersonnelInverseTable is the table name for the Personnel entity.
	// It exists in this package in order to avoid circular dependency with the "personnel" package.
	PersonnelInverseTable = "personnels"
	// PersonnelColumn is the table column denoting the personnel relation/edge.
	PersonnelColumn = "Personnel_id"
	// DoctorordersheetTable is the table the holds the doctorordersheet relation/edge.
	DoctorordersheetTable = "checksymptoms"
	// DoctorordersheetInverseTable is the table name for the Doctorordersheet entity.
	// It exists in this package in order to avoid circular dependency with the "doctorordersheet" package.
	DoctorordersheetInverseTable = "doctorordersheets"
	// DoctorordersheetColumn is the table column denoting the doctorordersheet relation/edge.
	DoctorordersheetColumn = "Doctorordersheet_id"
	// DiseaseTable is the table the holds the disease relation/edge.
	DiseaseTable = "checksymptoms"
	// DiseaseInverseTable is the table name for the Disease entity.
	// It exists in this package in order to avoid circular dependency with the "disease" package.
	DiseaseInverseTable = "diseases"
	// DiseaseColumn is the table column denoting the disease relation/edge.
	DiseaseColumn = "disease_id"
)

// Columns holds all SQL columns for checksymptom fields.
var Columns = []string{
	FieldID,
	FieldDate,
	FieldTimes,
	FieldNote,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Checksymptom type.
var ForeignKeys = []string{
	"disease_id",
	"Doctorordersheet_id",
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

var (
	// TimesValidator is a validator for the "times" field. It is called by the builders before save.
	TimesValidator func(string) error
	// NoteValidator is a validator for the "note" field. It is called by the builders before save.
	NoteValidator func(string) error
)
