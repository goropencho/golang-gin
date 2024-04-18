package main

import (
	"time"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/goropencho/golang-gin/controller"
	"github.com/goropencho/golang-gin/database"
	"github.com/goropencho/golang-gin/helper"
	"github.com/goropencho/golang-gin/model/users"
	repository "github.com/goropencho/golang-gin/repository/users"
	"github.com/goropencho/golang-gin/router"
	"github.com/goropencho/golang-gin/service"
)

func main() {

	db := database.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&users.User{})

	userRepository := repository.NewUsersRepositoryImpl(db)

	userService := service.NewUserServiceImpl(userRepository, validate)

	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
