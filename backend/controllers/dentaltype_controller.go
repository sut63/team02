package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/dentaltype"
)

// DentaltypeController defines the struct for the dentaltype controller
type DentaltypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDentaltype handles POST requests for adding dentaltype entities
// @Summary Create dentaltype
// @Description Create dentaltype
// @ID create-dentaltype
// @Accept   json
// @Produce  json
// @Param dentaltype body ent.Dentaltype true "dentaltype entity"
// @Success 200 {object} ent.Dentaltype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentaltypes [post]
func (ctl *DentaltypeController) CreateDentaltype(c *gin.Context) {
	obj := ent.Dentaltype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dentaltype binding failed",
		})
		return
	}

	u, err := ctl.client.Dentaltype.
		Create().
		SetDentaltype(obj.Dentaltype).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetDentaltype handles GET requests to retrieve a dentaltype entity
// @Summary Get a dentaltype entity by ID
// @Description get dentaltype by ID
// @ID get-dentaltype
// @Produce  json
// @Param id path int true "Dentaltype ID"
// @Success 200 {object} ent.Dentaltype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentaltypes/{id} [get]
func (ctl *DentaltypeController) GetDentaltype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Dentaltype.
		Query().
		Where(dentaltype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListDentaltype handles request to get a list of dentaltype entities
// @Summary List dentaltype entities
// @Description list dentaltype entities
// @ID list-dentaltype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Dentaltype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentaltypes [get]
func (ctl *DentaltypeController) ListDentaltype(c *gin.Context) {
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

	dentaltypes, err := ctl.client.Dentaltype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dentaltypes)
}

// DeleteDentaltype handles DELETE requests to delete a dentaltype entity
// @Summary Delete a dentaltype entity by ID
// @Description get dentaltype by ID
// @ID delete-dentaltype
// @Produce  json
// @Param id path int true "Dentaltype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentaltypes/{id} [delete]
func (ctl *DentaltypeController) DeleteDentaltype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Dentaltype.
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

// UpdateDentaltype handles PUT requests to update a dentaltype entity
// @Summary Update a dentaltype entity by ID
// @Description update dentaltype by ID
// @ID update-dentaltype
// @Accept   json
// @Produce  json
// @Param id path int true "Dentaltype ID"
// @Param dentaltype body ent.Dentaltype true "Dentaltype entity"
// @Success 200 {object} ent.Dentaltype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentaltypes/{id} [put]
func (ctl *DentaltypeController) UpdateDentaltype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Dentaltype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dentaltype binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Dentaltype.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewDentaltypeController creates and registers handles for the dentaltype controller
func NewDentaltypeController(router gin.IRouter, client *ent.Client) *DentaltypeController {
	uc := &DentaltypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDentaltypeController registers routes to the main engine
func (ctl *DentaltypeController) register() {
	dentaltypes := ctl.router.Group("/dentaltypes")

	dentaltypes.GET("", ctl.ListDentaltype)

	// CRUD
	dentaltypes.POST("", ctl.CreateDentaltype)
	dentaltypes.GET(":id", ctl.GetDentaltype)
	dentaltypes.PUT(":id", ctl.UpdateDentaltype)
	dentaltypes.DELETE(":id", ctl.DeleteDentaltype)
}
