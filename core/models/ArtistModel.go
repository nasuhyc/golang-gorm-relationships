package models

type Artist struct {
	ID      float64 `gorm:"primaryKey"`
	Name    string  `json:"name"`
	Surname string  `json:"surname" `
	Songs   []Song  `gorm:"many2many:artist_at_song" json:"songs" `
}
