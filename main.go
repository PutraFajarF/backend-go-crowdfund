package main

import (
	"fmt"
	"go-crowdfunding/handler"
	"go-crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=ktl123 dbname=go_crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	input := user.LoginInput{
		Email:    "putra@gmail.com",
		Password: "password",
	}

	user, err := userService.Login(input)
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	// api versioning
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}
