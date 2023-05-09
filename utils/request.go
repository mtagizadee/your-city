package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"your-city/packages/common"

	"github.com/gin-gonic/gin"
)

func ValidateBody[K comparable](c *gin.Context) (*K, *common.ErrorType)  {
	dto := new(K)
	
	if err := c.ShouldBindJSON(dto); err != nil {
    return nil, &common.ErrorType{Status: http.StatusBadGateway, Message: err.Error()}
  }

	return dto, nil
}

func ValidateParam(param string, c *gin.Context) (*string, *common.ErrorType) {
	p := c.Param("id") 

  if p == "" {
    return nil, &common.ErrorType{Status: http.StatusBadRequest, Message: fmt.Sprintf("%v is not provided", param)}
  }

	return &p, nil
}

func ValidateId(c *gin.Context) (int, *common.ErrorType) {
	id, err1 := ValidateParam("id", c)	
	if err1 != nil {
		return 0, err1
	}

  nId, err := strconv.Atoi(*id)
  if err != nil {
    return 0, &common.ErrorType{Status: http.StatusBadRequest, Message: "invalid id, must be an integer"}
  }

	return nId, nil
}