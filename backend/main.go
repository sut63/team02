package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/to63/app/controllers"
	"github.com/to63/app/ent"
)

//Personnels with
type Personnels struct {
	Personnel []Personnel
}

//Personnel with
type Personnel struct {
	Name       string
	Department int
	User       string
	Password   string
}

//Patients with
type Patients struct {
	Patient []Patient
}

//Patient with
type Patient struct {
	Name     string
	Birthday string
	Gender   int
}

//Remedys with
type Remedys struct {
	Remedy []Remedy
}

//Remedy with
type Remedy struct {
	Remedy string
}

//Diseases with
type Diseases struct {
	Disease []Disease
}

//Disease with
type Disease struct {
	Name string
}

//Doctorordersheets with
type Doctorordersheets struct {
	Doctorordersheet []Doctorordersheet
}

//Doctorordersheet with
type Doctorordersheet struct {
	Name string
}

//Deltalkinds with
type Dentalkinds struct {
	Dentalkind []Dentalkind
}

//Deltalkind with
type Dentalkind struct {
	KindName string
}

//Pregnancystatuss with
type Pregnancystatuss struct {
	Pregnancystatus []Pregnancystatus
}

//Pregnancystatus with
type Pregnancystatus struct {
	Pregnancystatus string
}

//Riskss with
type Riskss struct {
	Risks []Risks
}

//Risks with
type Risks struct {
	Risks string
}

//Physicaltherapyrooms with
type Physicaltherapyrooms struct {
	Physicaltherapyroom []Physicaltherapyroom
}

//Physicaltherapyroom with
type Physicaltherapyroom struct {
	Physicaltherapyroom string
}

//Statuss with
type Statuss struct {
	Status []Status
}

//Status with
type Status struct {
	Status string
}

//Surgerytypes with
type Surgerytypes struct {
	Surgerytype []Surgerytype
}

//Surgerytype with
type Surgerytype struct {
	Surgerytype string
}

//Genders with
type Genders struct {
	Gender []Gender
}

//Gender with
type Gender struct {
	Gender string
}

//Departments with
type Departments struct {
	Department []Department
}

