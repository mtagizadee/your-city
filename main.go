package main

import (
	"fmt"
	"your-city/packages/config"
	"your-city/packages/db"
	"your-city/packages/users"

	"github.com/gin-gonic/gin"
)

func main() {
  config.LoadConfig()

  serverConfig := config.GetServerConfig()
  
  db.Connect()
  db.Migrate()

  router := gin.Default()

  var usersController = new(users.UserController)

  router.POST("/users", usersController.GetById)
  router.Run(fmt.Sprintf("%v:%v",serverConfig.Host,serverConfig.Port))
}
