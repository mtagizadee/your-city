package utils

import (
	"your-city/packages/common"

	"github.com/gin-gonic/gin"
)

func SendError(err *common.ErrorType, c *gin.Context) bool {
	if err != nil {
    c.JSON(err.Status, gin.H{"error": err.Message})
    return true
  }

	return false // if the error was not found
}