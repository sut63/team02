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
		{Name: "gestationalage", Type: field.TypeString},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "Personnel_id", Type: field.TypeInt, Nullable: true},
		{Name: "pregnancystatus_antenatalinformation", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "risks_antenatalinformation", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// AntenatalinformationsTable holds the schema information for the "antenatalinformations" table.
	AntenatalinformationsTable = &schema.Table{
		Name:       "antenatalinformations",
		Columns:    AntenatalinformationsColumns,
		PrimaryKey: []*schema.Column{AntenatalinformationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "antenatalinformations_patients_Antenatalinformation",
				Columns: []*schema.Column{AntenatalinformationsColumns[3]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "antenatalinformations_personnels_Antenatalinformation",
				Columns: []*schema.Column{AntenatalinformationsColumns[4]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "antenatalinformations_pregnancystatuses_Antenatalinformation",
				Columns: []*schema.Column{AntenatalinformationsColumns[5]},

				RefColumns: []*schema.Column{PregnancystatusesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "antenatalinformations_risks_Antenatalinformation",
				Columns: []*schema.Column{AntenatalinformationsColumns[6]},

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
				Columns: []*schema.Column{BonediseasesColumns[3]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bonediseases_personnels_Bonedisease",
				Columns: []*schema.Column{BonediseasesColumns[4]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bonediseases_remedies_Bonedisease",
				Columns: []*schema.Column{BonediseasesColumns[5]},

				RefColumns: []*schema.Column{RemediesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChecksymptomsColumns holds the columns for the "checksymptoms" table.
	ChecksymptomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "Disease_id", Type: field.TypeInt, Nullable: true},
		{Name: "DoctorOrderSheet_id", Type: field.TypeInt, Nullable: true},
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
				Symbol:  "checksymptoms_diseases_Checksymptoms",
				Columns: []*schema.Column{ChecksymptomsColumns[1]},

				RefColumns: []*schema.Column{DiseasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checksymptoms_doctor_order_sheets_Checksymptoms",
				Columns: []*schema.Column{ChecksymptomsColumns[2]},

				RefColumns: []*schema.Column{DoctorOrderSheetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checksymptoms_patients_Checksymptoms",
				Columns: []*schema.Column{ChecksymptomsColumns[3]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checksymptoms_personnels_Checksymptoms",
				Columns: []*schema.Column{ChecksymptomsColumns[4]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DentalappointmentsColumns holds the columns for the "dentalappointments" table.
	DentalappointmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "appoint_time", Type: field.TypeTime},
		{Name: "typename", Type: field.TypeInt, Nullable: true},
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
				Symbol:  "dentalappointments_dentaltypes_Dentalappointment",
				Columns: []*schema.Column{DentalappointmentsColumns[2]},

				RefColumns: []*schema.Column{DentaltypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dentalappointments_patients_Dentalappointment",
				Columns: []*schema.Column{DentalappointmentsColumns[3]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dentalappointments_personnels_Dentalappointment",
				Columns: []*schema.Column{DentalappointmentsColumns[4]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DentaltypesColumns holds the columns for the "dentaltypes" table.
	DentaltypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "typename", Type: field.TypeString},
	}
	// DentaltypesTable holds the schema information for the "dentaltypes" table.
	DentaltypesTable = &schema.Table{
		Name:        "dentaltypes",
		Columns:     DentaltypesColumns,
		PrimaryKey:  []*schema.Column{DentaltypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DiseasesColumns holds the columns for the "diseases" table.
	DiseasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// DiseasesTable holds the schema information for the "diseases" table.
	DiseasesTable = &schema.Table{
		Name:        "diseases",
		Columns:     DiseasesColumns,
		PrimaryKey:  []*schema.Column{DiseasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DoctorOrderSheetsColumns holds the columns for the "doctor_order_sheets" table.
	DoctorOrderSheetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "time", Type: field.TypeString},
	}
	// DoctorOrderSheetsTable holds the schema information for the "doctor_order_sheets" table.
	DoctorOrderSheetsTable = &schema.Table{
		Name:        "doctor_order_sheets",
		Columns:     DoctorOrderSheetsColumns,
		PrimaryKey:  []*schema.Column{DoctorOrderSheetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "birthday", Type: field.TypeString},
		{Name: "gender", Type: field.TypeString},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:        "patients",
		Columns:     PatientsColumns,
		PrimaryKey:  []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PersonnelsColumns holds the columns for the "personnels" table.
	PersonnelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "department", Type: field.TypeString},
		{Name: "user", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// PersonnelsTable holds the schema information for the "personnels" table.
	PersonnelsTable = &schema.Table{
		Name:        "personnels",
		Columns:     PersonnelsColumns,
		PrimaryKey:  []*schema.Column{PersonnelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PhysicaltherapyrecordsColumns holds the columns for the "physicaltherapyrecords" table.
	PhysicaltherapyrecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "Personnel_id", Type: field.TypeInt, Nullable: true},
		{Name: "physicaltherapyroom_physicaltherapyrecord", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "status_physicaltherapyrecord", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PhysicaltherapyrecordsTable holds the schema information for the "physicaltherapyrecords" table.
	PhysicaltherapyrecordsTable = &schema.Table{
		Name:       "physicaltherapyrecords",
		Columns:    PhysicaltherapyrecordsColumns,
		PrimaryKey: []*schema.Column{PhysicaltherapyrecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "physicaltherapyrecords_patients_physicaltherapyrecord",
				Columns: []*schema.Column{PhysicaltherapyrecordsColumns[2]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "physicaltherapyrecords_personnels_physicaltherapyrecord",
				Columns: []*schema.Column{PhysicaltherapyrecordsColumns[3]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "physicaltherapyrecords_physicaltherapyrooms_physicaltherapyrecord",
				Columns: []*schema.Column{PhysicaltherapyrecordsColumns[4]},

				RefColumns: []*schema.Column{PhysicaltherapyroomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "physicaltherapyrecords_status_physicaltherapyrecord",
				Columns: []*schema.Column{PhysicaltherapyrecordsColumns[5]},

				RefColumns: []*schema.Column{StatusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PhysicaltherapyroomsColumns holds the columns for the "physicaltherapyrooms" table.
	PhysicaltherapyroomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "physical_therapy_room_name", Type: field.TypeString},
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
		{Name: "status_name", Type: field.TypeString},
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
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "Personnel_id", Type: field.TypeInt, Nullable: true},
		{Name: "surgerytype_surgeryappointment", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// SurgeryappointmentsTable holds the schema information for the "surgeryappointments" table.
	SurgeryappointmentsTable = &schema.Table{
		Name:       "surgeryappointments",
		Columns:    SurgeryappointmentsColumns,
		PrimaryKey: []*schema.Column{SurgeryappointmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "surgeryappointments_patients_Surgeryappointment",
				Columns: []*schema.Column{SurgeryappointmentsColumns[2]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "surgeryappointments_personnels_Surgeryappointment",
				Columns: []*schema.Column{SurgeryappointmentsColumns[3]},

				RefColumns: []*schema.Column{PersonnelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "surgeryappointments_surgerytypes_Surgeryappointment",
				Columns: []*schema.Column{SurgeryappointmentsColumns[4]},

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
		DentaltypesTable,
		DiseasesTable,
		DoctorOrderSheetsTable,
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
	ChecksymptomsTable.ForeignKeys[1].RefTable = DoctorOrderSheetsTable
	ChecksymptomsTable.ForeignKeys[2].RefTable = PatientsTable
	ChecksymptomsTable.ForeignKeys[3].RefTable = PersonnelsTable
	DentalappointmentsTable.ForeignKeys[0].RefTable = DentaltypesTable
	DentalappointmentsTable.ForeignKeys[1].RefTable = PatientsTable
	DentalappointmentsTable.ForeignKeys[2].RefTable = PersonnelsTable
	PhysicaltherapyrecordsTable.ForeignKeys[0].RefTable = PatientsTable
	PhysicaltherapyrecordsTable.ForeignKeys[1].RefTable = PersonnelsTable
	PhysicaltherapyrecordsTable.ForeignKeys[2].RefTable = PhysicaltherapyroomsTable
	PhysicaltherapyrecordsTable.ForeignKeys[3].RefTable = StatusTable
	SurgeryappointmentsTable.ForeignKeys[0].RefTable = PatientsTable
	SurgeryappointmentsTable.ForeignKeys[1].RefTable = PersonnelsTable
	SurgeryappointmentsTable.ForeignKeys[2].RefTable = SurgerytypesTable
}
