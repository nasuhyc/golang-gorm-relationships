package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Title    string  `json:"title"`      //şarkı adı
	SongTime float64 `json:"song_time" ` //süresi

	Artists []*Artist `gorm:"many2many:artist_at_song;" `
}
