package main

import (
	"digital-bank-api/internal/config"
	"digital-bank-api/internal/handlers"
	"digital-bank-api/internal/middleware"
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
		&models.Account{},
		&models.Transaction{},
	)

	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepo)
	userHandler := handlers.NewUserHandler(*userService)

	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(*accountRepo)
	accountHandler := handlers.NewAccountHandler(*accountService)

	router := gin.Default()

	router.POST("/cadastro", userHandler.Signup)
	router.POST("/login", userHandler.Login)

	protected := router.Group("/users")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", userHandler.GetProfile)

	protected.POST("/account", accountHandler.CreateAccount)

	router.Run(":8080")
}