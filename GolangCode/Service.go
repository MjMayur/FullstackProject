package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func LoginWithEmailPassword(c *gin.Context, loginReq LoginRequest) (*User, error) {
	userEntity, err := GetUserData(c, loginReq.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, err
	}
	response := User{
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
	return &response, nil
}
