package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var (
	R2_ENDPOINT string
	BUCKET_NAME string
	DEST_PATH   string
	ACCESS_KEY  string
	SECRET_KEY  string
)

func Init() {
	LoadEnv()
	R2_ENDPOINT = os.Getenv("R2_ENDPOINT")
	BUCKET_NAME = os.Getenv("BUCKET_NAME")
	DEST_PATH = os.Getenv("DEST_PATH")
	ACCESS_KEY = os.Getenv("ACCESS_KEY")
	SECRET_KEY = os.Getenv("SECRET_KEY")
}
