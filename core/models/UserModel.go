package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `json:"id"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname" `
	Phone     string         `json:"phone"`
	Email     string         `gorm:"unique" json:"email"`
	Password  []byte         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
