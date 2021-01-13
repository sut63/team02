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

type Personnels struct {
	Personnel []personnel
}

type Personnel struct {
	Name 		string
	Department 	string
	User		string
	Password	string
}

type Patients struct {
	Patient []Patient
}

type Patient struct {
	Name 		string
	Birthday	string
	Gender		string

}

type Remedys struct {
	Remedy []Remedy
}

type Remedys struct {
	Remedys 	string
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
	controllers.NewBonediseaseController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewPersonnelController(v1, client)
	controllers.NewRemedyController(v1, client)

	// Set Personnel Data
	personnels := Personnel{
		Personnel: []Personnel{
			Personnel{"Anusorn", "Orthopedics", "Anusorn@gmail.com", "12321"},}
			Personnel{"Kamkeaw", "Orthopedics", "Kamkeaw@gmail.com", "1234"},
		},
	}

	for _, personnels := range personnels.Personnel {
		client.Personnel.
			Create().
			SetName(personnels.Name).
			SetDepartment(personnels.Department).
			SetUser(personnels.User).
			SetPassword(personnels.Password).
			Save(context.Background())
	}

	// Set Patient Data
	patients := Patient{
		Patient: []Patient{
			Patient{"นายแดง", "15/02/1980", "Male"},
			Patient{"นายดำ", "11/03/1980", "Male"},
			Patient{"นางหล่า", "13/11/1980", "Female"},
			
		},
	}

	for _, patients := range patients.Patient {

		client.Patient.
			Create().
			SetName(patients.Name).
			SetBirthday(patients.Birthday).
			SetGender(patients.Gender).
			Save(context.Background())
	}

	// Set Patient Data
	patients := Patient{
		Patient: []Patient{
			Patient{"ทำการผ่าตัดกระดูก"},
			Patient{"ขีดไขข้อ"},
		},
	}

	for _, patients := range patients.Patient {

		client.Patient.
			Create().
			SetRemedy(patients.Remedy).
			Save(context.Background())
	}




	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()

}
