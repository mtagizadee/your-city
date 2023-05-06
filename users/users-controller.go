package users

import (
	"your-city/packages/users/dto"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var service usersService

func (uc *UserController) GetById(c *gin.Context) {
  // temporary implementation for checking functionality


  var d dto.CreateUserDto
  if err := c.BindJSON(&d); err != nil {
    return // handle error 
  }

  created, err := service.Create(&d)

  if err != nil {
    return // handle error
  }
  c.IndentedJSON(200, created)
}
