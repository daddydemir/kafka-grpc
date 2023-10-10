package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"producerProject/config"
)

var DB *gorm.DB

func GetConnect() {
	dsn := config.GetEnv("POSTGRE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	DB = db
}
