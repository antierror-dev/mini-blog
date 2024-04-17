package routes

import (
	controllers_basics "myapp/controllers/basics"
	"time"

	"github.com/gofiber/fiber/v3"
)

func setup_basics(prefix string, app *fiber.App) {
	bs := app.Group(prefix)

	bs.Get("/", controllers_basics.Home)
	bs.Static("/public", "./statics/public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Download:      false,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})
}
