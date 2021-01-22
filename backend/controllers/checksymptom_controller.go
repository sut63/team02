package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/disease"
	"github.com/to63/app/ent/doctorordersheet"
	
)

// ChecksymptomController defines the struct for the Checksymptom controller
type ChecksymptomController struct {
	client *ent.Client
	router gin.IRouter
}

// Checksymptom defines the struct for the Checksymptom
type Checksymptom struct {
	PatientID			int
	PersonnelID         int
	DoctorordersheetID  int
	DiseaseID			int
	Date				string
	Phone          		string 
	Note				string
	Identitycard		string

}

// CreateChecksymptom handles POST requests for adding Checksymptom entities
// @Summary Create checksymptom
// @Description Create checksymptom
// @ID create-checksymptom
// @Accept   json
// @Produce  json
// @Param checksymptom body Checksymptom true "Checksymptom entity"
// @Success 200 {object} ent.Checksymptom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checksymptoms [post]
func (ctl *ChecksymptomController) CreateChecksymptom(c *gin.Context) {
	obj := Checksymptom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Checksymptom binding failed",
		})
		return
	}

	patient, err := ctl.client.Patient.
	Query().
	Where(patient.IDEQ(int(obj.PatientID))).
	Only(context.Background())
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	personnel, err := ctl.client.Personnel.
	Query().
	Where(personnel.IDEQ(int(obj.PersonnelID))).
	Only(context.Background())
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": "personnel diagnostic  not found",
		})
		return
	}
	
	doctorordersheet, err := ctl.client.Doctorordersheet.
	Query().
	Where(doctorordersheet.IDEQ(int(obj.DoctorordersheetID))).
	Only(context.Background())
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Doctorordersheet diagnostic  not found",
		})
		return
	}

	disease, err := ctl.client.Disease.
	Query().
	Where(disease.IDEQ(int(obj.DiseaseID))).
	Only(context.Background())
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": "personnel diagnostic  not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Date)

	checksymptom, err := ctl.client.Checksymptom.
	Create().
	SetPersonnel(personnel).
	SetPatient(patient).
	SetDoctorordersheet(doctorordersheet).
	SetDisease(disease).
	SetPhone(obj.Phone).
	SetDate(time).

	SetIdentitycard(obj.Identitycard).
	SetNote(obj.Note).
	Save(context.Background())
	
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

    c.JSON(200, gin.H{
        "status": true,
        "data":   checksymptom,
    })

}


// ListChecksymptom handles request to get a list of Checksymptom entities
// @Summary List checksymptom entities
// @Description list checksymptom entities
// @ID list-checksymptom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Checksymptom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checksymptoms [get]
func (ctl *ChecksymptomController) ListChecksymptom(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	checksymptoms, err := ctl.client.Checksymptom.
		Query().
		WithPersonnel().
		WithPatient().
		WithDoctorordersheet().
		WithDisease().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, checksymptoms)
}



// DeleteChecksymptom handles DELETE requests to delete a Checksymptom entity
// @Summary Delete a checksymptom entity by ID
// @Description get checksymptom by ID
// @ID delete-checksymptom
// @Produce  json
// @Param id path int true "Checksymptom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checksymptoms/{id} [delete]
func (ctl *ChecksymptomController) DeleteChecksymptom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Checksymptom.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewChecksymptomController creates and registers handles for the Checksymptom controller
func NewChecksymptomController(router gin.IRouter, client *ent.Client) *ChecksymptomController {
	checksymptomController := &ChecksymptomController{
		client: client,
		router: router,
	}
	checksymptomController.register()
	return checksymptomController
}

// InitChecksymptomController registers routes to the main engine
func (ctl *ChecksymptomController) register() {
	checksymptoms := ctl.router.Group("/checksymptoms")

	checksymptoms.GET("", ctl.ListChecksymptom)

	// CRUD
	checksymptoms.POST("", ctl.CreateChecksymptom)
	checksymptoms.DELETE(":id", ctl.DeleteChecksymptom)
}
