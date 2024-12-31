package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE,PATCH, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

const (
	DRIVER    = "mysql"
	MYSQLPORT = "3306"
	HOST      = "localhost"
	USER      = "root"
	PASSWORD  = "password"
	DBNAME    = "demo_project"
)

var db *sqlx.DB
var err error

func ConnectDatabase() {
	// connectionString := "USER:PASSWORD@tcp(HOST:MYSQLPORT)/DBNAME"
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASSWORD, HOST, MYSQLPORT, DBNAME)
	db, err = sqlx.Open(DRIVER, connectionStr)
	if err != nil {
		log.Fatal(err)
		return
	}
}
func main() {
	ConnectDatabase()
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/user/create/", HandleRegistration)
	// router.POST("/user/create/", CreateUserHandler)
	// router.GET("/user/list/", ListUser)
	// router.DELETE("/user/delete/:id", DeleteUser)
	// router.GET("/user/get/:id", GetDetails)
	// router.PATCH("/user/update/:id", UpdateUser)
	router.Run(":8000")
}
