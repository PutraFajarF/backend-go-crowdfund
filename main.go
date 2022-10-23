package main

import (
	"go-crowdfunding/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// dsn := "host=localhost user=postgres password=ktl123 dbname=go_crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Connection to database is good!")

	// var users []user.User
	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("================")
	// }

	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
	dsn := "host=localhost user=postgres password=ktl123 dbname=go_crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
