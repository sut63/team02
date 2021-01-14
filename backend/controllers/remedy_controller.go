package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/remedy"
)

// RemedyController defines the struct for the Remedy controller
type RemedyController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateRemedy handles POST requests for adding Remedy entities
// @Summary Create remedy
// @Description Create remedy
// @ID create-remedy
// @Accept   json
// @Produce  json
// @Param remedy body ent.Remedy true "Remedy entity"
// @Success 200 {object} ent.Remedy
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /remedys [post]
func (ctl *RemedyController) CreateRemedy(c *gin.Context) {
	obj := ent.Remedy{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Remedy binding failed",
		})
		return
	}

	ar, err := ctl.client.Remedy.
		Create().
		SetRemedy(obj.Remedy).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ar)
}

// GetRemedy handles GET requests to retrieve a Remedy entity
// @Summary Get a remedy entity by ID
// @Description get remedy by ID
// @ID get-remedy
// @Produce  json
// @Param id path int true "Remedy ID"
// @Success 200 {object} ent.Remedy
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /remedys/{id} [get]
func (ctl *RemedyController) GetRemedy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	remedy, err := ctl.client.Remedy.
		Query().
		Where(remedy.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, remedy)
}

// ListRemedy handles request to get a list of Remedy entities
// @Summary List remedy entities
// @Description list remedy entities
// @ID list-remedy
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Remedy
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /remedys [get]
func (ctl *RemedyController) ListRemedy(c *gin.Context) {
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

	remedy, err := ctl.client.Remedy.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, remedy)
}

// DeleteRemedy handles DELETE requests to delete a Remedy entity
// @Summary Delete a remedy entity by ID
// @Description get remedy by ID
// @ID delete-remedy
// @Produce  json
// @Param id path int true "Remedy ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /remedys/{id} [delete]
func (ctl *RemedyController) DeleteRemedy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Remedy.
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

// UpdateRemedy handles PUT requests to update a Remedy entity
// @Summary Update a remedy entity by ID
// @Description update remedy by ID
// @ID update-remedy
// @Accept   json
// @Produce  json
// @Param id path int true "Remedy ID"
// @Param Remedy body ent.Remedy true "Remedy entity"
// @Success 200 {object} ent.Remedy
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /remedys/{id} [put]
func (ctl *RemedyController) UpdateRemedy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Remedy{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Remedy binding failed",
		})
		return
	}
	obj.ID = int(id)
	remedy, err := ctl.client.Remedy.
		UpdateOneID(int(id)).
		SetRemedy(obj.Remedy).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, remedy)
}

// NewRemedyController creates and registers handles for the RemedyController
func NewRemedyController(router gin.IRouter, client *ent.Client) *RemedyController {
	remedyController := &RemedyController{
		client: client,
		router: router,
	}
	remedyController.register()
	return remedyController
}

// InitRemedyController registers routes to the main engine
func (ctl *RemedyController) register() {
	Remedy := ctl.router.Group("/remedys")
	Remedy.GET("", ctl.ListRemedy)
	// CRUD
	Remedy.POST("", ctl.CreateRemedy)
	Remedy.GET(":id", ctl.GetRemedy)
	Remedy.PUT(":id", ctl.UpdateRemedy)
	Remedy.DELETE(":id", ctl.DeleteRemedy)
}
