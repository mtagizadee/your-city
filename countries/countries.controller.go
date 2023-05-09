package countries

import (
	"net/http"
	"your-city/packages/utils"

	"github.com/gin-gonic/gin"
)

type createCountryDto struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"max=3"`
}

type CountriesController struct{}

var service = new(coutriesService)

func (controller *CountriesController) getAll(c *gin.Context) {
	code := c.Query("code")

	countries, err := service.getAll(code)
	if utils.SendError(err, c) { return }

	c.IndentedJSON(http.StatusOK, countries)
}

func (controller *CountriesController) create(c *gin.Context) {
	dto, err := utils.ValidateBody[createCountryDto](c)
	if utils.SendError(err, c) { return }

	country, err := service.create(dto)
	if utils.SendError(err, c) {return }
	
	c.IndentedJSON(http.StatusCreated, country)
}

func (controller *CountriesController) AssignRoutes(router *gin.Engine) {
	users := router.Group("/countries")

  users.GET("/", controller.getAll)
  users.POST("/", controller.create)
}
