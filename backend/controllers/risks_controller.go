package controllers
 
import (
   "context"
   "strconv"

   "github.com/to63/app/ent"
   "github.com/to63/app/ent/risks"
   "github.com/gin-gonic/gin"
)
 
// RisksController defines the struct for the risks controller
type RisksController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateRisks handles POST requests for adding risks entities
// @Summary Create risks
// @Description Create risks
// @ID create-risks
// @Accept   json
// @Produce  json
// @Param risks body ent.Risks true "Risks entity"
// @Success 200 {object} ent.Risks
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Riskss [post]
func (ctl *RisksController) CreateRisks(c *gin.Context) {
	obj := ent.Risks{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "risks binding failed",
		})
		return
	}
  
	r, err := ctl.client.Risks.
		Create().
		SetRisks(obj.Risks).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, r)
 }
 // GetRisks handles GET requests to retrieve a risks entity
// @Summary Get a risks entity by ID
// @Description get risks by ID
// @ID get-risks
// @Produce  json
// @Param id path int true "Risks ID"
// @Success 200 {object} ent.Risks
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /riskss/{id} [get]
func (ctl *RisksController) GetRisks(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	r, err := ctl.client.Risks.
		Query().
		Where(risks.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, r)
 }
 // ListRisks handles request to get a list of risks entities
// @Summary List risks entities
// @Description list risks entities
// @ID list-risks
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Risks
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /riskss [get]
func (ctl *RisksController) ListRisks(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	riskss, err := ctl.client.Risks.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, riskss)
 }

 // NewRisksController creates and registers handles for the risks controller
func NewRisksController(router gin.IRouter, client *ent.Client) *RisksController {
	rc := &RisksController{
		client: client,
		router: router,
	}
	rc.register()
	return rc
 }
  
 // InitRisksController registers routes to the main engine
 func (ctl *RisksController) register() {
	riskss := ctl.router.Group("/riskss")
  
	riskss.GET("", ctl.ListRisks)
  
	// CRUD
	riskss.POST("", ctl.CreateRisks)
	riskss.GET(":id", ctl.GetRisks)
 }
 