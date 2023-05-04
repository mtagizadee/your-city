package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "your-city/packages/config"
)

func main() {
  config.LoadConfig()

  server := config.GetServerConfig()

  router := gin.Default()
  router.Run(fmt.Sprintf("%v:%v",server.Host,server.Port))
}
