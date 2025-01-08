package main

import "github.com/gin-gonic/gin"

type HandlerServices interface {
	AddRecordService(c *gin.Context, req AddUserEntity) (*ResponseStruct, string)
}
