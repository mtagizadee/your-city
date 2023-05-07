package main

import (
	"fmt"
	"your-city/packages/auth"
	"your-city/packages/config"
	"your-city/packages/db"
	"your-city/packages/migration"
	"your-city/packages/users"

	"github.com/gin-gonic/gin"
)

func main() {
  config.LoadConfig()
  serverConfig := config.GetServerConfig()
  
  // connect to database and migrate it
  db.Connect()
  migration.Migrate()

  router := gin.Default()
  
  // auth routes
  authController := new(auth.AuthController)
  authController.AssignRoutes(router)

  // users routes
  usersController := new(users.UserController)
  usersController.AssignRoutes(router)

  router.Run(fmt.Sprintf("%v:%v",serverConfig.Host,serverConfig.Port))
}
