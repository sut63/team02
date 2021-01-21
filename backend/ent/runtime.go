// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/to63/app/ent/antenatalinformation"
	"github.com/to63/app/ent/bonedisease"
	"github.com/to63/app/ent/checksymptom"
	"github.com/to63/app/ent/dentalappointment"
	"github.com/to63/app/ent/disease"
	"github.com/to63/app/ent/doctorordersheet"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/physicaltherapyroom"
	"github.com/to63/app/ent/pregnancystatus"
	"github.com/to63/app/ent/remedy"
	"github.com/to63/app/ent/risks"
	"github.com/to63/app/ent/schema"
	"github.com/to63/app/ent/status"
	"github.com/to63/app/ent/surgerytype"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	antenatalinformationFields := schema.Antenatalinformation{}.Fields()
	_ = antenatalinformationFields
	// antenatalinformationDescGestationalage is the schema descriptor for gestationalage field.
	antenatalinformationDescGestationalage := antenatalinformationFields[0].Descriptor()
	// antenatalinformation.GestationalageValidator is a validator for the "gestationalage" field. It is called by the builders before save.
	antenatalinformation.GestationalageValidator = antenatalinformationDescGestationalage.Validators[0].(func(int) error)
	bonediseaseFields := schema.Bonedisease{}.Fields()
	_ = bonediseaseFields
	// bonediseaseDescAddedTime is the schema descriptor for addedTime field.
	bonediseaseDescAddedTime := bonediseaseFields[0].Descriptor()
	// bonedisease.DefaultAddedTime holds the default value on creation for the addedTime field.
	bonedisease.DefaultAddedTime = bonediseaseDescAddedTime.Default.(func() time.Time)
	// bonediseaseDescAdvice is the schema descriptor for advice field.
	bonediseaseDescAdvice := bonediseaseFields[1].Descriptor()
	// bonedisease.AdviceValidator is a validator for the "advice" field. It is called by the builders before save.
	bonedisease.AdviceValidator = bonediseaseDescAdvice.Validators[0].(func(string) error)
	checksymptomFields := schema.Checksymptom{}.Fields()
	_ = checksymptomFields
	// checksymptomDescTimes is the schema descriptor for times field.
	checksymptomDescTimes := checksymptomFields[1].Descriptor()
	// checksymptom.TimesValidator is a validator for the "times" field. It is called by the builders before save.
	checksymptom.TimesValidator = checksymptomDescTimes.Validators[0].(func(string) error)
	// checksymptomDescNote is the schema descriptor for note field.
	checksymptomDescNote := checksymptomFields[2].Descriptor()
	// checksymptom.NoteValidator is a validator for the "note" field. It is called by the builders before save.
	checksymptom.NoteValidator = checksymptomDescNote.Validators[0].(func(string) error)
	dentalappointmentFields := schema.Dentalappointment{}.Fields()
	_ = dentalappointmentFields
	// dentalappointmentDescAmount is the schema descriptor for amount field.
	dentalappointmentDescAmount := dentalappointmentFields[1].Descriptor()
	// dentalappointment.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	dentalappointment.AmountValidator = dentalappointmentDescAmount.Validators[0].(func(int) error)
	// dentalappointmentDescPrice is the schema descriptor for price field.
	dentalappointmentDescPrice := dentalappointmentFields[2].Descriptor()
	// dentalappointment.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	dentalappointment.PriceValidator = dentalappointmentDescPrice.Validators[0].(func(int) error)
	// dentalappointmentDescNote is the schema descriptor for note field.
	dentalappointmentDescNote := dentalappointmentFields[3].Descriptor()
	// dentalappointment.NoteValidator is a validator for the "note" field. It is called by the builders before save.
	dentalappointment.NoteValidator = func() func(string) error {
		validators := dentalappointmentDescNote.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(note string) error {
			for _, fn := range fns {
				if err := fn(note); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	diseaseFields := schema.Disease{}.Fields()
	_ = diseaseFields
	// diseaseDescDisease is the schema descriptor for disease field.
	diseaseDescDisease := diseaseFields[0].Descriptor()
	// disease.DiseaseValidator is a validator for the "disease" field. It is called by the builders before save.
	disease.DiseaseValidator = diseaseDescDisease.Validators[0].(func(string) error)
	doctorordersheetFields := schema.Doctorordersheet{}.Fields()
	_ = doctorordersheetFields
	// doctorordersheetDescName is the schema descriptor for Name field.
	doctorordersheetDescName := doctorordersheetFields[0].Descriptor()
	// doctorordersheet.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	doctorordersheet.NameValidator = doctorordersheetDescName.Validators[0].(func(string) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescName is the schema descriptor for name field.
	patientDescName := patientFields[0].Descriptor()
	// patient.NameValidator is a validator for the "name" field. It is called by the builders before save.
	patient.NameValidator = patientDescName.Validators[0].(func(string) error)
	// patientDescBirthday is the schema descriptor for birthday field.
	patientDescBirthday := patientFields[1].Descriptor()
	// patient.BirthdayValidator is a validator for the "birthday" field. It is called by the builders before save.
	patient.BirthdayValidator = patientDescBirthday.Validators[0].(func(string) error)
	// patientDescGender is the schema descriptor for gender field.
	patientDescGender := patientFields[2].Descriptor()
	// patient.GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	patient.GenderValidator = patientDescGender.Validators[0].(func(string) error)
	personnelFields := schema.Personnel{}.Fields()
	_ = personnelFields
	// personnelDescName is the schema descriptor for name field.
	personnelDescName := personnelFields[0].Descriptor()
	// personnel.NameValidator is a validator for the "name" field. It is called by the builders before save.
	personnel.NameValidator = personnelDescName.Validators[0].(func(string) error)
	// personnelDescDepartment is the schema descriptor for department field.
	personnelDescDepartment := personnelFields[1].Descriptor()
	// personnel.DepartmentValidator is a validator for the "department" field. It is called by the builders before save.
	personnel.DepartmentValidator = personnelDescDepartment.Validators[0].(func(string) error)
	// personnelDescUser is the schema descriptor for user field.
	personnelDescUser := personnelFields[2].Descriptor()
	// personnel.UserValidator is a validator for the "user" field. It is called by the builders before save.
	personnel.UserValidator = personnelDescUser.Validators[0].(func(string) error)
	// personnelDescPassword is the schema descriptor for password field.
	personnelDescPassword := personnelFields[3].Descriptor()
	// personnel.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	personnel.PasswordValidator = personnelDescPassword.Validators[0].(func(string) error)
	physicaltherapyroomFields := schema.Physicaltherapyroom{}.Fields()
	_ = physicaltherapyroomFields
	// physicaltherapyroomDescPhysicaltherapyroomname is the schema descriptor for physicaltherapyroomname field.
	physicaltherapyroomDescPhysicaltherapyroomname := physicaltherapyroomFields[0].Descriptor()
	// physicaltherapyroom.PhysicaltherapyroomnameValidator is a validator for the "physicaltherapyroomname" field. It is called by the builders before save.
	physicaltherapyroom.PhysicaltherapyroomnameValidator = physicaltherapyroomDescPhysicaltherapyroomname.Validators[0].(func(string) error)
	pregnancystatusFields := schema.Pregnancystatus{}.Fields()
	_ = pregnancystatusFields
	// pregnancystatusDescPregnancystatus is the schema descriptor for Pregnancystatus field.
	pregnancystatusDescPregnancystatus := pregnancystatusFields[0].Descriptor()
	// pregnancystatus.PregnancystatusValidator is a validator for the "Pregnancystatus" field. It is called by the builders before save.
	pregnancystatus.PregnancystatusValidator = pregnancystatusDescPregnancystatus.Validators[0].(func(string) error)
	remedyFields := schema.Remedy{}.Fields()
	_ = remedyFields
	// remedyDescRemedy is the schema descriptor for remedy field.
	remedyDescRemedy := remedyFields[0].Descriptor()
	// remedy.RemedyValidator is a validator for the "remedy" field. It is called by the builders before save.
	remedy.RemedyValidator = remedyDescRemedy.Validators[0].(func(string) error)
	risksFields := schema.Risks{}.Fields()
	_ = risksFields
	// risksDescRisks is the schema descriptor for Risks field.
	risksDescRisks := risksFields[0].Descriptor()
	// risks.RisksValidator is a validator for the "Risks" field. It is called by the builders before save.
	risks.RisksValidator = risksDescRisks.Validators[0].(func(string) error)
	statusFields := schema.Status{}.Fields()
	_ = statusFields
	// statusDescStatusname is the schema descriptor for statusname field.
	statusDescStatusname := statusFields[0].Descriptor()
	// status.StatusnameValidator is a validator for the "statusname" field. It is called by the builders before save.
	status.StatusnameValidator = statusDescStatusname.Validators[0].(func(string) error)
	surgerytypeFields := schema.Surgerytype{}.Fields()
	_ = surgerytypeFields
	// surgerytypeDescTypename is the schema descriptor for typename field.
	surgerytypeDescTypename := surgerytypeFields[0].Descriptor()
	// surgerytype.TypenameValidator is a validator for the "typename" field. It is called by the builders before save.
	surgerytype.TypenameValidator = surgerytypeDescTypename.Validators[0].(func(string) error)
}
