package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/antenatalinformation"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/pregnancystatus"
	"github.com/to63/app/ent/risks"
)

// AntenatalinformationController defines the struct for the antenatalinformation controller
type AntenatalinformationController struct {
	client *ent.Client
	router gin.IRouter
}

//Antenatalinformation struct
type Antenatalinformation struct {
	Personnel       int
	Patient         int
	Pregnancystatus int
	Risks           int
	Gestationalage  int
	Time            string
}

// CreateAntenatalinformation handles POST requests for adding antenatalinformation entities
// @Summary Create antenatalinformation
// @Description Create antenatalinformation
// @ID create-antenatalinformation
// @Accept   json
// @Produce  json
// @Param antenatalinformation body ent.Antenatalinformation true "Antenatalinformation entity"
// @Success 200 {object} ent.Antenatalinformation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatalinformations [post]
func (ctl *AntenatalinformationController) CreateAntenatalinformation(c *gin.Context) {
	obj := Antenatalinformation{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "antenatalinformation binding failed",
		})
		return
	}

	per, err := ctl.client.Personnel.
		Query().
		Where(personnel.IDEQ(int(obj.Personnel))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "personnel not found",
		})
		return
	}

	pat, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	pre, err := ctl.client.Pregnancystatus.
		Query().
		Where(pregnancystatus.IDEQ(int(obj.Pregnancystatus))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "pregnancystatus not found",
		})
		return
	}

	r, err := ctl.client.Risks.
		Query().
		Where(risks.IDEQ(int(obj.Risks))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "risks not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Time)

	ai, err := ctl.client.Antenatalinformation.
		Create().
		SetPersonnel(per).
		SetPatient(pat).
		SetPregnancystatus(pre).
		SetRisks(r).
		SetGestationalage(obj.Gestationalage).
		SetTime(times).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ai)
}

// GetAntenatalinformation handles GET requests to retrieve a antenatalinformation entity
// @Summary Get a antenatalinformation entity by ID
// @Description get antenatalinformation by ID
// @ID get-antenatalinformation
// @Produce  json
// @Param id path int true "Antenatalinformation ID"
// @Success 200 {object} ent.Antenatalinformation
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatalinformations/{id} [get]
func (ctl *AntenatalinformationController) GetAntenatalinformation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ai, err := ctl.client.Antenatalinformation.
		Query().
		Where(antenatalinformation.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ai)
}

// ListAntenatalinformation handles request to get a list of antenatalinformation entities
// @Summary List antenatalinformation entities
// @Description list antenatalinformation entities
// @ID list-antenatalinformation
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Antenatalinformation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatalinformations [get]
func (ctl *AntenatalinformationController) ListAntenatalinformation(c *gin.Context) {
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

	antenatalinformations, err := ctl.client.Antenatalinformation.
		Query().
		WithPatient().
		WithPersonnel().
		WithPregnancystatus().
		WithRisks().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, antenatalinformations)
}

// DeleteAntenatalinformation handles DELETE requests to delete a antenatalinformation entity
// @Summary Delete a antenatalinformation entity by ID
// @Description get antenatalinformation by ID
// @ID delete-antenatalinformation
// @Produce  json
// @Param id path int true "Antenatalinformation ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatalinformations/{id} [delete]
func (ctl *AntenatalinformationController) DeleteAntenatalinformation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Antenatalinformation.
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

// UpdateAntenatalinformation handles PUT requests to update a antenatalinformation entity
// @Summary Update a antenatalinformation entity by ID
// @Description update antenatalinformation by ID
// @ID update-antenatalinformation
// @Accept   json
// @Produce  json
// @Param id path int true "Antenatalinformation ID"
// @Param antenatalinformation body ent.Antenatalinformation true "Antenatalinformation entity"
// @Success 200 {object} ent.Antenatalinformation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /antenatalinformations/{id} [put]
func (ctl *AntenatalinformationController) UpdateAntenatalinformation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Antenatalinformation{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "antenatalinformation binding failed",
		})
		return
	}
	obj.ID = int(id)
	ai, err := ctl.client.Antenatalinformation.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ai)
}

// NewAntenatalinformationController creates and registers handles for the antenatalinformation controller
func NewAntenatalinformationController(router gin.IRouter, client *ent.Client) *AntenatalinformationController {
	aic := &AntenatalinformationController{
		client: client,
		router: router,
	}
	aic.register()
	return aic
}

// InitAntenatalinformationController registers routes to the main engine
func (ctl *AntenatalinformationController) register() {
	antenatalinformations := ctl.router.Group("/antenatalinformations")

	antenatalinformations.GET("", ctl.ListAntenatalinformation)

	// CRUD
	antenatalinformations.POST("", ctl.CreateAntenatalinformation)
	antenatalinformations.GET(":id", ctl.GetAntenatalinformation)
	antenatalinformations.PUT(":id", ctl.UpdateAntenatalinformation)
	antenatalinformations.DELETE(":id", ctl.DeleteAntenatalinformation)
}
