package users

import (
	"net/http"
	"your-city/packages/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var service = new(usersService)

func (controller *UserController) GetById(c *gin.Context)  {
  id, err := utils.ValidateId(c)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user, err := service.GetById(id)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.IndentedJSON(http.StatusOK, user)
}

func (controller *UserController) AssignRoutes(router *gin.Engine) {
  users := router.Group("/users")

  users.GET("/:id", controller.GetById)
}