package service

import (
	"makromusic/pkg/database"
	"makromusic/pkg/models"
)

func GetImageByID(ID string) *models.FacialAnalyse {
	var analyse = new(models.FacialAnalyse)
	database.DB.Find(analyse, "image_id = ?", ID)
	return analyse
}

func UpdateAnalyse(analyse models.FacialAnalyse) {
	database.DB.Model(&models.FacialAnalyse{}).Where("image_id = ?", analyse.ImageID).Save(&analyse)
}

func Pagination() { // todo date?
	database.DB.Order("(joy + sorrow + anger + surprise + under_exposed + blurred + headwear ) / 7 DESC").
		Offset(0).
		Limit(10).
		Find(&[]models.FacialAnalyse{})
}
