package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/sawer-app/controller"
	"github.com/prayogatriady/sawer-app/db"
	"github.com/prayogatriady/sawer-app/middleware"
	"github.com/prayogatriady/sawer-app/repository"
	"github.com/prayogatriady/sawer-app/service"
)

func main() {

	// set default environment
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9000"
	}
	DB_USER := os.Getenv("DB_USER")
	if DB_USER == "" {
		DB_USER = "root"
	}
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		DB_PASSWORD = ""
	}
	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST == "" {
		DB_HOST = "127.0.0.1"
	}
	DB_PORT := os.Getenv("DB_PORT")
	if DB_PORT == "" {
		DB_PORT = "3306"
	}
	DB_NAME := os.Getenv("DB_NAME")
	if DB_NAME == "" {
		DB_NAME = "sawer"
	}

	db, err := db.NewConnectDB(
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	).InitMySQL()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userServ := service.NewUserService(userRepo)
	userCont := controller.NewUserController(userServ)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/signup", userCont.Signup)
		api.POST("/signin", userCont.Signin)

		r.Use(middleware.AuthMiddleware)

		api.GET("/profile", userCont.Profile)
		api.PUT("/edit", userCont.EditProfile)
		api.DELETE("/delete", userCont.DeleteUser)
	}

	log.Fatal(r.Run(":" + PORT))
}
