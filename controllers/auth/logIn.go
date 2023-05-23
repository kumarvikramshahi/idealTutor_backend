package auth

import (
	"github.com/gofiber/fiber/v2"
)

func LogIn(c *fiber.Ctx) error {
	return c.SendString("in LogIn controllerr")
}