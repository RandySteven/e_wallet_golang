package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseHandler struct{}

func (rh *ResponseHandler) ResponseEncoder(c *gin.Context, statusCode int, status bool, dataName string, data interface{}) {
	response := make(map[string]interface{})
	response["status"] = status
	response["responseCode"] = statusCode
	response[dataName] = data
	c.JSON(statusCode, response)
}
