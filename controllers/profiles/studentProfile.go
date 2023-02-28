package profiles

import (
	"github.com/gofiber/fiber/v2"
)

func GetStudentProfile(c *fiber.Ctx) error {
	
	// return c.Status(200).JSON(response)
	return c.SendString("getting response in profile controller")

}
