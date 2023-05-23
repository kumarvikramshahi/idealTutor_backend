package sessions

import (
	"github.com/gofiber/fiber/v2"
)

func PostTeachSession(c *fiber.Ctx) error {
	return c.SendString("in post teach session")
}

func PatchTeachSession(c *fiber.Ctx) error {
	return c.SendString("in patch teach session")
}

func FetchTeachSession(c *fiber.Ctx) error {
	return c.SendString("in fetch teach session")
}
