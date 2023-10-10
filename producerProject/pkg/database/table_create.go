package database

import (
	"log"
	"producerProject/pkg/models"
)

func CreateTables() {
	err := DB.AutoMigrate(&models.Image{}, &models.FacialAnalyse{})
	if err != nil {
		log.Println("error: ", err)
	}
}
