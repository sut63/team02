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
	"github.com/to63/app/ent/remedy"
	
)

// BonediseaseController defines the struct for the Bonedisease controller
type BonediseaseController struct {
	client *ent.Client
	router gin.IRouter
}

// Bonedisease defines the struct for the Bonedisease
type Bonedisease struct {
	PatientID			 int
	PersonnelID          int
	RemedyID             int
	Advice				 string
	AddedTime            string
}

// CreateBonedisease handles POST requests for adding Bonedisease entities
// @Summary Create Bonedisease
// @Description Create Bonedisease
// @ID create-Bonedisease
// @Accept   json
// @Produce  json
// @Param Bonedisease body Bonedisease true "Bonedisease entity"
// @Success 200 {object} ent.Bonedisease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Bonediseases [post]
func (ctl *BonediseaseController) CreateBonedisease(c *gin.Context) {
	obj := Bonedisease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Bonedisease binding failed",
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
	
	remedy, err := ctl.client.Remedy.
	Query().
	Where(remedy.IDEQ(int(obj.RemedyID))).
	Only(context.Background())
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Remedy diagnostic  not found",
		})
		return
	}

	t1 := time.Now()
	t2 := t1.Format("2020-01-02T15:04:14Z07:00")
	time,err := time.Parse(time.RFC3339, t2)

	bondisease, err := ctl.client.Bonedisease.

	Create().
	SetPatient(patient).
	SetPersonnel(personnel).
	SetRemedy(remedy).
	SetAdvice(obj.Advice).
	SetAddedTime(time).
	Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bondisease)
}

// ListBonedisease handles request to get a list of Bonedisease entities
// @Summary List Bonedisease entities
// @Description list Bonedisease entities
// @ID list-Bonedisease
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bonedisease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Bonediseases [get]
func (ctl *BonediseaseController) ListBonedisease(c *gin.Context) {
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

	bonediseases, err := ctl.client.Bonedisease.
		Query().
		WithPersonnel().
		WithPatient().
		WithRemedy().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bonediseases)
}

// DeleteBonedisease handles DELETE requests to delete a Bonedisease entity
// @Summary Delete a Bonedisease entity by ID
// @Description get Bonedisease by ID
// @ID delete-Bonedisease
// @Produce  json
// @Param id path int true "Bonedisease ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Bonediseases/{id} [delete]
func (ctl *BonediseaseController) DeleteBonedisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bonedisease.
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

// NewBonediseaseController creates and registers handles for the Bonedisease controller
func NewBonediseaseController(router gin.IRouter, client *ent.Client) *BonediseaseController {
	bondiseaseController := &BonediseaseController{
		client: client,
		router: router,
	}
	bondiseaseController.register()
	return bondiseaseController
}

// InitBonediseaseController registers routes to the main engine
func (ctl *BonediseaseController) register() {
	Bonediseases := ctl.router.Group("/Bonediseases")

	Bonediseases.GET("", ctl.ListBonedisease)

	// CRUD
	Bonediseases.POST("", ctl.CreateBonedisease)
	Bonediseases.DELETE(":id", ctl.DeleteBonedisease)
}
