package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Port string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8000"
	}

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Printf("Database connected successfully after %d attempt(s)", i+1)
			Migrate(DB)
			return
		}
		log.Printf("Failed to connect to database, retrying (%d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Error connecting to database after %d retries", maxRetries)

	Migrate(DB)
	log.Printf("Database connected")
}