//Department with
type Department struct {
	Department string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewChecksymptomController(v1, client)
	controllers.NewDiseaseController(v1, client)          //ของเจม
	controllers.NewDoctorordersheetController(v1, client) //ของเจม
	controllers.NewBonediseaseController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewPersonnelController(v1, client)
	controllers.NewRemedyController(v1, client)
	controllers.NewDentalkindController(v1, client)
	controllers.NewDentalappointmentController(v1, client)
	controllers.NewPregnancystatusController(v1, client)
	controllers.NewRisksController(v1, client)
	controllers.NewAntenatalinformationController(v1, client)
	controllers.NewPhysicaltherapyrecordController(v1, client)
	controllers.NewPhysicaltherapyroomController(v1, client)
	controllers.NewStatusController(v1, client)
	controllers.NewSurgerytypeController(v1, client)
	controllers.NewSurgeryappointmentController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewDepartmentController(v1, client)

	// Set Gender Data
	genders := Genders{
		Gender: []Gender{
			Gender{"ชาย"},
			Gender{"หญิง"},
		},
	}

	for _, genders := range genders.Gender {

		client.Gender.
			Create().
			SetGender(genders.Gender).
			Save(context.Background())
	}

	// Set Department Data
	departments := Departments{
		Department: []Department{
			Department{"Orthopedics"},
			Department{"Dentist"},
			Department{"Doctor"},
			Department{"Physicaltherapy"},
			Department{"Obstetrician"},
		},
	}

	for _, departments := range departments.Department {

		client.Department.
			Create().
			SetDepartment(departments.Department).
			Save(context.Background())
	}

	// Set Personnel Data
	personnels := Personnels{
		Personnel: []Personnel{
			Personnel{"Anusorn", 1, "Anusorn@email.com", "12321"},
			Personnel{"Kamkeaw", 1, "Kamkeaw@email.com", "12321"},
			Personnel{"Nutchaya", 2, "nutchaya@email.com", "12321"},
			Personnel{"Thanawut", 3, "thanawut@email.com", "12321"},
			Personnel{"Jay", 4, "Jay@email.com", "12321"},
			Personnel{"Varissara", 5, "Varissara@email.com", "12321"},
		},
	}

	for _, personnels := range personnels.Personnel {
		client.Personnel.
			Create().
			SetName(personnels.Name).
			SetDepartmentID(personnels.Department).
			SetUser(personnels.User).
			SetPassword(personnels.Password).
			Save(context.Background())
	}

	// Set Patient Data
	patients := Patients{
		Patient: []Patient{
			Patient{"นายแดง", "15/02/1980", 1},
			Patient{"นายดำ", "11/03/1980", 1},
			Patient{"นางหล่า", "13/11/1980", 2},
		},
	}

	for _, patients := range patients.Patient {

		client.Patient.
			Create().
			SetName(patients.Name).
			SetBirthday(patients.Birthday).
			SetGenderID(patients.Gender).
			Save(context.Background())
	}

	diseases := Diseases{
		Disease: []Disease{
			Disease{"โรคกระดูก"},
			Disease{"กายภาพบำบัด"},
		},
	}

	for _, diseases := range diseases.Disease {

		client.Disease.
			Create().
			SetDisease(diseases.Name).
			Save(context.Background())
	}

	// Set DoctorOrderSheet Data
	doctorordersheets := []string{"นายแพทย์อนุสรณ์ ศรีพรหม", "นายแพทย์วัชรพงศ์ ซึ้งศิริทรัพย์", "นายแพทย์หญิงวริศรา ตันดิลกตระกูล"}
	for _, doctorordersheet := range doctorordersheets {
		client.Doctorordersheet.
			Create().
			SetName(doctorordersheet).
			Save(context.Background())
	}

	// Set Remedy Data
	remedys := Remedys{
		Remedy: []Remedy{
			Remedy{"ทำการผ่าตัดกระดูก"},
			Remedy{"ขีดไขข้อ"},
		},
	}

	for _, remedys := range remedys.Remedy {

		client.Remedy.
			Create().
			SetRemedy(remedys.Remedy).
			Save(context.Background())
	}

	// Set Dentalkind Data
	dentalkinds := Dentalkinds{
		Dentalkind: []Dentalkind{
			Dentalkind{"ตรวจช่องปาก"},
			Dentalkind{"ถอนฟัน"},
			Dentalkind{"ขูดหินปูน"},
			Dentalkind{"ผ่าฟันคุด"},
			Dentalkind{"เอ็กซเรย์ฟัน"},
			Dentalkind{"รักษารากฟัน"},
			Dentalkind{"อุดฟัน"},
			Dentalkind{"เติมฟัน"},
			Dentalkind{"ทำฟันปลอม"},
		},
	}

	for _, dentalkinds := range dentalkinds.Dentalkind {

		client.Dentalkind.
			Create().
			SetKindname(dentalkinds.KindName).
			Save(context.Background())
	}

	// Set Pregnancystatus Data
	Pregnancystatuss := Pregnancystatuss{
		Pregnancystatus: []Pregnancystatus{
			Pregnancystatus{"ตั้งครรภ์"},
			Pregnancystatus{"ภาวะเสี่ยง"},
			Pregnancystatus{"รอคลอด"},
			Pregnancystatus{"แท้งบุตร"},
		},
	}

	for _, Pregnancystatuss := range Pregnancystatuss.Pregnancystatus {

		client.Pregnancystatus.
			Create().
			SetPregnancystatus(Pregnancystatuss.Pregnancystatus).
			Save(context.Background())
	}

	// Set Risks Data
	Riskss := Riskss{
		Risks: []Risks{
			Risks{"เคยผ่าตัดที่มดลูก"},
			Risks{"มีประวัติโรคแทรกซ้อน"},
			Risks{"มีประวัติแท้งตั้งแต่ 3 ครั้งขึ้นไป"},
			Risks{"เป็นโรคโลหิตจาง"},
			Risks{"ตรวจ HIV ได้ผลบวก"},
		},
	}

	for _, Riskss := range Riskss.Risks {

		client.Risks.
			Create().
			SetRisks(Riskss.Risks).
			Save(context.Background())
	}

	// Set Physicaltherapyroom Data
	physicaltherapyrooms := Physicaltherapyrooms{
		Physicaltherapyroom: []Physicaltherapyroom{
			Physicaltherapyroom{"Rhysicaltherapyroom01"},
			Physicaltherapyroom{"Rhysicaltherapyroom02"},
		},
	}

	for _, physicaltherapyrooms := range physicaltherapyrooms.Physicaltherapyroom {

		client.Physicaltherapyroom.
			Create().
			SetPhysicaltherapyroomname(physicaltherapyrooms.Physicaltherapyroom).
			Save(context.Background())
	}

	// Set Status Data
	statuss := Statuss{
		Status: []Status{
			Status{"ดีขึ้น"},
			Status{"ต้องติดตามอาการ"},
		},
	}

	for _, statuss := range statuss.Status {

		client.Status.
			Create().
			SetStatusname(statuss.Status).
			Save(context.Background())
	}

	// Set Surgerytype Data
	surgerytypes := Surgerytypes{
		Surgerytype: []Surgerytype{
			Surgerytype{"ผ่าตัดทั่วไป"},
			Surgerytype{"ผ่าคลอด"},
			Surgerytype{"ผ่าตัดกระดูก"},
		},
	}

	for _, surgerytypes := range surgerytypes.Surgerytype {

		client.Surgerytype.
			Create().
			SetTypename(surgerytypes.Surgerytype).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()

}
