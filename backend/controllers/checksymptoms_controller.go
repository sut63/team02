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

//ChecksymptomsController struct
type ChecksymptomsController struct {
	client *ent.Client
	router gin.IRouter
}

//Checksymptoms struct
type Checksymptoms struct {
	Personnel  			int
	Disease 			int
	DoctorOrderSheet   	int
	Patient      		int
	Date				string
}

// CreateChecksymptoms handles POST requests for adding checksymptoms entities
// @Summary Create checksymptoms
// @Description Create checksymptoms
// @ID create-checksymptoms
// @Accept   json
// @Produce  json
// @Param checksymptoms body Checksymptoms true "Checksymptoms entity"
// @Success 200 {object} ent.Checksymptoms
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checksymptomss [post]
func (ctl *ChecksymptomsController) CreateChecksymptoms(c *gin.Context) {
	obj := Checksymptoms{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "checksymptoms binding failed",
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
			"error": "recordsource not found",
		})
		return
	}

	doctorordersheet, err := ctl.client.DoctorOrderSheet.
		Query().
		Where(doctorordersheet.IDEQ(int(obj.DoctorOrderSheet))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "doctorordersheet not found",
		})
		return
	}
	t1 := time.Now()
	t2 := t1.Format("2020-01-02T15:04:14Z07:00")
	time,err := time.Parse(time.RFC3339, t2)

	checksymptoms, err := ctl.client.Checksymptoms.
		Create().
		SetPatient(patient).
		SetPersonnel(personnel).
		SetDoctorordersheet(doctorordersheet).
		SetDisease(disease).
		SetDate(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, checksymptoms)
}

// ListChecksymptoms handles request to get a list of checksymptoms entities
// @Summary List checksymptoms entities
// @Description list checksymptoms entities
// @ID list-checksymptoms
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Checksymptoms
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checksymptomss [get]
func (ctl *ChecksymptomsController) ListChecksymptoms(c *gin.Context) {
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

	checksymptoms, err := ctl.client.Checksymptoms.
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

	c.JSON(200, checksymptoms)
}

// DeleteChecksymptoms handles DELETE requests to delete a checksymptoms entity
// @Summary Delete a checksymptoms entity by ID
// @Description get checksymptoms by ID
// @ID delete-checksymptoms
// @Produce  json
// @Param id path int true "Checksymptoms ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router/checksymptomss/{id} [delete]
func (ctl *ChecksymptomsController) DeleteChecksymptoms(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Checksymptoms.
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

// NewChecksymptomsController creates and registers handles for the checksymptoms controller
func NewChecksymptomsController(router gin.IRouter, client *ent.Client) *ChecksymptomsController {
	check := &ChecksymptomsController{
		client: client,
		router: router,
	}

	check.register()

	return check

}

func (ctl *ChecksymptomsController) register() {
	checksymptomss := ctl.router.Group("/checksymptomss")

	checksymptomss.POST("", ctl.CreateChecksymptoms)
	checksymptomss.GET("", ctl.ListChecksymptoms)
	checksymptomss.DELETE(":id", ctl.DeleteChecksymptoms)

}
