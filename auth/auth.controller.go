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

type AuthController struct {}

var service authService

func (controller *AuthController) Signup(c *gin.Context) {
	dto, err := utils.ValidateBody[createUserDto](c)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})    
    return
  }

  user, err := service.Signup(dto)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.IndentedJSON(http.StatusCreated, user)
}

func (controller* AuthController) AssignRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/signup", controller.Signup)
}