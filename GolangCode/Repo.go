package main

import (
	"database/sql"
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
		ErrorRes(c, http.StatusInternalServerError, "Internal Server Error", nil)
	}
	return nil
}

func GetUserData(c *gin.Context, email string) (*UserEntity, error) {
	type tempModel struct {
		ID       int            `db:"id"`
		Name     sql.NullString `db:"name"`
		Email    sql.NullString `db:"email"`
		Password sql.NullString `db:"password"`
	}
	userModel := tempModel{}
	query := `SELECT * FROM demo_project.users WHERE users.email=?`
	err := db.Get(&userModel, query, email)
	if err != nil {
		return nil, err
	}
	userData := UserEntity{
		Name:     userModel.Name.String,
		Email:    userModel.Email.String,
		Password: userModel.Password.String,
	}
	return &userData, nil
}
