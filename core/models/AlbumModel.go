package models

import (
	"time"

	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Title       string      `json:"title"` //şarkı adı
	Year        *time.Time  `json:"year" ` //süresi
	SongAtAlbum SongAtAlbum `gorm:"foreignkey:Album_ID"references:Id`
}
