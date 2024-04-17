package routes

import (
	controllers_api "myapp/controllers/api"

	"github.com/gofiber/fiber/v3"
)

func setup_api(prefix string, app *fiber.App) {
	api := app.Group(prefix)

	api.Get("/", controllers_api.Home)
	api.Post("/register", controllers_api.Register)

}
