// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AntenatalinformationsColumns holds the columns for the "antenatalinformations" table.
	AntenatalinformationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gestationalage", Type: field.TypeInt},
		{Name: "examinationresult", Type: field.TypeString},
		{Name: "advice", Type: field.TypeString},
		{Name: "time", Type: field.TypeTime},
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "Personnel_id", Type: field.TypeInt, Nullable: true},
		{Name: "pregnancystatus_antenatalinformation", Type: field.TypeInt, Nullable: true},
		{Name: "risks_antenatalinformation", Type: field.TypeInt, Nullable: true},
	}
	// AntenatalinformationsTable holds the schema information for the "antenatalinformations" table.
	AntenatalinformationsTable = &schema.Table{
		Name:       "antenatalinformations",
		Columns:    AntenatalinformationsColumns,
		PrimaryKey: []*schema.Column{AntenatalinformationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "antenatalinformations_patients_Antenatalinformation",
				Columns: []*schema.Column{AntenatalinformationsColumns[5]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "antenatalinformations_personnels_Antenatalinformation",
				Columns: []*schema.Column{AntenatalinformationsColumns[6]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "antenatalinformations_pregnancystatuses_Antenatalinformation",
				Columns: []*schema.Column{AntenatalinformationsColumns[7]},

				RefColumns: []*schema.Column{PregnancystatusesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "antenatalinformations_risks_Antenatalinformation",
				Columns: []*schema.Column{AntenatalinformationsColumns[8]},

				RefColumns: []*schema.Column{RisksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BonediseasesColumns holds the columns for the "bonediseases" table.
	BonediseasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "advice", Type: field.TypeString},
		{Name: "tel", Type: field.TypeString, Size: 10},
		{Name: "identification_card", Type: field.TypeString, Size: 13},
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "Personnel_id", Type: field.TypeInt, Nullable: true},
		{Name: "remedy_id", Type: field.TypeInt, Nullable: true},
	}
	// BonediseasesTable holds the schema information for the "bonediseases" table.
	BonediseasesTable = &schema.Table{
		Name:       "bonediseases",
		Columns:    BonediseasesColumns,
		PrimaryKey: []*schema.Column{BonediseasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bonediseases_patients_Bonedisease",
				Columns: []*schema.Column{BonediseasesColumns[5]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bonediseases_personnels_Bonedisease",
				Columns: []*schema.Column{BonediseasesColumns[6]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bonediseases_remedies_Bonedisease",
				Columns: []*schema.Column{BonediseasesColumns[7]},

				RefColumns: []*schema.Column{RemediesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChecksymptomsColumns holds the columns for the "checksymptoms" table.
	ChecksymptomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "note", Type: field.TypeString, Size: 40},
		{Name: "identitycard", Type: field.TypeString, Size: 13},
		{Name: "phone", Type: field.TypeString, Size: 10},
		{Name: "disease_id", Type: field.TypeInt, Nullable: true},
		{Name: "Doctorordersheet_id", Type: field.TypeInt, Nullable: true},
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "Personnel_id", Type: field.TypeInt, Nullable: true},
	}
	// ChecksymptomsTable holds the schema information for the "checksymptoms" table.
	ChecksymptomsTable = &schema.Table{
		Name:       "checksymptoms",
		Columns:    ChecksymptomsColumns,
		PrimaryKey: []*schema.Column{ChecksymptomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "checksymptoms_diseases_Checksymptom",
				Columns: []*schema.Column{ChecksymptomsColumns[5]},

				RefColumns: []*schema.Column{DiseasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checksymptoms_doctorordersheets_Checksymptom",
				Columns: []*schema.Column{ChecksymptomsColumns[6]},

				RefColumns: []*schema.Column{DoctorordersheetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checksymptoms_patients_Checksymptom",
				Columns: []*schema.Column{ChecksymptomsColumns[7]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checksymptoms_personnels_Checksymptom",
				Columns: []*schema.Column{ChecksymptomsColumns[8]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DentalappointmentsColumns holds the columns for the "dentalappointments" table.
	DentalappointmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "appointtime", Type: field.TypeTime},
		{Name: "amount", Type: field.TypeInt},
		{Name: "price", Type: field.TypeInt},
		{Name: "note", Type: field.TypeString, Size: 25},
		{Name: "kindname", Type: field.TypeInt, Nullable: true},
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "Personnel_id", Type: field.TypeInt, Nullable: true},
	}
	// DentalappointmentsTable holds the schema information for the "dentalappointments" table.
	DentalappointmentsTable = &schema.Table{
		Name:       "dentalappointments",
		Columns:    DentalappointmentsColumns,
		PrimaryKey: []*schema.Column{DentalappointmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dentalappointments_dentalkinds_Dentalappointment",
				Columns: []*schema.Column{DentalappointmentsColumns[5]},

				RefColumns: []*schema.Column{DentalkindsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dentalappointments_patients_Dentalappointment",
				Columns: []*schema.Column{DentalappointmentsColumns[6]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dentalappointments_personnels_Dentalappointment",
				Columns: []*schema.Column{DentalappointmentsColumns[7]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DentalkindsColumns holds the columns for the "dentalkinds" table.
	DentalkindsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "kindname", Type: field.TypeString},
	}
	// DentalkindsTable holds the schema information for the "dentalkinds" table.
	DentalkindsTable = &schema.Table{
		Name:        "dentalkinds",
		Columns:     DentalkindsColumns,
		PrimaryKey:  []*schema.Column{DentalkindsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "department", Type: field.TypeString},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:        "departments",
		Columns:     DepartmentsColumns,
		PrimaryKey:  []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DiseasesColumns holds the columns for the "diseases" table.
	DiseasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "disease", Type: field.TypeString},
	}
	// DiseasesTable holds the schema information for the "diseases" table.
	DiseasesTable = &schema.Table{
		Name:        "diseases",
		Columns:     DiseasesColumns,
		PrimaryKey:  []*schema.Column{DiseasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DoctorordersheetsColumns holds the columns for the "doctorordersheets" table.
	DoctorordersheetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// DoctorordersheetsTable holds the schema information for the "doctorordersheets" table.
	DoctorordersheetsTable = &schema.Table{
		Name:        "doctorordersheets",
		Columns:     DoctorordersheetsColumns,
		PrimaryKey:  []*schema.Column{DoctorordersheetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gender", Type: field.TypeString},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "birthday", Type: field.TypeString},
		{Name: "Gender", Type: field.TypeInt, Nullable: true},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:       "patients",
		Columns:    PatientsColumns,
		PrimaryKey: []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "patients_genders_Patient",
				Columns: []*schema.Column{PatientsColumns[3]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PersonnelsColumns holds the columns for the "personnels" table.
	PersonnelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "Department", Type: field.TypeInt, Nullable: true},
	}
	// PersonnelsTable holds the schema information for the "personnels" table.
	PersonnelsTable = &schema.Table{
		Name:       "personnels",
		Columns:    PersonnelsColumns,
		PrimaryKey: []*schema.Column{PersonnelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "personnels_departments_Personnel",
				Columns: []*schema.Column{PersonnelsColumns[4]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PhysicaltherapyrecordsColumns holds the columns for the "physicaltherapyrecords" table.
	PhysicaltherapyrecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeInt},
		{Name: "note", Type: field.TypeString},
		{Name: "appointtime", Type: field.TypeTime},
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "Personnel_id", Type: field.TypeInt, Nullable: true},
		{Name: "physicaltherapyroomname_id", Type: field.TypeInt, Nullable: true},
		{Name: "statusname_id", Type: field.TypeInt, Nullable: true},
	}
	// PhysicaltherapyrecordsTable holds the schema information for the "physicaltherapyrecords" table.
	PhysicaltherapyrecordsTable = &schema.Table{
		Name:       "physicaltherapyrecords",
		Columns:    PhysicaltherapyrecordsColumns,
		PrimaryKey: []*schema.Column{PhysicaltherapyrecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "physicaltherapyrecords_patients_physicaltherapyrecord",
				Columns: []*schema.Column{PhysicaltherapyrecordsColumns[4]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "physicaltherapyrecords_personnels_physicaltherapyrecord",
				Columns: []*schema.Column{PhysicaltherapyrecordsColumns[5]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "physicaltherapyrecords_physicaltherapyrooms_physicaltherapyrecord",
				Columns: []*schema.Column{PhysicaltherapyrecordsColumns[6]},

				RefColumns: []*schema.Column{PhysicaltherapyroomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "physicaltherapyrecords_status_physicaltherapyrecord",
				Columns: []*schema.Column{PhysicaltherapyrecordsColumns[7]},

				RefColumns: []*schema.Column{StatusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PhysicaltherapyroomsColumns holds the columns for the "physicaltherapyrooms" table.
	PhysicaltherapyroomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "physicaltherapyroomname", Type: field.TypeString},
	}
	// PhysicaltherapyroomsTable holds the schema information for the "physicaltherapyrooms" table.
	PhysicaltherapyroomsTable = &schema.Table{
		Name:        "physicaltherapyrooms",
		Columns:     PhysicaltherapyroomsColumns,
		PrimaryKey:  []*schema.Column{PhysicaltherapyroomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PregnancystatusesColumns holds the columns for the "pregnancystatuses" table.
	PregnancystatusesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "pregnancystatus", Type: field.TypeString},
	}
	// PregnancystatusesTable holds the schema information for the "pregnancystatuses" table.
	PregnancystatusesTable = &schema.Table{
		Name:        "pregnancystatuses",
		Columns:     PregnancystatusesColumns,
		PrimaryKey:  []*schema.Column{PregnancystatusesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RemediesColumns holds the columns for the "remedies" table.
	RemediesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "remedy", Type: field.TypeString},
	}
	// RemediesTable holds the schema information for the "remedies" table.
	RemediesTable = &schema.Table{
		Name:        "remedies",
		Columns:     RemediesColumns,
		PrimaryKey:  []*schema.Column{RemediesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RisksColumns holds the columns for the "risks" table.
	RisksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "risks", Type: field.TypeString},
	}
	// RisksTable holds the schema information for the "risks" table.
	RisksTable = &schema.Table{
		Name:        "risks",
		Columns:     RisksColumns,
		PrimaryKey:  []*schema.Column{RisksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "statusname", Type: field.TypeString},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:        "status",
		Columns:     StatusColumns,
		PrimaryKey:  []*schema.Column{StatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SurgeryappointmentsColumns holds the columns for the "surgeryappointments" table.
	SurgeryappointmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "appoint_time", Type: field.TypeTime},
		{Name: "phone", Type: field.TypeString},
		{Name: "note", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "Personnel_id", Type: field.TypeInt, Nullable: true},
		{Name: "Surgerytype", Type: field.TypeInt, Nullable: true},
	}
	// SurgeryappointmentsTable holds the schema information for the "surgeryappointments" table.
	SurgeryappointmentsTable = &schema.Table{
		Name:       "surgeryappointments",
		Columns:    SurgeryappointmentsColumns,
		PrimaryKey: []*schema.Column{SurgeryappointmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "surgeryappointments_patients_Surgeryappointment",
				Columns: []*schema.Column{SurgeryappointmentsColumns[5]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "surgeryappointments_personnels_Surgeryappointment",
				Columns: []*schema.Column{SurgeryappointmentsColumns[6]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "surgeryappointments_surgerytypes_Surgeryappointment",
				Columns: []*schema.Column{SurgeryappointmentsColumns[7]},

				RefColumns: []*schema.Column{SurgerytypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SurgerytypesColumns holds the columns for the "surgerytypes" table.
	SurgerytypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "typename", Type: field.TypeString},
	}
	// SurgerytypesTable holds the schema information for the "surgerytypes" table.
	SurgerytypesTable = &schema.Table{
		Name:        "surgerytypes",
		Columns:     SurgerytypesColumns,
		PrimaryKey:  []*schema.Column{SurgerytypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AntenatalinformationsTable,
		BonediseasesTable,
		ChecksymptomsTable,
		DentalappointmentsTable,
		DentalkindsTable,
		DepartmentsTable,
		DiseasesTable,
		DoctorordersheetsTable,
		GendersTable,
		PatientsTable,
		PersonnelsTable,
		PhysicaltherapyrecordsTable,
		PhysicaltherapyroomsTable,
		PregnancystatusesTable,
		RemediesTable,
		RisksTable,
		StatusTable,
		SurgeryappointmentsTable,
		SurgerytypesTable,
	}
)

func init() {
	AntenatalinformationsTable.ForeignKeys[0].RefTable = PatientsTable
	AntenatalinformationsTable.ForeignKeys[1].RefTable = PersonnelsTable
	AntenatalinformationsTable.ForeignKeys[2].RefTable = PregnancystatusesTable
	AntenatalinformationsTable.ForeignKeys[3].RefTable = RisksTable
	BonediseasesTable.ForeignKeys[0].RefTable = PatientsTable
	BonediseasesTable.ForeignKeys[1].RefTable = PersonnelsTable
	BonediseasesTable.ForeignKeys[2].RefTable = RemediesTable
	ChecksymptomsTable.ForeignKeys[0].RefTable = DiseasesTable
	ChecksymptomsTable.ForeignKeys[1].RefTable = DoctorordersheetsTable
	ChecksymptomsTable.ForeignKeys[2].RefTable = PatientsTable
	ChecksymptomsTable.ForeignKeys[3].RefTable = PersonnelsTable
	DentalappointmentsTable.ForeignKeys[0].RefTable = DentalkindsTable
	DentalappointmentsTable.ForeignKeys[1].RefTable = PatientsTable
	DentalappointmentsTable.ForeignKeys[2].RefTable = PersonnelsTable
	PatientsTable.ForeignKeys[0].RefTable = GendersTable
	PersonnelsTable.ForeignKeys[0].RefTable = DepartmentsTable
	PhysicaltherapyrecordsTable.ForeignKeys[0].RefTable = PatientsTable
	PhysicaltherapyrecordsTable.ForeignKeys[1].RefTable = PersonnelsTable
	PhysicaltherapyrecordsTable.ForeignKeys[2].RefTable = PhysicaltherapyroomsTable
	PhysicaltherapyrecordsTable.ForeignKeys[3].RefTable = StatusTable
	SurgeryappointmentsTable.ForeignKeys[0].RefTable = PatientsTable
	SurgeryappointmentsTable.ForeignKeys[1].RefTable = PersonnelsTable
	SurgeryappointmentsTable.ForeignKeys[2].RefTable = SurgerytypesTable
}
