package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(c *gin.Context, user UserEntity) error {
	p, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
	}

	UserEntity := UserEntity{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(p),
	}
	err = CreateUserRepo(c, UserEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return err
	}
	return nil
}

func LoginWithEmailPassword(c *gin.Context, loginReq LoginRequest) (*User, string) {
	userEntity, err := GetUserData(c, loginReq.Email)
	if err != nil {
		return nil, ""
	}
	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, "Invalid Credentials"
	}

	token := uuid.New().String()
	response := User{
		Name:  userEntity.Name,
		Email: userEntity.Email,
		Token: token,
	}
	return &response, ""
}

func AddRecordService(c *gin.Context, req AddUserEntity) (*ResponseStruct, string) {
	res, err := AddRecordRepo(c, req)
	if err != "" {
		return nil, "Internal Server Error"
	}
	return res, ""
}
