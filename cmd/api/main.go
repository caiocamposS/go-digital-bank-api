package main

import (
	"digital-bank-api/internal/config"
	"digital-bank-api/internal/models"
	"log"
	"net/http"

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

	router := gin.Default()
	
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}