package controllers

import (
	"golang-gorm-relationships/config"
	"golang-gorm-relationships/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//DESC - Get all songs and their artists

func SongIndex(c *fiber.Ctx) error {
	var songs []models.Song
	err := config.DB.Model(&models.Song{}).Preload("Artists").Find(&songs).Error
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(songs)

}

//DESC - Create a new song and add artists to it
func SongStore(c *fiber.Ctx) error {

	type SongData struct {
		ID       uint            `gorm:"primaryKey"`
		Title    string          `json:"title"`
		SongTime float64         `json:"song_time"`
		Artists  []models.Artist `json:"artists"`
	}

	songData := SongData{}

	if err := c.BodyParser(&songData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	song := models.Song{
		Title:    songData.Title,
		SongTime: songData.SongTime,
		ID:       songData.ID,
		Artists:  songData.Artists,
	}
	config.DB.Omit("Artist").Save(&song)

	return c.Status(200).JSON(fiber.Map{"data": song, "message": "Create song successfully."})
}

//DESC - Update a song and replace artists to it
func SongUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	U, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		return err
	}
	newid := uint(U)
	var client_query *models.Song
	config.DB.Where("id = ?", id).Preload("Artists").First(&client_query)
	config.DB.Model(&models.Song{ID: newid}).Association("Artists").Replace(&client_query.Artists)
	return c.Status(200).JSON(fiber.Map{"data": &client_query, "message": "Updated is Successfully."})
}

//DESC - Bringing songs and artists according to ID

func SongGetId(c *fiber.Ctx) error {
	id := c.Params("id")
	var songs []models.Song
	err := config.DB.Model(&models.Artist{}).Preload("Songs").Find(&songs, id).Error
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(songs)
}

//DESC - Delete a song and delete artists to it
func SongDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	U, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		return err
	}
	newid := uint(U)
	var client_query *models.Song
	config.DB.Where("id = ?", id).Preload("Artists").First(&client_query).Delete(client_query)
	config.DB.Model(&models.Song{ID: newid}).Association("Artists").Delete(&client_query.Artists)
	return c.Status(200).JSON(fiber.Map{"data": &client_query, "message": "Deleted is Successfully."})
}
