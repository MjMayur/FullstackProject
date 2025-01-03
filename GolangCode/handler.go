package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRegistration(c *gin.Context) {
	var Request CreateUserRequest
	// Bind the JSON request to the request struct
	err := c.ShouldBindJSON(&Request)
	if err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid JSON format",
		})
		return
	}
	UserEntity := ConvertUserEntity(Request)

	if err := validationEmail(Request.Email); err != nil {
		log.Println("Email validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	// if err := validationPass(request.Password); err != nil {
	// 	log.Println("Password must be strong:", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"code":    400,
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	err = CreateUserService(c, UserEntity)
	if err != nil {
		log.Println("Password must be strong:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	SuccessRes(c, http.StatusOK, "Data fetched Successfully", nil)
}

func HandleLogin(c *gin.Context) {
	var request LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid JSON format",
		})
		return
	}

	if err := validationEmail(request.Email); err != nil {
		log.Println("Email validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	userEntity, errorRes := LoginWithEmailPassword(c, request)
	if errorRes != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": errorRes,
		})
		return
	}
	SuccessRes(c, http.StatusOK, "Data fetched Successfully", userEntity)

}
func CreateUserHandler(c *gin.Context) {
	var request AddUserRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"FAILED": err})
		return
	}
	if err := validationEmail(request.Email); err != nil {
		log.Println("Email validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return

	}
	addUser := ConvertAddUserEntity(request)
	_, err := AddRecordService(c, addUser)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": err,
		})
		return
	}
	SuccessRes(c, http.StatusOK, "Data fetched Successfully", nil)

}
