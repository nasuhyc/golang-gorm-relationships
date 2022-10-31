package controllers

import (
	"golang-gorm-relationships/config"
	"golang-gorm-relationships/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserIndex(c *fiber.Ctx) error {

	user := new([]models.User)
	config.DB.Find(&user)
	return c.Status(200).JSON(fiber.Map{"datas": user})
}
func UserStore(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	user := models.User{
		Name:     data["name"],
		Surname:  data["surname"],
		Phone:    data["phone"],
		Email:    data["email"],
		Password: password,
	}
	config.DB.Create(&user)
	return c.Status(200).JSON(fiber.Map{"data": user, "message": "Create user successfully."})

}
func UserUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	config.DB.Find(&user, id)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	config.DB.Save(&user)
	return c.JSON(fiber.Map{"data": &user, "message": "Updated is Successfully."})
}
func UserGetId(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	config.DB.Find(&user, id)
	return c.Status(200).JSON(&user)
}
func UserDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	config.DB.Find(&user, id).Delete(&user)
	return c.Status(200).JSON(fiber.Map{"data": &user, "message": "Deleted is Successfully."})

}
