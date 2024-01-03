package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var PORT string
var DATABASE_URL string
var JWT_SECRET_KEY string
var GOOGLE_KEY_ID string
var GOOGLE_SECRET_KEY string
var GITHUB_KEY_ID string
var GITHUB_SECRET_KEY string
var FACEBOOK_KEY_ID string
var FACEBOOK_SECRET_KEY string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	PORT = os.Getenv("PORT")
	DATABASE_URL = os.Getenv("DATABASE_URL")
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

	GOOGLE_KEY_ID = os.Getenv("GOOGLE_KEY_ID")
	GOOGLE_SECRET_KEY = os.Getenv("GOOGLE_SECRET_KEY")

	GITHUB_KEY_ID = os.Getenv("GITHUB_KEY_ID")
	GITHUB_SECRET_KEY = os.Getenv("GITHUB_SECRET_KEY")

	FACEBOOK_KEY_ID = os.Getenv("FACEBOOK_KEY_ID")
	FACEBOOK_SECRET_KEY = os.Getenv("FACEBOOK_SECRET_KEY")
}
