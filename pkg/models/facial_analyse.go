package models

import (
	"github.com/google/uuid"
)

type FacialAnalyse struct {
	ImageID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Joy          int
	Sorrow       int
	Anger        int
	Surprise     int
	UnderExposed int
	Blurred      int
	Headwear     int
	Confidence   float64
}
