package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Image struct {
	gorm.Model
	ImageId    uuid.UUID
	ImagePath  string
	Status     string
	CreateDate time.Time
	UpdateDate time.Time
	Analyses   []FacialAnalyse
}
