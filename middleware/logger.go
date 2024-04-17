package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

var loggerconfig = logger.New(logger.Config{
	Format:     "[${ip}] ${status} - ${path} ${latency} (pid:${pid})\n",
	TimeFormat: "15:04:05",
	TimeZone:   "Local",
})

func setup_logger(app *fiber.App) {
	app.Use(loggerconfig)
}
