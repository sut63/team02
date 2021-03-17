// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/to63/app/ent/antenatalinformation"
	"github.com/to63/app/ent/bonedisease"
	"github.com/to63/app/ent/checksymptom"
	"github.com/to63/app/ent/dentalappointment"
	"github.com/to63/app/ent/department"
	"github.com/to63/app/ent/disease"
	"github.com/to63/app/ent/doctorordersheet"
	"github.com/to63/app/ent/gender"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/physicaltherapyrecord"
	"github.com/to63/app/ent/physicaltherapyroom"
	"github.com/to63/app/ent/pregnancystatus"
	"github.com/to63/app/ent/remedy"
	"github.com/to63/app/ent/risks"
	"github.com/to63/app/ent/schema"
	"github.com/to63/app/ent/status"
	"github.com/to63/app/ent/surgeryappointment"
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
	// antenatalinformationDescExaminationresult is the schema descriptor for examinationresult field.
	antenatalinformationDescExaminationresult := antenatalinformationFields[1].Descriptor()
	// antenatalinformation.ExaminationresultValidator is a validator for the "examinationresult" field. It is called by the builders before save.
	antenatalinformation.ExaminationresultValidator = func() func(string) error {
		validators := antenatalinformationDescExaminationresult.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(examinationresult string) error {
			for _, fn := range fns {
				if err := fn(examinationresult); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// antenatalinformationDescAdvice is the schema descriptor for advice field.
	antenatalinformationDescAdvice := antenatalinformationFields[2].Descriptor()
	// antenatalinformation.AdviceValidator is a validator for the "advice" field. It is called by the builders before save.
	antenatalinformation.AdviceValidator = func() func(string) error {
		validators := antenatalinformationDescAdvice.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(advice string) error {
			for _, fn := range fns {
				if err := fn(advice); err != nil {
					return err
				}
			}
			return nil
		}
	}()
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
	// bonediseaseDescTel is the schema descriptor for tel field.
	bonediseaseDescTel := bonediseaseFields[2].Descriptor()
	// bonedisease.TelValidator is a validator for the "tel" field. It is called by the builders before save.
	bonedisease.TelValidator = func() func(string) error {
		validators := bonediseaseDescTel.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
			validators[3].(func(string) error),
		}
		return func(tel string) error {
			for _, fn := range fns {
				if err := fn(tel); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// bonediseaseDescIdentificationCard is the schema descriptor for identificationCard field.
	bonediseaseDescIdentificationCard := bonediseaseFields[3].Descriptor()
	// bonedisease.IdentificationCardValidator is a validator for the "identificationCard" field. It is called by the builders before save.
	bonedisease.IdentificationCardValidator = func() func(string) error {
		validators := bonediseaseDescIdentificationCard.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
			validators[3].(func(string) error),
		}
		return func(identificationCard string) error {
			for _, fn := range fns {
				if err := fn(identificationCard); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	checksymptomFields := schema.Checksymptom{}.Fields()
	_ = checksymptomFields
	// checksymptomDescNote is the schema descriptor for note field.
	checksymptomDescNote := checksymptomFields[1].Descriptor()
	// checksymptom.NoteValidator is a validator for the "note" field. It is called by the builders before save.
	checksymptom.NoteValidator = checksymptomDescNote.Validators[0].(func(string) error)
	// checksymptomDescIdentitycard is the schema descriptor for Identitycard field.
	checksymptomDescIdentitycard := checksymptomFields[2].Descriptor()
	// checksymptom.IdentitycardValidator is a validator for the "Identitycard" field. It is called by the builders before save.
	checksymptom.IdentitycardValidator = func() func(string) error {
		validators := checksymptomDescIdentitycard.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Identitycard string) error {
			for _, fn := range fns {
				if err := fn(_Identitycard); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// checksymptomDescPhone is the schema descriptor for phone field.
	checksymptomDescPhone := checksymptomFields[3].Descriptor()
	// checksymptom.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	checksymptom.PhoneValidator = func() func(string) error {
		validators := checksymptomDescPhone.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(phone string) error {
			for _, fn := range fns {
				if err := fn(phone); err != nil {
					return err
				}
			}
			return nil
		}
	}()
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
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescDepartment is the schema descriptor for department field.
	departmentDescDepartment := departmentFields[0].Descriptor()
	// department.DepartmentValidator is a validator for the "department" field. It is called by the builders before save.
	department.DepartmentValidator = departmentDescDepartment.Validators[0].(func(string) error)
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
	genderFields := schema.Gender{}.Fields()
	_ = genderFields
	// genderDescGender is the schema descriptor for gender field.
	genderDescGender := genderFields[0].Descriptor()
	// gender.GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	gender.GenderValidator = genderDescGender.Validators[0].(func(string) error)
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
	personnelFields := schema.Personnel{}.Fields()
	_ = personnelFields
	// personnelDescName is the schema descriptor for name field.
	personnelDescName := personnelFields[0].Descriptor()
	// personnel.NameValidator is a validator for the "name" field. It is called by the builders before save.
	personnel.NameValidator = personnelDescName.Validators[0].(func(string) error)
	// personnelDescUser is the schema descriptor for user field.
	personnelDescUser := personnelFields[1].Descriptor()
	// personnel.UserValidator is a validator for the "user" field. It is called by the builders before save.
	personnel.UserValidator = personnelDescUser.Validators[0].(func(string) error)
	// personnelDescPassword is the schema descriptor for password field.
	personnelDescPassword := personnelFields[2].Descriptor()
	// personnel.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	personnel.PasswordValidator = personnelDescPassword.Validators[0].(func(string) error)
	physicaltherapyrecordFields := schema.Physicaltherapyrecord{}.Fields()
	_ = physicaltherapyrecordFields
	// physicaltherapyrecordDescIdnumber is the schema descriptor for idnumber field.
	physicaltherapyrecordDescIdnumber := physicaltherapyrecordFields[0].Descriptor()
	// physicaltherapyrecord.IdnumberValidator is a validator for the "idnumber" field. It is called by the builders before save.
	physicaltherapyrecord.IdnumberValidator = func() func(string) error {
		validators := physicaltherapyrecordDescIdnumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(idnumber string) error {
			for _, fn := range fns {
				if err := fn(idnumber); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// physicaltherapyrecordDescAge is the schema descriptor for age field.
	physicaltherapyrecordDescAge := physicaltherapyrecordFields[1].Descriptor()
	// physicaltherapyrecord.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	physicaltherapyrecord.AgeValidator = physicaltherapyrecordDescAge.Validators[0].(func(int) error)
	// physicaltherapyrecordDescTelephone is the schema descriptor for telephone field.
	physicaltherapyrecordDescTelephone := physicaltherapyrecordFields[2].Descriptor()
	// physicaltherapyrecord.TelephoneValidator is a validator for the "telephone" field. It is called by the builders before save.
	physicaltherapyrecord.TelephoneValidator = physicaltherapyrecordDescTelephone.Validators[0].(func(string) error)
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
	surgeryappointmentFields := schema.Surgeryappointment{}.Fields()
	_ = surgeryappointmentFields
	// surgeryappointmentDescPhone is the schema descriptor for phone field.
	surgeryappointmentDescPhone := surgeryappointmentFields[1].Descriptor()
	// surgeryappointment.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	surgeryappointment.PhoneValidator = func() func(string) error {
		validators := surgeryappointmentDescPhone.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(phone string) error {
			for _, fn := range fns {
				if err := fn(phone); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// surgeryappointmentDescNote is the schema descriptor for note field.
	surgeryappointmentDescNote := surgeryappointmentFields[2].Descriptor()
	// surgeryappointment.NoteValidator is a validator for the "note" field. It is called by the builders before save.
	surgeryappointment.NoteValidator = func() func(string) error {
		validators := surgeryappointmentDescNote.Validators
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
	// surgeryappointmentDescAge is the schema descriptor for age field.
	surgeryappointmentDescAge := surgeryappointmentFields[3].Descriptor()
	// surgeryappointment.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	surgeryappointment.AgeValidator = surgeryappointmentDescAge.Validators[0].(func(int) error)
	surgerytypeFields := schema.Surgerytype{}.Fields()
	_ = surgerytypeFields
	// surgerytypeDescTypename is the schema descriptor for typename field.
	surgerytypeDescTypename := surgerytypeFields[0].Descriptor()
	// surgerytype.TypenameValidator is a validator for the "typename" field. It is called by the builders before save.
	surgerytype.TypenameValidator = surgerytypeDescTypename.Validators[0].(func(string) error)
}
