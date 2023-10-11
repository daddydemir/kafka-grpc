package service

import (
	"producerProject/pkg/database"
	"producerProject/pkg/models"
)

func GetImageByID(ID string) *models.FacialAnalyse {
	var analyse = new(models.FacialAnalyse)
	database.DB.Find(analyse, "image_id = ?", ID)
	return analyse
}

func UpdateAnalyse(analyse models.FacialAnalyse) {
	database.DB.Model(&models.FacialAnalyse{}).Where("image_id = ?", analyse.ImageID).Save(&analyse)
}

func GetPaginatedResult(limit, page int) []models.FacialAnalyse {
	var result []models.FacialAnalyse
	database.DB.Scopes(database.NewPaginate(limit, page).PaginatedResult).Find(&result)
	return result
}
