package handlers

import (
	"context"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gofiber/fiber/v2"

	"la-cipollina-budgeter-api/db"
	"la-cipollina-budgeter-api/models"
)

func GetDayByDate(c *fiber.Ctx) error {
	var day models.Day
	err := db.Pool.QueryRow(
		context.Background(),
		`SELECT id, date, hours_worked, worked_today,
current_employees, food_costs, created_at, updated_at
FROM days WHERE date = $1`,
		c.Params("Date"),
	).Scan(
		&day.ID,
		&day.Date,
		&day.HoursWorked,
		&day.WorkedToday,
		&day.CurrentEmployees,
		&day.FoodCosts,
		&day.CreatedAt,
		&day.UpdatedAt,
	)
	if err != nil {
		log.Print("Error in GetDayByDate: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while getting row from days by date"})
	}
	return c.JSON(day)
}

func RecordDay(c *fiber.Ctx) error {
	var day models.Day
	if err := c.BodyParser(&day); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := db.Pool.QueryRow(
		context.Background(),
		`INSERT INTO days (
	date, hours_worked, worked_today, current_employees, food_costs
) VALUES (
	$1, $2, $3, $4, $5
) RETURNING id, created_at, updated_at`,
		day.Date,
		day.HoursWorked,
		day.WorkedToday,
		day.CurrentEmployees,
		day.FoodCosts,
	).Scan(
		&day.ID,
		&day.CreatedAt,
		&day.UpdatedAt,
	)
	if err != nil {
		log.Print("Error in RecordDay: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while adding day"})
	}
	return c.Status(201).JSON(day)
}

func UpdateDay(c *fiber.Ctx) error {
	var day models.Day
	if err := c.BodyParser(&day); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := db.Pool.QueryRow(
		context.Background(),
		`UPDATE employees SET name = $2,
wage = $3, type = $4, special_pay = $5, updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at`,
		c.Params("id"),
		employee.Name,
		employee.Wage,
		employee.Type,
		employee.SpecialPay,
	).Scan(&employee.ID, &employee.CreatedAt, &employee.UpdatedAt)
	if err != nil {
		log.Print("Error in UpdateEmployee: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while updating employee"})
	}
	return c.Status(200).JSON(employee)
}

func RemoveEmployee(c *fiber.Ctx) error {
	_, err := db.Pool.Exec(
		context.Background(),
		`DELETE FROM employees WHERE id = $1`,
		c.Params("id"),
	)
	if err != nil {
		log.Print("Error in RemoveEmployee: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while removing employee"})
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}
