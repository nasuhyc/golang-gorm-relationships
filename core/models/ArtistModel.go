package models

import (
	"time"

	"gorm.io/gorm"
)

type Artist struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname" `
	Songs     []Song         `gorm:"many2many:artist_at_song" json:"songs" `
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
