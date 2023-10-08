package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FacialAnalyse struct {
	gorm.Model
	AnalyseId uuid.UUID
	ImageId   uuid.UUID
	Name      string
	Value     string
	Rate      float32
}
