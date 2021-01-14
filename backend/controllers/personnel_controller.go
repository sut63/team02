package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/to63/app/ent"
	"github.com/to63/app/ent/personnel"
)

// PersonnelController defines the struct for the Personnel controller
type PersonnelController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePersonnel handles POST requests for adding Personnel entities
// @Summary Create personnel
// @Description Create personnel
// @ID create-personnel
// @Accept   json
// @Produce  json
// @Param personnel body ent.Personnel true "Personnel entity"
// @Success 200 {object} ent.Personnel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personnels [post]
func (ctl *PersonnelController) CreatePersonnel(c *gin.Context) {
	obj := ent.Personnel{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Personnel binding failed",
		})
		return
	}

	personnel, err := ctl.client.Personnel.
		Create().
		SetName(obj.Name).
		SetDepartment(obj.Department).
		SetUser(obj.User).
		SetPassword(obj.Password).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, personnel)
}

// GetPersonnel handles GET requests to retrieve a Personnel entity
// @Summary Get a personnel entity by ID
// @Description get personnel by ID
// @ID get-personnel
// @Produce  json
// @Param id path int true "Personnel ID"
// @Success 200 {object} ent.Personnel
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personnels/{id} [get]
func (ctl *PersonnelController) GetPersonnel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	personnel, err := ctl.client.Personnel.
		Query().
		Where(personnel.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, personnel)
}

// ListPersonnel handles request to get a list of Personnel entities
// @Summary List personnel entities
// @Description list Personnel entities
// @ID list-personnel
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Personnel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personnels [get]
func (ctl *PersonnelController) ListPersonnel(c *gin.Context) {
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

	personnel, err := ctl.client.Personnel.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, personnel)
}

// DeletePersonnel handles DELETE requests to delete a Personnel entity
// @Summary Delete a personnel entity by ID
// @Description get personnel by ID
// @ID delete-personnel
// @Produce  json
// @Param id path int true "Personnel ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personnels/{id} [delete]
func (ctl *PersonnelController) DeletePersonnel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Personnel.
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

// UpdatePersonnel handles PUT requests to update a Personnel entity
// @Summary Update a personnel entity by ID
// @Description update personnel by ID
// @ID update-personnel
// @Accept   json
// @Produce  json
// @Param id path int true "Personnel ID"
// @Param Personnel body ent.Personnel true "Personnel entity"
// @Success 200 {object} ent.Personnel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /personnels/{id} [put]
func (ctl *PersonnelController) UpdatePersonnel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Personnel{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Personnel binding failed",
		})
		return
	}
	obj.ID = int(id)
	personnel, err := ctl.client.Personnel.
		UpdateOneID(int(id)).
		SetName(obj.Name).
		SetDepartment(obj.Department).
		SetUser(obj.User).
		SetPassword(obj.Password).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, personnel)
}

// NewPersonnelController creates and registers handles for the PersonnelController
func NewPersonnelController(router gin.IRouter, client *ent.Client) *PersonnelController {
	personnelController := &PersonnelController{
		client: client,
		router: router,
	}
	personnelController.register()
	return personnelController
}

// InitPersonnelController registers routes to the main engine
func (ctl *PersonnelController) register() {
	Personnel := ctl.router.Group("/personnels")
	Personnel.GET("", ctl.ListPersonnel)
	// CRUD
	Personnel.POST("", ctl.CreatePersonnel)
	Personnel.GET(":id", ctl.GetPersonnel)
	Personnel.PUT(":id", ctl.UpdatePersonnel)
	Personnel.DELETE(":id", ctl.DeletePersonnel)
}
