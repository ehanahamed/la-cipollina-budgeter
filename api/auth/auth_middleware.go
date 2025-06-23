package auth

import (
	"context"
	
    "la-cipollina-budgeter-api/db"
)

func AuthMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
	if authHeader == "" || len(authHeader) < 7 || !strings.EqualFold(authHeader[:7], "Bearer ") {
	    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid auth header"})
	}
	token := authHeader[7:] /* `Bearer ` is 7 characters */

	var userID int
	err := db.Pool.QueryRow(
		context.Background(),
		`SELECT user_id FROM auth.sessions
		WHERE token = $1 AND
		expire_at > (SELECT now())`,
		token,
	).Scan(&userID)
    if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			/* 0 rows means token isn't there or is expired */
        	return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error in auth middleware"})
    }

	/* store userID in context */
    c.Locals("userID", userID)

    return c.Next()
}
