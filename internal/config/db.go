package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connection with the database postgres
func ConnectToDB() (*gorm.DB, error) {
	var err error

	dsn := os.Getenv("DB")
	println("DSN:", dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	return DB, nil
}