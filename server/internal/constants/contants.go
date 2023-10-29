package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var configured = false

func Config() {
	if configured == false {
		log.Println("Configuring environment variables...")

		// running local
		if os.Getenv("ENV") == "" {
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file", err.Error())
			}
		}

		// Load into constants
		DB_USER = os.Getenv("PG_USER")
		DB_PWD = os.Getenv("PG_PASSWORD")
		DB_NAME = os.Getenv("PG_DATABASE")
		DB_ADDR = os.Getenv("DB_ADDR")
		ES_ADDR = os.Getenv("ES_ADDR")
		JWT_SECRET = os.Getenv("JWT_SECRET")

		log.Println("Environment variables configured!")
		configured = true
	} else {
		log.Println("Environment variables already configured!")
	}
}

var DB_USER string
var DB_PWD string
var DB_NAME string
var DB_ADDR string
var ES_ADDR string
var JWT_SECRET string
var MAX_MULTIPART_SIZE int64 = 10 << 20 // 10 MiB
