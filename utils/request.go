package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateBody[K comparable](c *gin.Context) (*K, error)  {
	dto := new(K)
	
	if err := c.ShouldBindJSON(dto); err != nil {
    return nil, err
  }

	return dto, nil
}

func ValidateParam(param string, c *gin.Context) (*string, error) {
	p := c.Param("id") 

  if p == "" {
    return nil, fmt.Errorf(fmt.Sprintf("%v is not provided", param)) 
  }

	return &p, nil
}

func ValidateId(c *gin.Context) (int, error) {
	id, err := ValidateParam("id", c)	
	if err != nil {
		return 0, err
	}

  nId, err := strconv.Atoi(*id)
  if err != nil {
    return 0, fmt.Errorf("invalid id, must be an integer")
  }

	return nId, nil
}