package main

import (
	"fmt"
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
	fmt.Println("==================>its in service")
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
