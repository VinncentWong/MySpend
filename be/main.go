package main

import (
	"log"
	"module/app/user/handler"
	"module/app/user/repository"
	"module/app/user/usecase"
	"module/infrastructure"
	"module/rest"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("can't load env properties, %s", err.Error())
		return
	}

	err = infrastructure.InitDb()
	if err != nil {
		log.Fatal("exception on database initialization")
		return
	}
	err = infrastructure.Migrate()
	if err != nil {
		log.Fatal("exception when migrating")
		return
	}

	// Routing Initialization
	r := gin.Default()

	// User App
	userRepo := repository.NewUserRepository()
	userService := usecase.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	rest.UserRouting(r, userHandler)

	r.Run()
}
