package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/favicon"
)

func setup_favicon(app *fiber.App) {
	app.Use(favicon.New(favicon.Config{
		URL:  "/favicon.ico",
		File: "./statics/favicon/icon.ico",
	}))
}
