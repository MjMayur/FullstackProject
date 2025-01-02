package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

const (
	STATUS_SUCCESS               = "SUCCESS"
	STATUS_FAILED                = "FAILED"
	STATUS_INTERNAL_SERVER_ERROR = "internal server error"
)

func SuccessRes(c *gin.Context, statusCode int, msg string, payload interface{}) {
	res := SuccessResponse{
		StatusCode: statusCode,
		Status:     STATUS_SUCCESS,
		Message:    msg,
		Data:       payload,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(res)
	c.Writer.WriteHeader(statusCode)
	c.Writer.Write(response)
}
func ErrorRes(c *gin.Context, statusCode int, msg string, payload interface{}) {
	res := SuccessResponse{
		StatusCode: statusCode,
		Status:     STATUS_SUCCESS,
		Message:    msg,
		Data:       payload,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(res)
	c.Writer.WriteHeader(statusCode)
	c.Writer.Write(response)
}
