package users

import (
	"net/http"
	"your-city/packages/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var service = new(usersService)

func (controller *UserController) getById(c *gin.Context)  {
  id, err := utils.ValidateId(c)
  if utils.SendError(err, c) { return }

  user, err := service.GetById(id)
  if utils.SendError(err, c) { return }

  c.IndentedJSON(http.StatusOK, user)
}

func (controller *UserController) AssignRoutes(router *gin.Engine) {
  users := router.Group("/users")

  users.GET("/:id", controller.getById)
}