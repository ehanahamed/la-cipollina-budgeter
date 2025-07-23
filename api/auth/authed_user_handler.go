package auth

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"la-cipollina-budgeter-api/db"
	"la-cipollina-budgeter-api/models"
)

func GetAuthedUser(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	token := authHeader[7:] /* `Bearer ` prefix before actual token is 7 characters */

	var authedUser models.User
	err := db.Pool.QueryRow(
		context.Background(),
		`SELECT u.id, u.username, u.created_at, u.updated_at
FROM auth.sessions s
JOIN auth.users u ON s.user_id = u.id
WHERE s.token = $1`,
		token,
	).Scan(
		&authedUser.ID,
		&authedUser.Username,
		&authedUser.CreatedAt,
		&authedUser.UpdatedAt,
	)
	if err != nil {
		log.Print("Error in GetAuthedUser: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error when trying to fetch authed user"})
	}
	return c.JSON(authedUser)
}
