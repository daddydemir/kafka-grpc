package service

import (
	"consumerProject/pkg/database"
	"consumerProject/pkg/models"
)

func (c *ConsumerService) SaveAnalyse(m models.FacialAnalyse) {
	database.DB.Save(&m)
}

func (c *ConsumerService) UpdateImageStatus(id, text string) {
	database.DB.Model(&models.Image{}).Where("image_id = ?", id).Update("status", text)
}
