package main

import (
	"bwa_go/auth"
	"bwa_go/handler"
	"bwa_go/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	UserRepository := user.NewRepository(db)
	UserService := user.NewService(UserRepository)
	AuthService := auth.NewService()

	userHandler := handler.NewUserHandler(UserService, AuthService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/emailcheckers", userHandler.CheckEmail)
	api.POST("/avatar", userHandler.UploadAvatar)
	router.Run()
}

//input dari User
//handler, mapping input user => struct input
//service : melakukan mapping dari struct input ke struct user
//repository
//db
