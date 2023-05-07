package auth

import (
	"net/http"
	"your-city/packages/utils"

	"github.com/gin-gonic/gin"
)

type createUserDto struct {
  Name string `json:"name" binding:"required"`
  Surname string `json:"surname" binding:"required"`
  Email string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required"`
}

type LoginUserDto struct {
  Email string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required"`
}

type AuthController struct {}

var service authService

func (controller *AuthController) signup(c *gin.Context) {
	dto, err := utils.ValidateBody[createUserDto](c)
  if utils.SendError(err, c) { return }

  user, err := service.signup(dto)
  if utils.SendError(err, c) { return }

  c.IndentedJSON(http.StatusCreated, user)
}

func (controller *AuthController) login(c *gin.Context) {
  dto, err := utils.ValidateBody[LoginUserDto](c)
  if utils.SendError(err, c) { return }

  user, token, err := service.login(dto)
  if utils.SendError(err, c) { return }

  c.Header("Authorization", token)
  c.IndentedJSON(http.StatusOK, user)
}

func (controller* AuthController) AssignRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/signup", controller.signup)
  auth.POST("/login", controller.login)
}