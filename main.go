package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/sawer-app/controller"
	"github.com/prayogatriady/sawer-app/db"
	"github.com/prayogatriady/sawer-app/repository"
	"github.com/prayogatriady/sawer-app/service"
)

func main() {

	// set environtment variable for for PORT
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("Environment variable PORT must be set")
	}

	// db, err := db.InitMySQL()
	db, err := db.InitPostgreSQL()
	if err != nil {
		log.Println(err)
	}

	userRepo := repository.NewUserRepository(db)
	if err := userRepo.CreateTables(); err != nil {
		log.Fatal(err)
	}

	userServ := service.NewUserService(userRepo)
	userCont := controller.NewUserController(userServ)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/signup", userCont.Signup)
	}

	log.Fatal(r.Run(":" + PORT))
}
