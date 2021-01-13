package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/doctorordersheet"
)

// DoctorOrderSheetController defines the struct for the doctorordersheet controller
type DoctorOrderSheetController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDoctorOrderSheet handles POST requests for adding doctorordersheet entities
// @Summary Create doctorordersheet
// @Description Create doctorordersheet
// @ID create-doctorordersheet
// @Accept   json
// @Produce  json
// @Param doctorordersheet body ent.DoctorOrderSheet true "doctorordersheet entity"
// @Success 200 {object} ent.DoctorOrderSheet
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets [post]
func (ctl *DoctorOrderSheetController) CreateDoctorOrderSheet(c *gin.Context) {
	obj := ent.DoctorOrderSheet{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctorordersheet binding failed",
		})
		return
	}

	doctorordersheet, err := ctl.client.DoctorOrderSheet.
		Create().
		SetName(obj.Name).
		SetTime(obj.Time).
		SetNote(obj.Note).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	} 

	c.JSON(200, doctorordersheet)
}

// GetDoctorOrderSheet handles GET requests to retrieve a doctorordersheet entity
// @Summary Get a doctorordersheet entity by ID
// @Description get doctorordersheet by ID
// @ID get-doctorordersheet
// @Produce  json
// @Param id path int true "DoctorOrderSheet ID"
// @Success 200 {object} ent.DoctorOrderSheet
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets/{id} [get]
func (ctl *DoctorOrderSheetController) GetDoctorOrderSheet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	doctorordersheet, err := ctl.client.DoctorOrderSheet.
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

// ListDoctorOrderSheet handles request to get a list of doctorordersheet entities
// @Summary List doctorordersheet entities
// @Description list doctorordersheet entities
// @ID list-doctorordersheet
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.DoctorOrderSheet
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets [get]
func (ctl *DoctorOrderSheetController) ListDoctorOrderSheet(c *gin.Context) {
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

	doctorordersheets, err := ctl.client.DoctorOrderSheet.
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

// DeleteDoctorOrderSheet handles DELETE requests to delete a doctorordersheet entity
// @Summary Delete a doctorordersheet entity by ID
// @Description get doctorordersheet by ID
// @ID delete-doctorordersheet
// @Produce  json
// @Param id path int true "DeleteDoctorOrderSheet ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets/{id} [delete]
func (ctl *DoctorOrderSheetController) DeleteDoctorOrderSheet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.DoctorOrderSheet.
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

// UpdateDoctorOrderSheet handles PUT requests to update a doctorordersheet entity
// @Summaery Update a doctorordersheet entity by ID
// @Description update doctorordersheet by ID
// @ID update-doctorordersheet
// @Accept   json
// @Produce  json
// @Param id path int true "DoctorOrderSheet ID"
// @Param doctorordersheet body ent.DoctorOrderSheet true "DoctorOrderSheet entity"
// @Success 200 {object} ent.DoctorOrderSheet
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorordersheets/{id} [put]
func (ctl *DoctorOrderSheetController) UpdateDoctorOrderSheet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.DoctorOrderSheet{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctorordersheet binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.DoctorOrderSheet.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewDoctorOrderSheetController creates and registers handles for the doctorordersheet controller
func NewDoctorOrderSheetController(router gin.IRouter, client *ent.Client) *DoctorOrderSheetController {
	uc := &DoctorOrderSheetController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDoctorOrderSheetController registers routes to the main engine
func (ctl *DoctorOrderSheetController) register() {
	diseases := ctl.router.Group("/doctorordersheets")

	diseases.GET("", ctl.ListDoctorOrderSheet)

	// CRUD
	diseases.POST("", ctl.CreateDoctorOrderSheet)
	diseases.GET(":id", ctl.GetDoctorOrderSheet)
	diseases.PUT(":id", ctl.UpdateDoctorOrderSheet)
	diseases.DELETE(":id", ctl.DeleteDoctorOrderSheet)
}