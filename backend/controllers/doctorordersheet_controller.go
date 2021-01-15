package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/doctorordersheet"
)

// DoctorordersheetController defines the struct for the doctorordersheet controller
type DoctorordersheetController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDoctorordersheet handles POST requests for adding doctorordersheet entities
// @Summary Create doctorordersheet
// @Description Create doctorordersheet
// @ID create-doctorordersheet
// @Accept   json
// @Produce  json
// @Param doctorordersheet body ent.Doctorordersheet true "doctorordersheet entity"
// @Success 200 {object} ent.Doctorordersheet
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets [post]
func (ctl *DoctorordersheetController) CreateDoctorordersheet(c *gin.Context) {
	obj := ent.Doctorordersheet{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctorordersheet binding failed",
		})
		return
	}

	doctorordersheet, err := ctl.client.Doctorordersheet.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	} 

	c.JSON(200, doctorordersheet)
}

// GetDoctorordersheet handles GET requests to retrieve a doctorordersheet entity
// @Summary Get a doctorordersheet entity by ID
// @Description get doctorordersheet by ID
// @ID get-doctorordersheet
// @Produce  json
// @Param id path int true "Doctorordersheet ID"
// @Success 200 {object} ent.Doctorordersheet
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets/{id} [get]
func (ctl *DoctorordersheetController) GetDoctorordersheet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	doctorordersheet, err := ctl.client.Doctorordersheet.
		Query().
		Where(doctorordersheet.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, doctorordersheet)
}

// ListDoctorordersheet handles request to get a list of doctorordersheet entities
// @Summary List doctorordersheet entities
// @Description list doctorordersheet entities
// @ID list-doctorordersheet
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Doctorordersheet
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets [get]
func (ctl *DoctorordersheetController) ListDoctorordersheet(c *gin.Context) {
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

	doctorordersheets, err := ctl.client.Doctorordersheet.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, doctorordersheets)
}

// DeleteDoctorordersheet handles DELETE requests to delete a doctorordersheet entity
// @Summary Delete a doctorordersheet entity by ID
// @Description get doctorordersheet by ID
// @ID delete-doctorordersheet
// @Produce  json
// @Param id path int true "DeleteDoctorordersheet ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets/{id} [delete]
func (ctl *DoctorordersheetController) DeleteDoctorordersheet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Doctorordersheet.
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

// UpdateDoctorordersheet handles PUT requests to update a doctorordersheet entity
// @Summaery Update a doctorordersheet entity by ID
// @Description update doctorordersheet by ID
// @ID update-doctorordersheet
// @Accept   json
// @Produce  json
// @Param id path int true "Doctorordersheet ID"
// @Param doctorordersheet body ent.Doctorordersheet true "Doctorordersheet entity"
// @Success 200 {object} ent.Doctorordersheet
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets/{id} [put]
func (ctl *DoctorordersheetController) UpdateDoctorordersheet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Doctorordersheet{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctorordersheet binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Doctorordersheet.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	} 

	c.JSON(200, u)
}

// NewDoctorordersheetController creates and registers handles for the doctorordersheet controller
func NewDoctorordersheetController(router gin.IRouter, client *ent.Client) *DoctorordersheetController {
	uc := &DoctorordersheetController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDoctorordersheetController registers routes to the main engine
func (ctl *DoctorordersheetController) register() {
	doctorordersheets := ctl.router.Group("/doctorordersheets")

	doctorordersheets.GET("", ctl.ListDoctorordersheet)

	// CRUD
	doctorordersheets.POST("", ctl.CreateDoctorordersheet)
	doctorordersheets.GET(":id", ctl.GetDoctorordersheet)
	doctorordersheets.PUT(":id", ctl.UpdateDoctorordersheet)
	doctorordersheets.DELETE(":id", ctl.DeleteDoctorordersheet)
}