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

func AddRecordRepo(c *gin.Context, req AddUserEntity) (*ResponseStruct, string) {
	query := "INSERT INTO records (name, email, phone, message) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, req.Name, req.Email, req.Phone, req.Message)
	if err != nil {
		return nil, "Internal Server Error"
	}
	userData := ResponseStruct{
		Code:    200,
		Massage: "User Added Successfully.",
	}
	return &userData, ""
}

func UserListRepo(c *gin.Context) ([]ListResEntity, string) {
	tempModels := []ListResModal{}
	query := "SELECT * FROM records"
	err = db.Select(&tempModels, query)
	if err != nil {
		return nil, "Internal Server Error."
	}
	response := []ListResEntity{}
	for _, tempModel := range tempModels {
		userEntity := ConvertUserModalToUserEntity(tempModel)
		response = append(response, userEntity)
	}
	return response, ""
}

func DeleteUserRepo(c *gin.Context, userID int) (*ResponseStructEntity, string) {
	query := "DELETE FROM records WHERE id=?"
	_, err = db.Exec(query, userID)
	if err != nil {
		return nil, "Internal Server Error."
	}
	res := ResponseStructEntity{
		Code:    200,
		Massage: "Record Deleted successfully",
	}
	return &res, ""
}

func GetUserDetailRepo(c *gin.Context, userID int) (*ListResEntity, string) {
	tempModel := ListResModal{}
	query := "SELECT * FROM records WHERE id=?"
	err = db.Get(&tempModel, query, userID)
	if err != nil {
		return nil, "Internal Server Error."
	}
	userEntity := ConvertUserModalToUserEntity(tempModel)
	return &userEntity, ""
}

func UpdateUserRepo(c *gin.Context, req AddUserEntity, userID int) string {
	query := "UPDATE records SET name=?,  email=?,phone=?,message=? WHERE id=?"
	_, err := db.Exec(query, req.Name, req.Email, req.Phone, req.Message, userID)

	if err != nil {
		return "Internal Server Error"
	}

	return ""
}
