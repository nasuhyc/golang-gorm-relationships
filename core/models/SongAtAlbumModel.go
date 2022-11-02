package models

import (
	"gorm.io/gorm"
)

type SongAtAlbum struct {
	gorm.Model
	Song_ID  int16 `json:"song_id"`   //şarkı adı
	Album_ID int16 `json:"album_id" ` //süresi

}
