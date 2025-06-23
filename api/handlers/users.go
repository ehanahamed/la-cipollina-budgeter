package handlers

import (
    "context"
	"log"

    "github.com/gofiber/fiber/v2"
    "github.com/georgysavva/scany/v2/pgxscan"

    "la-cipollina-budgeter-api/db"
    "la-cipollina-budgeter-api/models"
)

func GetUsers(c *fiber.Ctx) error {
    var users []models.User
    err := pgxscan.Select(
		context.Background(),
		db.Pool,
		&users,
		`SELECT id, username, created_at, updated_at
FROM auth.users ORDER BY created_at`,
	)
    if err != nil {
		log.Print("Error in GetUsers: ", err)
        return c.Status(404).JSON(fiber.Map{"error": "Database error while getting users"})
    }
    return c.JSON(users)
}

func AddUser(c *fiber.Ctx) error {
	var newUser models.NewUser
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	var user models.User
	err := db.Pool.QueryRow(
		context.Background(),
		`INSERT INTO auth.users (username, encrypted_password) VALUES (
	$1, crypt($2, gen_salt('bf'))
) RETURNING id, username, created_at, updated_at`,
		newUser.Username,
		newUser.NewPassword,
	).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
    if err != nil {
		log.Print("Error in AddUser: ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Database error when trying create user"})
    }
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	var updatedUser models.NewUser
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	var user models.User
	err := db.Pool.QueryRow(
		context.Background(),
		`UPDATE auth.users SET username = $2,
encrypted_password = $crypt($3, gen_salt('bf')),
updated_at = now()
WHERE id = $1
RETURNING id, username, created_at, updated_at`,
		c.Params("id"),
		updatedUser.Username,
		updatedUser.NewPassword,
	).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
    if err != nil {
		log.Print("Error in UpdateUser: ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Database error while updating user"})
    }
	return c.Status(200).JSON(user)
}
