package profiles

import (
	"github.com/gofiber/fiber/v2"
)

func GetStudentProfile(c *fiber.Ctx) error {
	
	// return c.Status(200).JSON(response)
	return c.SendString("getting response in get student profile controller")
}

func PatchStudentProfile(c *fiber.Ctx) error {
	return c.SendString("response from patch student profile")
}