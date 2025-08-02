package auth

import (
	"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"

	"la-cipollina-budgeter-api/db"
	"la-cipollina-budgeter-api/models"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" || len(authHeader) < 7 || !strings.EqualFold(authHeader[:7], "Bearer ") {
		return c.Status(401).JSON(fiber.Map{"error": "Missing or invalid auth header"})
	}
	token := authHeader[7:] /* `Bearer ` is 7 characters */

	var user models.User
    err := db.Pool.QueryRow(
        context.Background(),
        `SELECT u.id, u.username, u.admin,
u.created_at, u.updated_at
FROM auth.sessions s
JOIN auth.users u ON s.user_id = u.id
WHERE s.token = $1
AND s.expire_at > now()`,
        token,
    ).Scan(
		&user.ID,
		&user.Username,
		&user.Admin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Print("Error in AuthMiddleware: ", err)
		return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	c.Locals("user", user)

	return c.Next()
}

func AuthOrShareLinkMiddlewareWDateParam(c *fiber.Ctx) error {
	shareLinkToken := c.Query("s")
	if shareLinkToken != "" {
		var isValid bool
		err := db.Pool.QueryRow(
			context.Background(),
			`SELECT EXISTS (
SELECT 1 FROM weeks
WHERE $1::date BETWEEN start_date AND end_date
AND share_link_token = $2
)`,
			c.Params("date"),
			shareLinkToken,
		).Scan(&isValid)

		if err == nil && isValid {
			return c.Next()
		}
	}

	return AuthMiddleware(c)
}
