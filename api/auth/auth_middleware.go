package auth

import (
	"context"
	
    "la-cipollina-budgeter-api/db"
)

func AuthMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
	if authHeader == "" || len(authHeader) < 7 || !strings.EqualFold(authHeader[:7], "Bearer ") {
	    return c.Status(401).JSON(fiber.Map{"error": "Missing or invalid auth header"})
	}
	token := authHeader[7:] /* `Bearer ` is 7 characters */

	var isTokenValid bool
	err := db.Pool.QueryRow(
		context.Background(),
		`SELECT EXISTS(
	SELECT 1 FROM auth.sessions
	WHERE token = $1 AND
	expire_at > (SELECT now())
)`,
		token,
	).Scan(&isTokenValid)
    if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error in auth middleware"})
    }
	if isTokenValid {
		/* if token is valid, allow user to use the route/endpoint's handler */
    	return c.Next()
	} else {
		/* if token isn't valid, return this error, which stops this route's handler from running */
        return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
	}
}
