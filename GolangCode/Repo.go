package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserRepo(c *gin.Context, user UserEntity) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	fmt.Println()
	result, err := db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		res := ResponseStruct{
			Code:    500,
			Massage: "Internal Server Error.",
		}
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Query Exection Error": res})
	}
	_, err = result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Last Insert ID getting Error": err})
	}

	return nil
}
