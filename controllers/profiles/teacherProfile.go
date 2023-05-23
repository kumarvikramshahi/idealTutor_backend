package profiles

import (
	"github.com/gofiber/fiber/v2"
)

func GetTeacherProfile(c *fiber.Ctx) error {
	
	// return c.Status(200).JSON(response)
	return c.SendString("getting response in get teacher profile controller")
}

func PatchTeacherProfile(c *fiber.Ctx) error {
	return c.SendString("response from patch teacher profile")
}