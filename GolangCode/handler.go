package main

import (
	"log"
	"net/http"
	"strconv"

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

func (h RegistryHandler) CreateUserHandler(c *gin.Context) {
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
	_, err := h.ServiceInterface.AddRecordService(c, addUser)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": err,
		})
		return
	}
	SuccessRes(c, http.StatusOK, "Data fetched Successfully", nil)

}

func ListUser(c *gin.Context) {
	userList, err := UserListService(c)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": err,
		})
		return
	}
	response := []UserListRes{}
	for _, userListRes := range userList {
		userJson := ConvertUserEntityToUserJson(userListRes)
		response = append(response, userJson)
	}
	SuccessRes(c, http.StatusOK, "Data fetched Successfully", response)
}

func DeleteUser(c *gin.Context) {
	userIDstr := c.Param("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": err,
		})
		return
	}
	_, error := DeleteUserService(c, userID)
	if error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": error,
		})
		return
	}
	SuccessRes(c, http.StatusOK, "User Deleted Successfully.", nil)
}

func GetDetails(c *gin.Context) {
	userIDstr := c.Param("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": err,
		})
		return
	}
	userDetails, error := GetDetailsService(c, userID)
	if error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": error,
		})
		return
	}
	SuccessRes(c, http.StatusOK, "User Deleted Successfully.", userDetails)
}

func UpdateUser(c *gin.Context) {
	var request AddUserRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"FAILED": err})
		return
	}
	userIDstr := c.Param("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": err,
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
	addUser := ConvertAddUserEntity(request)
	error := UpdateUserService(c, addUser, userID)
	if error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": error,
		})
		return
	}
	SuccessRes(c, http.StatusOK, "Data fetched Successfully", nil)

}
