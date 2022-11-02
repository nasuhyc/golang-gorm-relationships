package controllers

import (
	"golang-gorm-relationships/config"
	"golang-gorm-relationships/models"

	"github.com/gofiber/fiber/v2"
)

func SongIndex(c *fiber.Ctx) error {
	var songs []models.Song
	err := config.DB.Model(&models.Song{}).Preload("Artists").Find(&songs).Error
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(songs)

}

func SongStore(c *fiber.Ctx) error {
	song := models.Song{}
	artis := models.Artist{}
	if err := c.BodyParser(&song); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := c.BodyParser(&artis); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	song = models.Song{
		Title:    song.Title,
		SongTime: song.SongTime,
	}
	config.DB.Create(&song)
	config.DB.Model(&song).Association("Artists").Append(&artis)
	return c.Status(200).JSON(fiber.Map{"data": artis, "message": "Create artist successfully."})
}
func SongUpdate(c *fiber.Ctx) error {
	return c.SendString("sa")
}

func SongGetId(c *fiber.Ctx) error {
	return c.SendString("sa")
}

func SongDestroy(c *fiber.Ctx) error {
	return c.SendString("sa")
}
