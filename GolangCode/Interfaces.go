package main

import "github.com/gin-gonic/gin"

type HandlerServices interface {
	AddRecordService(c *gin.Context, req AddUserEntity) (*ResponseStruct, string)
	UserListService(c *gin.Context) ([]ListResEntity, string)
	DeleteUserService(c *gin.Context, userID int) (*ResponseStructEntity, string)
	GetDetailsService(c *gin.Context, userID int) (*ListResEntity, string)
	UpdateUserService(c *gin.Context, req AddUserEntity, userID int) string
}
