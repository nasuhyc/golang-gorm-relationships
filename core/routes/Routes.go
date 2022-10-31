package routes

import (
	"golang-gorm-relationships/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api") // /api
	//User Group
	user := api.Group("/user")
	user.Get("/gets", controllers.UserIndex)
	user.Post("/", controllers.UserStore)
	user.Post("/update/:id", controllers.UserUpdate)
	user.Get("/:id", controllers.UserGetId)
	user.Delete("/delete/:id", controllers.UserDestroy)
	//User Group End

}
