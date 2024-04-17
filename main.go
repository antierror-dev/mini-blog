package main

import (
	"myapp/config"
	"myapp/middleware"
	"myapp/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Custom config
	app := fiber.New(config.Appconfig)
	routes.SetUp(app)
	middleware.SetUp(app)

	// Custom config

	app.Listen("0.0.0.0:8000")
}
