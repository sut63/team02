package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/pregnancystatus"
)

// PregnancystatusController defines the struct for the pregnancystatus controller
type PregnancystatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePregnancystatus handles POST requests for adding pregnancystatus entities
// @Summary Create pregnancystatus
// @Description Create pregnancystatus
// @ID create-pregnancystatus
// @Accept   json
// @Produce  json
// @Param pregnancystatus body ent.Pregnancystatus true "Pregnancystatus entity"
// @Success 200 {object} ent.Pregnancystatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Pregnancystatuss [post]
func (ctl *PregnancystatusController) CreatePregnancystatus(c *gin.Context) {
	obj := ent.Pregnancystatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pregnancystatus binding failed",
		})
		return
	}

	pre, err := ctl.client.Pregnancystatus.
		Create().
		SetPregnancystatus(obj.Pregnancystatus).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pre)
}

// GetPregnancystatus handles GET requests to retrieve a pregnancystatus entity
// @Summary Get a pregnancystatus entity by ID
// @Description get pregnancystatus by ID
// @ID get-pregnancystatus
// @Produce  json
// @Param id path int true "Pregnancystatus ID"
// @Success 200 {object} ent.Pregnancystatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pregnancystatuss/{id} [get]
func (ctl *PregnancystatusController) GetPregnancystatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pre, err := ctl.client.Pregnancystatus.
		Query().
		Where(pregnancystatus.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pre)
}

// ListPregnancystatus handles request to get a list of pregnancystatus entities
// @Summary List pregnancystatus entities
// @Description list pregnancystatus entities
// @ID list-pregnancystatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Pregnancystatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pregnancystatuss [get]
func (ctl *PregnancystatusController) ListPregnancystatus(c *gin.Context) {
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

	pregnancystatuss, err := ctl.client.Pregnancystatus.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pregnancystatuss)
}

// NewPregnancystatusController creates and registers handles for the pregnancystatus controller
func NewPregnancystatusController(router gin.IRouter, client *ent.Client) *PregnancystatusController {
	prec := &PregnancystatusController{
		client: client,
		router: router,
	}
	prec.register()
	return prec
}

// InitPregnancystatusController registers routes to the main engine
func (ctl *PregnancystatusController) register() {
	pregnancystatuss := ctl.router.Group("/pregnancystatuss")

	pregnancystatuss.GET("", ctl.ListPregnancystatus)

	// CRUD
	pregnancystatuss.POST("", ctl.CreatePregnancystatus)
	pregnancystatuss.GET(":id", ctl.GetPregnancystatus)
}
