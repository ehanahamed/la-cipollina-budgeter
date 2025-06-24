package handlers

import (
    "context"
	"log"
	"fmt"

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
	var newUser models.UserInput
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	var user models.User
	err := db.Pool.QueryRow(
		context.Background(),
		`INSERT INTO auth.users (username, encrypted_password) VALUES (
	$1, crypt($2, gen_salt('bf'))
) RETURNING id, username, created_at, updated_at`,
		*newUser.Username,
		*newUser.NewPassword,
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
	var updatedUser models.UserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	/* Send an error if they're not updating anything */
	if input.Username == nil && input.NewPassword == nil {
	    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No fields to update"})
	}

	/* Add to/build the query */
	setParts := []string{}
	args := []interface{}{}
	argNum := 1

	if input.Username != nil {
	    setParts = append(setParts, fmt.Sprintf("username = $%d", argNum))
	    args = append(args, *input.Username)
	    argNum++
	}

	if input.NewPassword != nil {
	    setParts = append(setParts, fmt.Sprintf("encrypted_password = crypt($%d, gen_salt('bf'))", argNum))
	    args = append(args, *input.NewPassword)
	    argNum++
	}

	// Always update updated_at
	setParts = append(setParts, "updated_at = now()")

	// Add WHERE clause
	args = append(args, c.Params("id"))
	query := fmt.Sprintf(
		`UPDATE auth.users
SET %s
WHERE id = $%d
RETURNING id, username, created_at, updated_at`,
		strings.Join(setParts, ", "),
		argNum,
	)

	var user models.User
	err := db.Pool.QueryRow(context.Background(), query, args...).Scan(
	    &user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
	    log.Print("Error in UpdateUser: ", err)
	    return c.Status(500).JSON(fiber.Map{"error": "Database error while updating user"})
	}

	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	ctx := context.Background()
	_, err := db.Pool.Exec(
		ctx,
		`DELETE FROM auth.users WHERE id = $1`,
		c.Params("id"),
	)
	if err != nil {
		log.Print("Error in DeleteUser: ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Database error while deleting user"})
    }

	/* after deleting user, check if there's no users and recreate the default admin user/login if there's no users/logins */
	var count int
	err = db.Pool.QueryRow(ctx, `SELECT COUNT(*) FROM auth.users`).Scan(&count)
	if err != nil {
		log.Print("Error checking user count after successfully deleting a user: ", err)
	}
	if count == 0 {
		_, err := db.Pool.Exec(
			ctx,
			`INSERT INTO auth.users (username, encrypted_password)
VALUES ('admin', crypt('admin', gen_salt('bf')))`,
		)
		if err != nil {
			log.Print("Error inserting new default admin user: ", err)
		}
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}
