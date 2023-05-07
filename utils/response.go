package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(err error, c *gin.Context) bool {
	if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return true
  }

	return false // if the error was not found
}