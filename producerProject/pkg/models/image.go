package models

import (
	"github.com/google/uuid"
)

type Image struct {
	ImageId   uuid.UUID `gorm:"type:uuid;primaryKey"`
	ImagePath string
	Status    string
}
