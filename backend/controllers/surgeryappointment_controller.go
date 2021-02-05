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
	"github.com/to63/app/ent/surgeryappointment"
	"github.com/to63/app/ent/surgerytype"
)

// SurgeryappointmentController defines the struct for the surgeryappointment controller
type SurgeryappointmentController struct {
	client *ent.Client
	router gin.IRouter
}

type Surgeryappointment struct {
	Appointtime   string
	Personelid    int
	Patientid     int
	Surgerytypeid int
	Phone         string
	Age           int
	Note          string
}

// CreateSurgeryappointment handles POST requests for adding Surgeryappointment entities
// @Summary Create Surgeryappointment
// @Description Create Surgeryappointment
// @ID create-Surgeryappointment
// @Accept   json
// @Produce  json
// @Param Surgeryappointment body Surgeryappointment true "Surgeryappointment entity"
// @Success 200 {object} ent.Surgeryappointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /surgeryappointments [post]
func (ctl *SurgeryappointmentController) CreateSurgeryappointment(c *gin.Context) {
	obj := Surgeryappointment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Surgeryappointment binding failed",
		})
		return
	}

	ps, err := ctl.client.Personnel.
		Query().
		Where(personnel.IDEQ(int(obj.Personelid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "personnel not found",
		})
		return
	}

	s, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patientid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	ar, err := ctl.client.Surgerytype.
		Query().
		Where(surgerytype.IDEQ(int(obj.Surgerytypeid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "approvedresult not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Appointtime)
	pv, err := ctl.client.Surgeryappointment.
		Create().
		SetAppointTime(times).
		SetPersonnel(ps).
		SetPatient(s).
		SetSurgerytype(ar).
		SetPhone(obj.Phone).
		SetAge(obj.Age).
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
		"data":   pv,
	})
}

// GetSurgeryappointment handles GET requests to retrieve a surgeryappointment entity
// @Summary Get a surgeryappointment entity by ID
// @Description get surgeryappointment by ID
// @ID get-surgeryappointment
// @Produce  json
// @Param id path int true "surgeryappointment ID"
// @Success 200 {object} ent.Surgeryappointment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /surgeryappointments/{id} [get]
func (ctl *SurgeryappointmentController) GetSurgeryappointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Surgeryappointment.
		Query().
		Where(surgeryappointment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSurgeryappointment handles request to get a list of surgeryappointment entities
// @Summary List surgeryappointment entities
// @Description list surgeryappointment entities
// @ID list-surgeryappointment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Surgeryappointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /surgeryappointments [get]
func (ctl *SurgeryappointmentController) ListSurgeryappointment(c *gin.Context) {
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

	surgeryappointments, err := ctl.client.Surgeryappointment.
		Query().
		Limit(limit).
		Offset(offset).
		WithPersonnel().
		WithPatient().
		WithSurgerytype().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, surgeryappointments)
}

// DeleteSurgeryappointment handles DELETE requests to delete a surgeryappointment entity
// @Summary Delete a surgeryappointment entity by ID
// @Description get surgeryappointment by ID
// @ID delete-surgeryappointment
// @Produce  json
// @Param id path int true "Surgeryappointment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /surgeryappointments/{id} [delete]
func (ctl *SurgeryappointmentController) DeleteSurgeryappointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Surgeryappointment.
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

// UpdateSurgeryappointment handles PUT requests to update a surgeryappointment entity
// @Summary Update a surgeryappointment entity by ID
// @Description update surgeryappointment by ID
// @ID update-surgeryappointment
// @Accept   json
// @Produce  json
// @Param id path int true "Surgeryappointment ID"
// @Param surgeryappointment body ent.Surgeryappointment true "surgeryappointment entity"
// @Success 200 {object} ent.Surgeryappointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /surgeryappointments/{id} [put]
func (ctl *SurgeryappointmentController) UpdateSurgeryappointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Surgeryappointment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "surgeryappointment binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Surgeryappointment.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewSurgeryappointmentController creates and registers handles for the surgeryappointment controller
func NewSurgeryappointmentController(router gin.IRouter, client *ent.Client) *SurgeryappointmentController {
	uc := &SurgeryappointmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitSurgeryappointmentController registers routes to the main engine
func (ctl *SurgeryappointmentController) register() {
	surgeryappointments := ctl.router.Group("/surgeryappointments")

	surgeryappointments.GET("", ctl.ListSurgeryappointment)

	// CRUD
	surgeryappointments.POST("", ctl.CreateSurgeryappointment)
	surgeryappointments.GET(":id", ctl.GetSurgeryappointment)
	surgeryappointments.PUT(":id", ctl.UpdateSurgeryappointment)
	surgeryappointments.DELETE(":id", ctl.DeleteSurgeryappointment)
}
