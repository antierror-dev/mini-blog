package controllers_api

import "github.com/gofiber/fiber/v3"

func Home(c fiber.Ctx) error {
	return c.SendString("hello world")
}
