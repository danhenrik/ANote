package constants

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func Config() {
	once.Do(func() {
		log.Println("Configuring environment variables...")

		// running local
		if os.Getenv("ENV") == "" {
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file", err.Error())
			}
		}

		// Load into constants
		ES_ADDR = os.Getenv("ES_ADDR")
		DB_ADDR = os.Getenv("DB_ADDR")
		DB_USER = os.Getenv("PG_USER")
		DB_PWD = os.Getenv("PG_PASSWORD")
		DB_NAME = os.Getenv("PG_DATABASE")

		log.Println("Environment variables configured!")
	})
}

var ES_ADDR string
var DB_ADDR string
var DB_USER string
var DB_PWD string
var DB_NAME string
