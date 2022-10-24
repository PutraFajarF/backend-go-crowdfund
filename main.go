package main

import (
	"go-crowdfunding/user"
	"log"

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

	userInput := user.RegisterUserInput{}
	userInput.Name = "Test Simpan dari service"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "geologist"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
