package controllers

import (
	"context" 
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/disease"
)

// DiseaseController defines the struct for the Disease controller
type DiseaseController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDisease handles POST requests for adding Disease entities
// @Summary Create disease
// @Description Create disease
// @ID create-disease
// @Accept   json
// @Produce  json
// @Param disease body ent.Disease true "Disease entity"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [post]
func (ctl *DiseaseController) CreateDisease(c *gin.Context) {
	obj := ent.Disease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Disease binding failed",
		})
		return
	}
 
	ar, err := ctl.client.Disease.
		Create().
		SetDisease(obj.Disease).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "disease saving failed",
		})
		return
	}

	c.JSON(200, ar)
}

// GetDisease handles GET requests to retrieve a Disease entity
// @Summary Get a disease entity by ID
// @Description get disease by ID
// @ID get-disease
// @Produce  json
// @Param id path int true "Disease ID"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [get]
func (ctl *DiseaseController) GetDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	disease, err := ctl.client.Disease.
		Query().
		Where(disease.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, disease)
}

// ListDisease handles request to get a list of Disease entities
// @Summary List disease entities
// @Description list disease entities
// @ID list-disease
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [get]
func (ctl *DiseaseController) ListDisease(c *gin.Context) {
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

	disease, err := ctl.client.Disease.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, disease)
}

// DeleteDisease handles DELETE requests to delete a Disease entity
// @Summary Delete a disease entity by ID
// @Description get disease by ID
// @ID delete-disease
// @Produce  json
// @Param id path int true "Disease ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [delete]
func (ctl *DiseaseController) DeleteDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Disease.
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

// UpdateDisease handles PUT requests to update a Disease entity
// @Summary Update a disease entity by ID
// @Description update disease by ID
// @ID update-disease
// @Accept   json
// @Produce  json
// @Param id path int true "Disease ID"
// @Param disease body ent.Disease true "Disease entity"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [put]
func (ctl *DiseaseController) UpdateDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Disease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Disease binding failed",
		})
		return
	}
	obj.ID = int(id)
	disease, err := ctl.client.Disease.
		UpdateOneID(int(id)).
		SetDisease(obj.Disease).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, disease)
}

// NewDiseaseController creates and registers handles for the DiseaseController
func NewDiseaseController(router gin.IRouter, client *ent.Client) *DiseaseController {
	diseaseController := &DiseaseController{
		client: client,
		router: router,
	}
	diseaseController.register()
	return diseaseController
}

// InitDiseaseController registers routes to the main engine
func (ctl *DiseaseController) register() {
	Disease := ctl.router.Group("/diseases")
	Disease.GET("", ctl.ListDisease)
	// CRUD
	Disease.POST("", ctl.CreateDisease)
	Disease.GET(":id", ctl.GetDisease)
	Disease.PUT(":id", ctl.UpdateDisease)
	Disease.DELETE(":id", ctl.DeleteDisease)
}
