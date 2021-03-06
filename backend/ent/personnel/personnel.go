// Code generated by entc, DO NOT EDIT.

package personnel

const (
	// Label holds the string label denoting the personnel type in the database.
	Label = "personnel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgePhysicaltherapyrecord holds the string denoting the physicaltherapyrecord edge name in mutations.
	EdgePhysicaltherapyrecord = "physicaltherapyrecord"
	// EdgeBonedisease holds the string denoting the bonedisease edge name in mutations.
	EdgeBonedisease = "Bonedisease"
	// EdgeChecksymptom holds the string denoting the checksymptom edge name in mutations.
	EdgeChecksymptom = "Checksymptom"
	// EdgeDentalappointment holds the string denoting the dentalappointment edge name in mutations.
	EdgeDentalappointment = "Dentalappointment"
	// EdgeSurgeryappointment holds the string denoting the surgeryappointment edge name in mutations.
	EdgeSurgeryappointment = "Surgeryappointment"
	// EdgeAntenatalinformation holds the string denoting the antenatalinformation edge name in mutations.
	EdgeAntenatalinformation = "Antenatalinformation"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "Department"

	// Table holds the table name of the personnel in the database.
	Table = "personnels"
	// PhysicaltherapyrecordTable is the table the holds the physicaltherapyrecord relation/edge.
	PhysicaltherapyrecordTable = "physicaltherapyrecords"
	// PhysicaltherapyrecordInverseTable is the table name for the Physicaltherapyrecord entity.
	// It exists in this package in order to avoid circular dependency with the "physicaltherapyrecord" package.
	PhysicaltherapyrecordInverseTable = "physicaltherapyrecords"
	// PhysicaltherapyrecordColumn is the table column denoting the physicaltherapyrecord relation/edge.
	PhysicaltherapyrecordColumn = "Personnel_id"
	// BonediseaseTable is the table the holds the Bonedisease relation/edge.
	BonediseaseTable = "bonediseases"
	// BonediseaseInverseTable is the table name for the Bonedisease entity.
	// It exists in this package in order to avoid circular dependency with the "bonedisease" package.
	BonediseaseInverseTable = "bonediseases"
	// BonediseaseColumn is the table column denoting the Bonedisease relation/edge.
	BonediseaseColumn = "Personnel_id"
	// ChecksymptomTable is the table the holds the Checksymptom relation/edge.
	ChecksymptomTable = "checksymptoms"
	// ChecksymptomInverseTable is the table name for the Checksymptom entity.
	// It exists in this package in order to avoid circular dependency with the "checksymptom" package.
	ChecksymptomInverseTable = "checksymptoms"
	// ChecksymptomColumn is the table column denoting the Checksymptom relation/edge.
	ChecksymptomColumn = "Personnel_id"
	// DentalappointmentTable is the table the holds the Dentalappointment relation/edge.
	DentalappointmentTable = "dentalappointments"
	// DentalappointmentInverseTable is the table name for the Dentalappointment entity.
	// It exists in this package in order to avoid circular dependency with the "dentalappointment" package.
	DentalappointmentInverseTable = "dentalappointments"
	// DentalappointmentColumn is the table column denoting the Dentalappointment relation/edge.
	DentalappointmentColumn = "Personnel_id"
	// SurgeryappointmentTable is the table the holds the Surgeryappointment relation/edge.
	SurgeryappointmentTable = "surgeryappointments"
	// SurgeryappointmentInverseTable is the table name for the Surgeryappointment entity.
	// It exists in this package in order to avoid circular dependency with the "surgeryappointment" package.
	SurgeryappointmentInverseTable = "surgeryappointments"
	// SurgeryappointmentColumn is the table column denoting the Surgeryappointment relation/edge.
	SurgeryappointmentColumn = "Personnel_id"
	// AntenatalinformationTable is the table the holds the Antenatalinformation relation/edge.
	AntenatalinformationTable = "antenatalinformations"
	// AntenatalinformationInverseTable is the table name for the Antenatalinformation entity.
	// It exists in this package in order to avoid circular dependency with the "antenatalinformation" package.
	AntenatalinformationInverseTable = "antenatalinformations"
	// AntenatalinformationColumn is the table column denoting the Antenatalinformation relation/edge.
	AntenatalinformationColumn = "Personnel_id"
	// DepartmentTable is the table the holds the Department relation/edge.
	DepartmentTable = "personnels"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "departments"
	// DepartmentColumn is the table column denoting the Department relation/edge.
	DepartmentColumn = "Department"
)

// Columns holds all SQL columns for personnel fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldUser,
	FieldPassword,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Personnel type.
var ForeignKeys = []string{
	"Department",
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// UserValidator is a validator for the "user" field. It is called by the builders before save.
	UserValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
