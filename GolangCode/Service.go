package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type ServiceImplementation struct {
}

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
		return nil, "Internal Server error"
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
	fmt.Println(response, "----------->")
	return &response, ""
}

func (h ServiceImplementation) AddRecordService(c *gin.Context, req AddUserEntity) (*ResponseStruct, string) {
	res, err := AddRecordRepo(c, req)
	if err != "" {
		return nil, "Internal Server Error"
	}
	return res, ""
}

func (h ServiceImplementation) UserListService(c *gin.Context) ([]ListResEntity, string) {
	res, err := UserListRepo(c)
	if err != "" {
		return nil, "Internal Server Error"
	}
	return res, ""
}

func (h ServiceImplementation) DeleteUserService(c *gin.Context, userID int) (*ResponseStructEntity, string) {
	res, err := DeleteUserRepo(c, userID)
	if err != "" {
		return nil, "Internal Server Error."
	}
	return res, ""
}

func (h ServiceImplementation) GetDetailsService(c *gin.Context, userID int) (*ListResEntity, string) {
	userDetails, err := GetUserDetailRepo(c, userID)
	if err != "" {
		return nil, "Internal Server Error."
	}
	return userDetails, ""
}

func (h ServiceImplementation) UpdateUserService(c *gin.Context, req AddUserEntity, userID int) string {
	err := UpdateUserRepo(c, req, userID)
	if err != "" {
		return "Internal Server Error"
	}
	return ""
}
