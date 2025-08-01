package auth

import (
	"github.com/gofiber/fiber/v2"

	"la-cipollina-budgeter-api/models"
)

func GetAuthedUser(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	return c.JSON(user)
}
