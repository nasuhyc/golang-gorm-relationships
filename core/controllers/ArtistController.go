package controllers

import (
	"golang-gorm-relationships/config"
	"golang-gorm-relationships/models"

	"github.com/gofiber/fiber/v2"
)

func ArtistIndex(c *fiber.Ctx) error {
	var artists []models.Artist
	err := config.DB.Model(&models.Artist{}).Preload("Songs").Find(&artists).Error
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(artists)

}
func ArtistStore(c *fiber.Ctx) error {
	type FullData struct {
		ID      float64       `gorm:"primaryKey"`
		Name    string        `json:"name"`
		Surname string        `json:"surname"`
		Songs   []models.Song `json:"songs"`
	}

	fulldata := FullData{}
	if err := c.BodyParser(&fulldata); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// if err := c.BodyParser(&artist); err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }

	artist := models.Artist{
		Name:    fulldata.Name,
		Surname: fulldata.Surname,
		ID:      fulldata.ID,
		Songs:   fulldata.Songs,
	}
	config.DB.Omit("Song").Save(&artist)

	// artist = models.Artist{
	// 	Name:    artist.Name,
	// 	Surname: artist.Surname,
	// }
	// config.DB.Create(&artist)
	// config.DB.Model(&song).Association("Artists").Append(&artist)
	return c.Status(200).JSON(fiber.Map{"data": artist, "message": "Create artist successfully."})

	// var artistData map[string]string
	// if err := c.BodyParser(&artistData); err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }

	// artist := models.Artist{
	// 	Name:    artistData["name"],
	// 	Surname: artistData["surname"],
	// }
	// config.DB.Create(&artist)
	// return c.Status(200).JSON(fiber.Map{"data": artist, "message": "Create artist successfully."})

}
func ArtistUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var artist models.Artist
	config.DB.Find(&artist, id)
	if err := c.BodyParser(&artist); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	config.DB.Save(&artist)
	return c.JSON(fiber.Map{"data": &artist, "message": "Updated is Successfully."})
}
func ArtistGetId(c *fiber.Ctx) error {
	id := c.Params("id")
	var artist models.Artist
	config.DB.Find(&artist, id)
	return c.Status(200).JSON(&artist)
}
func ArtistDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	var artist models.Artist
	config.DB.Find(&artist, id).Delete(&artist)
	return c.Status(200).JSON(fiber.Map{"data": &artist, "message": "Deleted is Successfully."})

}
