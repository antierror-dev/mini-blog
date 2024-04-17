package middleware

import "github.com/gofiber/fiber/v3"

func SetUp(app *fiber.App) {
	setup_logger(app)
	setup_limiter(app)
	setup_favicon(app)
	//setup_monitor(app)
}
