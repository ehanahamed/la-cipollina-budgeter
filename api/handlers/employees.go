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
		`SELECT id, name, type, wage, special_pay, created_at, updated_at
FROM employees ORDER BY created_at`,
	)
    if err != nil {
		log.Print("Error in GetEmployees: ", err)
        return c.Status(404).JSON(fiber.Map{"error": "Not sure why this is an error"})
    }
    return c.JSON(employees)
}

func AddEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := db.Pool.QueryRow(
		context.Background(),
		`INSERT INTO employees (name, wage, type, special_pay) VALUES (
	$1, $2, $3, $4
) RETURNING id, created_at, updated_at`,
		employee.Name,
		employee.Wage,
		employee.Type,
		employee.SpecialPay,
	).Scan(
		&employee.ID,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)
    if err != nil {
		log.Print("Error in AddEmployee: ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Database error while adding employee"})
    }
	return c.Status(201).JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := db.Pool.QueryRow(
		context.Background(),
		`UPDATE employees SET name = $2,
wage = $3, type = $4, special_pay = $5, updated_at = now()
WHERE id = $1
RETURNING created_at, updated_at`,
		employee.ID,
		employee.Name,
		employee.Wage,
		employee.Type,
		employee.SpecialPay,
	).Scan(&employee.CreatedAt, &employee.UpdatedAt)
    if err != nil {
		log.Print("Error in UpdateEmployee: ", err)
        return c.Status(500).JSON(fiber.Map{"error": "Database error while updating employee"})
    }
	return c.Status(200).JSON(employee)
}
