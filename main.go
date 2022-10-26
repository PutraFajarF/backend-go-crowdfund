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

	userByEmail, err := userRepository.FindByEmail("putsra@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	if userByEmail.ID == 0 {
		fmt.Println("User not found")
	} else {
		fmt.Println(userByEmail.Name)
	}

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	// api versioning
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}
