package controllers

import (
	"golang-gorm-relationships/config"
	"golang-gorm-relationships/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//DESC - Get all artists and their songs
func ArtistIndex(c *fiber.Ctx) error {
	var artists []models.Artist
	err := config.DB.Model(&models.Artist{}).Preload("Songs").Find(&artists).Error
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(artists)

}

//DESC - Create a new artist and add songs to it
func ArtistStore(c *fiber.Ctx) error {
	type FullData struct {
		ID      uint          `gorm:"primaryKey"`
		Name    string        `json:"name"`
		Surname string        `json:"surname"`
		Songs   []models.Song `json:"songs"`
	}

	fulldata := FullData{}
	if err := c.BodyParser(&fulldata); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	artist := models.Artist{
		Name:    fulldata.Name,
		Surname: fulldata.Surname,
		ID:      fulldata.ID,
		Songs:   fulldata.Songs,
	}

	config.DB.Omit("Song").Save(&artist)
	return c.Status(200).JSON(fiber.Map{"data": artist, "message": "Create artist successfully."})

}

//DESC - Update a artist and replace songs to it
func ArtistUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	U, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		return err
	}
	newid := uint(U)
	var client_query *models.Artist
	config.DB.Where("id = ?", id).Preload("Songs").First(&client_query)
	config.DB.Model(&models.Artist{ID: newid}).Association("Songs").Replace(&client_query.Songs)
	return c.Status(200).JSON(fiber.Map{"data": &client_query, "message": "Updated is Successfully."})
}

//DESC - Bringing artists and songs according to ID

func ArtistGetId(c *fiber.Ctx) error {
	id := c.Params("id")
	var artists []models.Artist
	err := config.DB.Model(&models.Artist{}).Preload("Songs").Find(&artists, id).Error
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(artists)

}

//DESC - Delete a artist and delete songs to it
func ArtistDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	U, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		return err
	}
	newid := uint(U)
	var client_query *models.Artist
	config.DB.Where("id = ?", id).Preload("Songs").First(&client_query).Delete(client_query)
	config.DB.Model(&models.Artist{ID: newid}).Association("Songs").Delete(&client_query.Songs)
	return c.Status(200).JSON(fiber.Map{"data": &client_query, "message": "Deleted is Successfully."})

}
