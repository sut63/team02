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
	"github.com/to63/app/ent/physicaltherapyroom"
	"github.com/to63/app/ent/status"
)

//PhysicaltherapyrecordController struct
type PhysicaltherapyrecordController struct {
	client *ent.Client
	router gin.IRouter
}

//Physicaltherapyrecord struct
type Physicaltherapyrecord struct {
	Personnel           int
	Patient             int
	Physicaltherapyroom int
	Status              int
	Time                string
}

// CreatePhysicaltherapyrecord handles POST requests for adding physicaltherapyrecord entities
// @Summary Create physicaltherapyrecord
// @Description Create physicaltherapyrecord
// @ID create-physicaltherapyrecord
// @Accept   json
// @Produce  json
// @Param physicaltherapyrecord body Physicaltherapyrecord true "Physicaltherapyrecord entity"
// @Success 200 {object} ent.Physicaltherapyrecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicaltherapyrecords [post]
func (ctl *PhysicaltherapyrecordController) CreatePhysicaltherapyrecord(c *gin.Context) {
	obj := Physicaltherapyrecord{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "physicaltherapyrecord binding failed",
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

	physicaltherapyroom, err := ctl.client.Physicaltherapyroom.
		Query().
		Where(physicaltherapyroom.IDEQ(int(obj.Physicaltherapyroom))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "physicaltherapyroom not found",
		})
		return
	}

	status, err := ctl.client.Status.
		Query().
		Where(status.IDEQ(int(obj.Status))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "status not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Time)

	physicaltherapyrecord, err := ctl.client.Physicaltherapyrecord.
		Create().
		SetPatient(patient).
		SetPersonnel(personnel).
		SetStatus(status).
		SetPhysicaltherapyroom(physicaltherapyroom).
		SetAppointtime(times).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   physicaltherapyrecord,
	})
}

// ListPhysicaltherapyrecord handles request to get a list of physicaltherapyrecord entities
// @Summary List physicaltherapyrecord entities
// @Description list physicaltherapyrecord entities
// @ID list-physicaltherapyrecord
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Physicaltherapyrecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicaltherapyrecords [get]
func (ctl *PhysicaltherapyrecordController) ListPhysicaltherapyrecord(c *gin.Context) {
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

	physicaltherapyrecord, err := ctl.client.Physicaltherapyrecord.
		Query().
		WithPersonnel().
		WithPatient().
		WithStatus().
		WithPhysicaltherapyroom().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, physicaltherapyrecord)
}

// DeletePhysicaltherapyrecord handles DELETE requests to delete a physicaltherapyrecord entity
// @Summary Delete a physicaltherapyrecord entity by ID
// @Description get physicaltherapyrecord by ID
// @ID delete-physicaltherapyrecord
// @Produce  json
// @Param id path int true "Physicaltherapyrecord ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router/physicaltherapyrecords/{id} [delete]
func (ctl *PhysicaltherapyrecordController) DeletePhysicaltherapyrecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Physicaltherapyrecord.
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

// NewPhysicaltherapyrecordController creates and registers handles for the physicaltherapyrecord controller
func NewPhysicaltherapyrecordController(router gin.IRouter, client *ent.Client) *PhysicaltherapyrecordController {
	physicaltherapyrecordController := &PhysicaltherapyrecordController{
		client: client,
		router: router,
	}

	physicaltherapyrecordController.register()

	return physicaltherapyrecordController

}

func (ctl *PhysicaltherapyrecordController) register() {
	physicaltherapyrecords := ctl.router.Group("/physicaltherapyrecords")

	physicaltherapyrecords.POST("", ctl.CreatePhysicaltherapyrecord)
	physicaltherapyrecords.GET("", ctl.ListPhysicaltherapyrecord)
	physicaltherapyrecords.DELETE(":id", ctl.DeletePhysicaltherapyrecord)

}
