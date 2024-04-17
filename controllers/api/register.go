package controllers_api

import "github.com/gofiber/fiber/v3"

type Userdata struct {
	Username string
	Password string
}

func Register(c fiber.Ctx) error {
	var userdata *Userdata = new(Userdata)
	if err := c.BodyParser(&userdata); err != nil {
		return c.SendString("error in params")
	}
	c.SendString("ok")
}
