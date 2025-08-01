package auth

import (
	"github.com/gofiber/fiber/v2"

	"la-cipollina-budgeter-api/models"
)

func AdminMiddleware(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "User context missing"})
	}

	if user.Admin == nil || !*user.Admin {
		return c.Status(403).JSON(fiber.Map{"error": "Admin permissions required"})
	}

	return c.Next()
}
