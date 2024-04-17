package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func setup_limiter(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Max:        4,
		Expiration: 10 * time.Second,
		LimitReached: func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"error": "limitted message",
			})
		},
	}))
}
