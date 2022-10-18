package main

import (
	"golang/handler"
	"golang/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/ap/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.ChekEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "test simpan dari services"
	// userInput.Email = "asda@gmail.com"
	// userInput.Occupation = "anak"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)
}
