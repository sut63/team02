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
	"github.com/to63/app/ent/dentaltype"
)

//DentalappointmentController struct
type DentalappointmentController struct {
	client *ent.Client
	router gin.IRouter
}

//Dentalappointment struct
type Dentalappointment struct {
	RecordPersonnel  int
	RecordDentaltype int
	RecordPatient    int
	RecordTime       string
}

// CreateDentalappointment handles POST requests for adding dentalappointment entities
// @Summary Create dentalappointment
// @Description Create dentalappointment
// @ID create-dentalappointment
// @Accept   json
// @Produce  json
// @Param dentalappointment body Dentalappointment true "Dentalappointment entity"
// @Success 200 {object} ent.Dentalappointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalappointments [post]
func (ctl *DentalappointmentController) CreateDentalappointment(c *gin.Context) {
	obj := Dentalappointment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dentalappointment binding failed",
		})
		return
	}

	per, err := ctl.client.Personnel.
		Query().
		Where(personnel.IDEQ(int(obj.RecordPersonnel))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "recordsource not found",
		})
		return
	}

	pat, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.RecordPatient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "recordsource not found",
		})
		return
	}

	type, err := ctl.client.Dentaltype.
		Query().
		Where(dentaltype.IDEQ(int(obj.RecordDentaltype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "recordsource not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.RecordTime)

	pv, err := ctl.client.Dentalappointment.
		Create().
		//SetAddTime(times).
		//SetPa2re(pat).
		//SetEx2re(exp).
		//SetPer2re(per).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pv)
}

// ListDentalappointment handles request to get a list of dentalappointment entities
// @Summary List dentalappointment entities
// @Description list dentalappointment entities
// @ID list-dentalappointment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Dentalappointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalappointments [get]
func (ctl *DentalappointmentController) ListDentalappointment(c *gin.Context) {
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

	rec, err := ctl.client.Dentalappointment.
		Query().
		//WithPa2re().
		//WithEx2re().
		//WithPer2re().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, rec)
}

// DeleteDentalappointment handles DELETE requests to delete a dentalappointment entity
// @Summary Delete a dentalappointment entity by ID
// @Description get dentalappointment by ID
// @ID delete-dentalappointment
// @Produce  json
// @Param id path int true "Dentalappointment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router/dentalappointments/{id} [delete]
func (ctl *DentalappointmentController) DeleteDentalappointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Dentalappointment.
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

// NewDentalappointmentController creates and registers handles for the dentalappointment controller
func NewDentalappointmentController(router gin.IRouter, client *ent.Client) *DentalappointmentController {
	pvc := &DentalappointmentController{
		client: client,
		router: router,
	}

	pvc.register()

	return pvc

}

func (ctl *DentalappointmentController) register() {
	dentalappointments := ctl.router.Group("/dentalappointments")

	dentalappointments.POST("", ctl.CreateDentalappointment)
	dentalappointments.GET("", ctl.ListDentalappointment)
	dentalappointments.DELETE(":id", ctl.DeleteDentalappointment)

}
