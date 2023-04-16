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

	db, err := db.NewConnectDB(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
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

	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
