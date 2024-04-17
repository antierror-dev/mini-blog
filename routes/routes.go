package routes

import "github.com/gofiber/fiber/v3"

func SetUp(app *fiber.App) {
	setup_basics("/", app)
	setup_api("/api", app)
}
