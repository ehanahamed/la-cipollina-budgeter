package auth

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"la-cipollina-budgeter-api/db"
)

func LoginHandler(c *fiber.Ctx) error {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var genToken string
	err := db.Pool.QueryRow(
		context.Background(),
		`INSERT INTO auth.sessions (user_id)
SELECT id FROM auth.users
WHERE username = $1 AND encrypted_password = crypt($2, encrypted_password)
RETURNING token`,
		creds.Username,
		creds.Password,
	).Scan(&genToken)
	if err != nil {
		log.Print("Error in LoginHandler: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while logging in"})
	}
	return c.Status(200).JSON(fiber.Map{"token": genToken})
}
