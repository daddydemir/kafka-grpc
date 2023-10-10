package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load("../config/local.env")
	if err != nil {
		log.Fatal("Error load .env file :", err)
	}
	return os.Getenv(key)
}
