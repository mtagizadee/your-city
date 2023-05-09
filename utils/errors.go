package utils

import (
	"net/http"
	"strings"
	"your-city/packages/common"
)

func DefaultError(err error) *common.ErrorType {
	return &common.ErrorType{Status: http.StatusInternalServerError, Message: err.Error()}
}

func IsUniqueKeyError(err error) bool {
	return strings.Contains(err.Error(), "Duplicate entry")
}