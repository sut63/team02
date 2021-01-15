// Code generated by entc, DO NOT EDIT.

package disease

const (
	// Label holds the string label denoting the disease type in the database.
	Label = "disease"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDisease holds the string denoting the disease field in the database.
	FieldDisease = "disease"

	// EdgeChecksymptom holds the string denoting the checksymptom edge name in mutations.
	EdgeChecksymptom = "Checksymptom"

	// Table holds the table name of the disease in the database.
	Table = "diseases"
	// ChecksymptomTable is the table the holds the Checksymptom relation/edge.
	ChecksymptomTable = "checksymptoms"
	// ChecksymptomInverseTable is the table name for the Checksymptom entity.
	// It exists in this package in order to avoid circular dependency with the "checksymptom" package.
	ChecksymptomInverseTable = "checksymptoms"
	// ChecksymptomColumn is the table column denoting the Checksymptom relation/edge.
	ChecksymptomColumn = "disease_id"
)

// Columns holds all SQL columns for disease fields.
var Columns = []string{
	FieldID,
	FieldDisease,
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
	// DiseaseValidator is a validator for the "disease" field. It is called by the builders before save.
	DiseaseValidator func(string) error
)
