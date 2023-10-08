package constants

import (
	"os"

	"github.com/joho/godotenv"
)

func Config() {
	godotenv.Load("../../.env")
}

var DB_USER = os.Getenv("PG_USER")
var DB_PWD = os.Getenv("PG_PASSWORD")
var DB_NAME = os.Getenv("PG_DATABASE")
var DB_ADDR = os.Getenv("DB_ADDR")
var JWT_SECRET = os.Getenv("JWT_SECRET")
