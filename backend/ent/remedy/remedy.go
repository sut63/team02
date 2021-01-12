// Code generated by entc, DO NOT EDIT.

package remedy

const (
	// Label holds the string label denoting the remedy type in the database.
	Label = "remedy"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRemedy holds the string denoting the remedy field in the database.
	FieldRemedy = "remedy"

	// EdgeBonedisease holds the string denoting the bonedisease edge name in mutations.
	EdgeBonedisease = "Bonedisease"

	// Table holds the table name of the remedy in the database.
	Table = "remedies"
	// BonediseaseTable is the table the holds the Bonedisease relation/edge.
	BonediseaseTable = "bonediseases"
	// BonediseaseInverseTable is the table name for the Bonedisease entity.
	// It exists in this package in order to avoid circular dependency with the "bonedisease" package.
	BonediseaseInverseTable = "bonediseases"
	// BonediseaseColumn is the table column denoting the Bonedisease relation/edge.
	BonediseaseColumn = "remedy_id"
)

// Columns holds all SQL columns for remedy fields.
var Columns = []string{
	FieldID,
	FieldRemedy,
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
	// RemedyValidator is a validator for the "remedy" field. It is called by the builders before save.
	RemedyValidator func(string) error
)