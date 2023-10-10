package database

import (
	"consumerProject/pkg/models"
	"log"
)

func CreateTables() {
	err := DB.AutoMigrate(&models.Image{}, &models.FacialAnalyse{})
	if err != nil {
		log.Println("error: ", err)
	}
}
