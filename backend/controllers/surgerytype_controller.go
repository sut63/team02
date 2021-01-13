package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/surgerytype"
)

// SurgerytypeController defines the struct for the surgerytype controller
type SurgerytypeController struct {
	client *ent.Client
	router gin.IRouter
}

// GetSurgerytype handles GET requests to retrieve a surgerytype entity
// @Summary Get a surgerytype entity by ID
// @Description get surgerytype by ID
// @ID get-surgerytype
// @Produce  json
// @Param id path int true "surgerytype ID"
// @Success 200 {object} ent.surgerytype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /surgerytypes/{id} [get]
func (ctl *SurgerytypeController) GetSurgerytype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Surgerytype.
		Query().
		Where(surgerytype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSurgerytype handles request to get a list of surgerytype entities
// @Summary List surgerytype entities
// @Description list surgerytype entities
// @ID list-surgerytype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Surgerytype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /surgerytypes [get]
func (ctl *SurgerytypeController) ListSurgerytype(c *gin.Context) {
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

	surgerytypes, err := ctl.client.Surgerytype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, surgerytypes)
}

// DeleteSurgerytype handles DELETE requests to delete a surgerytype entity
// @Summary Delete a surgerytype entity by ID
// @Description get surgerytype by ID
// @ID delete-surgerytype
// @Produce  json
// @Param id path int true "Surgerytype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /surgerytypes/{id} [delete]
func (ctl *SurgerytypeController) DeleteSurgerytype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Surgerytype.
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

// UpdateSurgerytype handles PUT requests to update a surgerytype entity
// @Summary Update a surgerytype entity by ID
// @Description update surgerytype by ID
// @ID update-surgerytype
// @Accept   json
// @Produce  json
// @Param id path int true "Surgerytype ID"
// @Param surgerytype body ent.Surgerytype true "surgerytype entity"
// @Success 200 {object} ent.Surgerytype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /surgerytypes/{id} [put]
func (ctl *SurgerytypeController) UpdateSurgerytype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Surgerytype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "surgerytype binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Surgerytype.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewSurgerytypeController creates and registers handles for the surgerytype controller
func NewSurgerytypeController(router gin.IRouter, client *ent.Client) *SurgerytypeController {
	uc := &SurgerytypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitSurgerytypeController registers routes to the main engine
func (ctl *SurgerytypeController) register() {
	surgerytypes := ctl.router.Group("/surgerytypes")

	surgerytypes.GET("", ctl.ListSurgerytype)

	// CRUD
	surgerytypes.GET(":id", ctl.GetSurgerytype)
	surgerytypes.PUT(":id", ctl.UpdateSurgerytype)
	surgerytypes.DELETE(":id", ctl.DeleteSurgerytype)
}
