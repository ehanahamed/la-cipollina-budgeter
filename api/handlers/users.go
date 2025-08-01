package handlers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"strconv"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gofiber/fiber/v2"

	"la-cipollina-budgeter-api/db"
	"la-cipollina-budgeter-api/models"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	err := pgxscan.Select(
		context.Background(),
		db.Pool,
		&users,
		`SELECT id, username, admin, created_at, updated_at
FROM auth.users ORDER BY created_at`,
	)
	if err != nil {
		log.Print("Error in GetUsers: ", err)
		return c.Status(404).JSON(fiber.Map{"error": "Database error while getting users"})
	}
	return c.JSON(users)
}

/* authorization logic is in auth/admin_middleware.go,
only admin users can use AddUser */
func AddUser(c *fiber.Ctx) error {
	var newUser models.UserInput
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	var user models.User
	err := db.Pool.QueryRow(
		context.Background(),
		`INSERT INTO auth.users (username, encrypted_password, admin) VALUES (
	$1, crypt($2, gen_salt('bf')), $3
) RETURNING id, username, created_at, updated_at`,
		*newUser.Username,
		*newUser.NewPassword,
		*newUser.Admin,
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

/*
	authorization logic is HERE, in this handler
	
	admin users can update any user
	non-admin users can ONLY update themselves
	and non-admin users can NOT update the `admin` field
*/
func UpdateUser(c *fiber.Ctx) error {
	authedUser := c.Locals("user").(models.User)

	/* only allow admins or users updating themselves */
	if *authedUser.Admin || strconv.Itoa(authedUser.ID) == c.Params("id") {
		var input models.UserInput
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

			/* when we're updating password,
			sign user out on all devices
			by deleting all sessions for that user */
			_, err := db.Pool.Exec(
				context.Background(),
				`DELETE FROM auth.sessions
WHER	E user_id = $1`,
				c.Params("id"),
			)
			if err != nil {
				log.Print("Error deleting sessions during password update: ", err)
			}
		}

		/* ONLY admin users can update the admin field */
		if input.Admin != nil && *authedUser.Admin {
			setParts = append(setParts, fmt.Sprintf("admin = $%d", argNum))
			args = append(args, *input.Admin)
			argNum++
		}

		// Always update updated_at
		setParts = append(setParts, "updated_at = now()")

		// Add WHERE clause
		args = append(args, c.Params("id"))
		query := fmt.Sprintf(
			`UPDATE auth.users
SET 	%s
WHER	E id = $%d
RETU	RNING id, username, admin, created_at, updated_at`,
			strings.Join(setParts, ", "),
			argNum,
		)

		var user models.User
		err := db.Pool.QueryRow(context.Background(), query, args...).Scan(
			&user.ID, &user.Username, &user.Admin, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			log.Print("Error in UpdateUser: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "Database error while updating user"})
		}

		return c.Status(200).JSON(user)
	} else {
		return c.Status(403).JSON(fiber.Map{
			"error": "Non-admin users can only update their own username or password",
		})
	}
}

/* authorization logic is in auth/admin_middleware.go,
only admin users can use DeleteUser */
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
			`INSERT INTO auth.users (username, encrypted_password, admin)
VALUES ('admin', crypt('admin', gen_salt('bf')), true)`,
		)
		if err != nil {
			log.Print("Error inserting new default admin user: ", err)
		}
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}
