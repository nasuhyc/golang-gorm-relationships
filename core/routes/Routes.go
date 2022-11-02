package routes

import (
	"golang-gorm-relationships/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api") // /api
	//Artist Group
	artist := api.Group("/artist")
	artist.Get("/gets", controllers.ArtistIndex)
	artist.Post("/", controllers.ArtistStore)
	artist.Post("/update/:id", controllers.ArtistUpdate)
	artist.Get("/:id", controllers.ArtistGetId)
	artist.Delete("/delete/:id", controllers.ArtistDestroy)
	//Artist Group End

	//Song Group
	song := api.Group("/song")
	song.Get("/gets", controllers.SongIndex)
	song.Post("/", controllers.SongStore)
	song.Post("/update/:id", controllers.SongUpdate)
	song.Get("/:id", controllers.SongGetId)
	song.Delete("/delete/:id", controllers.SongDestroy)
	//Song Group

}
