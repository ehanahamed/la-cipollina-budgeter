package handlers

import (
    "context"
	"log"

    "github.com/gofiber/fiber/v2"
    "github.com/georgysavva/scany/v2/pgxscan"

    "la-cipollina-budgeter-api/db"
    "la-cipollina-budgeter-api/models"
)

func GetEmployees(c *fiber.Ctx) error {
    var employees []models.Employee
    err := pgxscan.Select(
		context.Background(),
		db.Pool,
		&employees,
		`SELECT id, type, wage, special_pay, created_at, updated_at FROM employees`,
	)
    if err != nil {
		log.Print("Error in GetEmployees: ", err)
        return c.Status(404).JSON(fiber.Map{"error": "Not sure why this is an error"})
    }
    return c.JSON(employees)
}

