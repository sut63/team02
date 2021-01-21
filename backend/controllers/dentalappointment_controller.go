package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/dentalkind"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
)

// DentalappointmentController defines the struct for the Dentalappointment controller
type DentalappointmentController struct {
	client *ent.Client
	router gin.IRouter
}

// Dentalappointment defines the struct for the Dentalappointment
type Dentalappointment struct {
	PatientID   int
	PersonnelID int
	KindName    int
	AppointTime string
	Amount      int
	Price       int
	Note        string
}

// CreateDentalappointment handles POST requests for adding dentalappointment entities
// @Summary Create dentalappointment
// @Description Create dentalappointment
// @ID create-dentalappointment
// @Accept   json
// @Produce  json
// @Param dentalappointment body Dentalappointment true "dentalappointment entity"
// @Success 200 {object} ent.Dentalappointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalappointments [post]
func (ctl *DentalappointmentController) CreateDentalappointment(c *gin.Context) {
	obj := Dentalappointment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Dentalappointment binding failed",
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

	kindname, err := ctl.client.Dentalkind.
		Query().
		Where(dentalkind.IDEQ(int(obj.KindName))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "personnel diagnostic  not found",
		})
		return
	}

	t1 := time.Now()
	t2 := t1.Format("2020-01-02T15:04:14Z07:00")
	time, err := time.Parse(time.RFC3339, t2)

	dentalapp, err := ctl.client.Dentalappointment.
		Create().
		SetPatient(patient).
		SetPersonnel(personnel).
		SetDentalkind(kindname).
		SetAppointtime(time).
		SetAmount(obj.Amount).
		SetPrice(obj.Price).
		SetNote(obj.Note).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   dentalapp,
	})
}

// ListDentalappointment handles request to get a list of Dentalappointment entities
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

	dentalappointments, err := ctl.client.Dentalappointment.
		Query().
		WithPersonnel().
		WithPatient().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dentalappointments)
}

// DeleteDentalappointment handles DELETE requests to delete a Dentalappointment entity
// @Summary Delete a dentalappointment entity by ID
// @Description get dentalappointment by ID
// @ID delete-dentalappointment
// @Produce  json
// @Param id path int true "dentalappointment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalappointments/{id} [delete]
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

// NewDentalappointmentController creates and registers handles for the Dentalappointment controller
func NewDentalappointmentController(router gin.IRouter, client *ent.Client) *DentalappointmentController {
	bondiseaseController := &DentalappointmentController{
		client: client,
		router: router,
	}
	bondiseaseController.register()
	return bondiseaseController
}

// InitDentalappointmentController registers routes to the main engine
func (ctl *DentalappointmentController) register() {
	Dentalappointments := ctl.router.Group("/dentalappointments")

	Dentalappointments.GET("", ctl.ListDentalappointment)

	// CRUD
	Dentalappointments.POST("", ctl.CreateDentalappointment)
	Dentalappointments.DELETE(":id", ctl.DeleteDentalappointment)
}
