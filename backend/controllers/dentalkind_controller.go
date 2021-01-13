package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/dentalkind"
)

// DentalkindController defines the struct for the dentalkind controller
type DentalkindController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDentalkind handles POST requests for adding dentalkind entities
// @Summary Create dentalkind
// @Description Create dentalkind
// @ID create-dentalkind
// @Accept   json
// @Produce  json
// @Param dentalkind body ent.Dentalkind true "dentalkind entity"
// @Success 200 {object} ent.Dentalkind
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalkinds [post]
func (ctl *DentalkindController) CreateDentalkind(c *gin.Context) {
	obj := ent.Dentalkind{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dentalkind binding failed",
		})
		return
	}

	u, err := ctl.client.Dentalkind.
		Create().
		SetKindname(obj.Kindname).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetDentalkind handles GET requests to retrieve a dentalkind entity
// @Summary Get a dentalkind entity by ID
// @Description get dentalkind by ID
// @ID get-dentalkind
// @Produce  json
// @Param id path int true "Dentalkind ID"
// @Success 200 {object} ent.Dentalkind
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalkinds/{id} [get]
func (ctl *DentalkindController) GetDentalkind(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Dentalkind.
		Query().
		Where(dentalkind.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListDentalkind handles request to get a list of dentalkind entities
// @Summary List dentalkind entities
// @Description list dentalkind entities
// @ID list-dentalkind
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Dentalkind
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalkinds [get]
func (ctl *DentalkindController) ListDentalkind(c *gin.Context) {
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

	dentalkinds, err := ctl.client.Dentalkind.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dentalkinds)
}

// DeleteDentalkind handles DELETE requests to delete a dentalkind entity
// @Summary Delete a dentalkind entity by ID
// @Description get dentalkind by ID
// @ID delete-dentalkind
// @Produce  json
// @Param id path int true "Dentalkind ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalkinds/{id} [delete]
func (ctl *DentalkindController) DeleteDentalkind(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Dentalkind.
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

// UpdateDentalkind handles PUT requests to update a dentalkind entity
// @Summary Update a dentalkind entity by ID
// @Description update dentalkind by ID
// @ID update-dentalkind
// @Accept   json
// @Produce  json
// @Param id path int true "Dentalkind ID"
// @Param dentalkind body ent.Dentalkind true "Dentalkind entity"
// @Success 200 {object} ent.Dentalkind
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalkinds/{id} [put]
func (ctl *DentalkindController) UpdateDentalkind(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Dentalkind{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dentalkind binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Dentalkind.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewDentalkindController creates and registers handles for the dentalkind controller
func NewDentalkindController(router gin.IRouter, client *ent.Client) *DentalkindController {
	uc := &DentalkindController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDentalkindController registers routes to the main engine
func (ctl *DentalkindController) register() {
	dentalkinds := ctl.router.Group("/dentalkinds")

	dentalkinds.GET("", ctl.ListDentalkind)

	// CRUD
	dentalkinds.POST("", ctl.CreateDentalkind)
	dentalkinds.GET(":id", ctl.GetDentalkind)
	dentalkinds.PUT(":id", ctl.UpdateDentalkind)
	dentalkinds.DELETE(":id", ctl.DeleteDentalkind)
}
