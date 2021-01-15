package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/physicaltherapyroom"
)

// PhysicaltherapyroomController  handles POST requests for adding physicaltherapyroom entities
type PhysicaltherapyroomController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePhysicaltherapyroom handles POST requests for adding physicaltherapyroom entities
// @Summary Create physicaltherapyroom
// @Description Create physicaltherapyroom
// @ID create-physicaltherapyroom
// @Accept   json
// @Produce  json
// @Param physicaltherapyroom body ent.Physicaltherapyroom true "Physicaltherapyroom entity"
// @Success 200 {object} ent.Physicaltherapyroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicaltherapyrooms [post]
func (ctl *PhysicaltherapyroomController) CreatePhysicaltherapyroom(c *gin.Context) {
	obj := ent.Physicaltherapyroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "physicaltherapyroom binding failed",
		})
		return
	}

	phr, err := ctl.client.Physicaltherapyroom.
		Create().
		SetPhysicaltherapyroomname(obj.Physicaltherapyroomname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, phr)
}

// GetPhysicaltherapyroom handles GET requests to retrieve a physicaltherapyroom entity
// @Summary Get a physicaltherapyroom entity by ID
// @Description get physicaltherapyroom by ID
// @ID get-physicaltherapyroom
// @Produce  json
// @Param id path int true "Physicaltherapyroom ID"
// @Success 200 {object} ent.Physicaltherapyroom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicaltherapyrooms/{id} [get]
func (ctl *PhysicaltherapyroomController) GetPhysicaltherapyroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	phr, err := ctl.client.Physicaltherapyroom.
		Query().
		Where(physicaltherapyroom.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, phr)
}

// ListPhysicaltherapyroom handles request to get a list of physicaltherapyroom entities
// @Summary List physicaltherapyroom entities
// @Description list physicaltherapyroom entities
// @ID list-physicaltherapyroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Physicaltherapyroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /physicaltherapyrooms [get]
func (ctl *PhysicaltherapyroomController) ListPhysicaltherapyroom(c *gin.Context) {
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

	physicaltherapyrooms, err := ctl.client.Physicaltherapyroom.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, physicaltherapyrooms)
}

// NewPhysicaltherapyroomController creates and physicaltherapyrooms handles for the physicaltherapyroom controller
func NewPhysicaltherapyroomController(router gin.IRouter, client *ent.Client) *PhysicaltherapyroomController {
	phrc := &PhysicaltherapyroomController{
		client: client,
		router: router,
	}

	phrc.physicaltherapyroom()

	return phrc

}

func (ctl *PhysicaltherapyroomController) physicaltherapyroom() {
	physicaltherapyrooms := ctl.router.Group("/physicaltherapyrooms")

	physicaltherapyrooms.POST("", ctl.CreatePhysicaltherapyroom)
	physicaltherapyrooms.GET(":id", ctl.GetPhysicaltherapyroom)
	physicaltherapyrooms.GET("", ctl.ListPhysicaltherapyroom)

}
