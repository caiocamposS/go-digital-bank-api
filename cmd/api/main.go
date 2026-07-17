package main

import (
	"digital-bank-api/internal/config"
	"digital-bank-api/internal/handlers"
	"digital-bank-api/internal/models"
	"digital-bank-api/internal/repository"
	"digital-bank-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := config.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepo)
	userHandler := handlers.NewUserHandler(*userService)

	router := gin.Default()

	router.POST("/cadastro", userHandler.Signup)
	router.POST("/login", userHandler.Login)

	router.Run(":8080")
}