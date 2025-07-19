package handlers

import (
    "context"
	"log"

    "github.com/gofiber/fiber/v2"
    "github.com/georgysavva/scany/v2/pgxscan"

    "la-cipollina-budgeter-api/db"
    "la-cipollina-budgeter-api/models"
)

func GetValentinos(c *fiber.Ctx) error {
    var valentinos []models.Valentino
    err := pgxscan.Select(
		context.Background(),
		db.Pool,
		&valentinos,
		`SELECT id, name, weekly_pay, created_at, updated_at
FROM valentinos ORDER BY created_at`,
	)
    if err != nil {
		log.Print("Error in GetValentinos: ", err)
        return c.Status(404).JSON(fiber.Map{"error": "Not sure why this is an error"})
    }
    return c.JSON(valentinos)
}

func AddValentino(c *fiber.Ctx) error {
	var valentino models.Valentino
	if err := c.BodyParser(&valentino); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := db.Pool.QueryRow(
		context.Background(),
		`INSERT INTO valentinos (name, weekly_pay) VALUES (
	$1, $2
) RETURNING id, created_at, updated_at`,
		valentino.Name,
		valentino.WeeklyPay,
	).Scan(
		&valentino.ID,
		&valentino.CreatedAt,
		&valentino.UpdatedAt,
	)
    if err != nil {
		log.Print("Error in AddValentino: ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Database error while adding valentino-type weekly-paid entity"})
    }
	return c.Status(201).JSON(valentino)
}

func UpdateValentino(c *fiber.Ctx) error {
	var valentino models.Valentino
	if err := c.BodyParser(&valentino); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := db.Pool.QueryRow(
		context.Background(),
		`UPDATE valentinos SET name = $2,
weekly_pay = $3, updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at`,
		c.Params("id"),
		valentino.Name,
		valentino.WeeklyPay,
	).Scan(&valentino.ID, &valentino.CreatedAt, &valentino.UpdatedAt)
    if err != nil {
		log.Print("Error in UpdateValentino: ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Database error while updating valentino-type weekly-paid entity"})
    }
	return c.Status(200).JSON(valentino)
}

func RemoveValentino(c *fiber.Ctx) error {
	_, err := db.Pool.Exec(
		context.Background(),
		`DELETE FROM valentinos WHERE id = $1`,
		c.Params("id"),
	)
    if err != nil {
		log.Print("Error in RemoveValentino: ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Database error while removing valentino-type weekly-paid entity"})
    }
	return c.Status(200).JSON(fiber.Map{"success": true})
}
