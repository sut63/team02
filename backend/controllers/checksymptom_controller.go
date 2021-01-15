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
	"github.com/to63/app/ent/doctorordersheet"
	"github.com/to63/app/ent/disease"
	
)

//ChecksymptomController struct
type ChecksymptomController struct {
	client *ent.Client
	router gin.IRouter
}

//Checksymptom struct
type Checksymptom struct {
	Personnel  			int
	Disease 			int
	Doctorordersheet   	int
	Patient      		int
	date				string
	times				string
	note				string
}

// CreateChecksymptom handles POST requests for adding checksymptom entities
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
			"error": "checksymptom binding failed",
		})
		return
	}

	personnel, err := ctl.client.Personnel.
		Query().
		Where(personnel.IDEQ(int(obj.Personnel))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "personnel not found",
		})
		return
	}

	patient, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	disease, err := ctl.client.Disease.
		Query().
		Where(disease.IDEQ(int(obj.Disease))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "disease not found",
		})
		return
	}

	

	doctorordersheet, err := ctl.client.Doctorordersheet.
		Query().
		Where(doctorordersheet.IDEQ(int(obj.Doctorordersheet))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "doctorordersheet not found",
		})
		return
	}


	



	dateee, err := time.Parse(time.RFC3339, obj.date)


	checksymptom, err := ctl.client.Checksymptom.
		Create().
		SetPatient(patient).
		SetPersonnel(personnel).
		SetDoctorordersheet(doctorordersheet).
		SetDisease(disease).
		SetDate(dateee).
		SetTimes(obj.times).
		SetNote(obj.note).
		Save(context.Background()) 

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   checksymptom,
	})
}

// ListChecksymptom handles request to get a list of checksymptom entities
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

	checksymptom, err := ctl.client.Checksymptom.
		Query().
		WithPersonnel().
		WithPatient().
		WithDoctorordersheet().
		WithDisease().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, checksymptom)
}

// DeleteChecksymptom handles DELETE requests to delete a checksymptom entity
// @Summary Delete a checksymptom entity by ID
// @Description get checksymptom by ID
// @ID delete-checksymptom
// @Produce  json
// @Param id path int true "Checksymptom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router/checksymptoms/{id} [delete]
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

// NewChecksymptomController creates and registers handles for the checksymptom controller
func NewChecksymptomController(router gin.IRouter, client *ent.Client) *ChecksymptomController {
	check := &ChecksymptomController{
		client: client,
		router: router,
	}

	check.register()

	return check

}

func (ctl *ChecksymptomController) register() {
	checksymptoms := ctl.router.Group("/checksymptoms")
	

	checksymptoms.POST("", ctl.CreateChecksymptom)
	checksymptoms.GET("", ctl.ListChecksymptom)
	checksymptoms.DELETE(":id", ctl.DeleteChecksymptom)

}
