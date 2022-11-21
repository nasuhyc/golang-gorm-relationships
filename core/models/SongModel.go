package models

import (
	"time"

	"gorm.io/gorm"
)

type Song struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string         `json:"title"`      //şarkı adı
	SongTime  float64        `json:"song_time" ` //süresi
	Artists   []Artist       `gorm:"many2many:artist_at_song;" json:"artists" `
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
