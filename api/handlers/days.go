package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"

	"la-cipollina-budgeter-api/db"
	"la-cipollina-budgeter-api/models"
)

func GetDayByDate(c *fiber.Ctx) error {
	var day models.Day
	err := db.Pool.QueryRow(
		context.Background(),
		`SELECT id, date, hours_worked, worked_today,
food_costs, created_at, updated_at
FROM days WHERE date = $1`,
		c.Params("Date"),
	).Scan(
		&day.ID,
		&day.Date,
		&day.HoursWorked,
		&day.WorkedToday,
		&day.FoodCosts,
		&day.CreatedAt,
		&day.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			/* if not found */
			return c.Status(404).JSON(fiber.Map{
				"error": "Day not found",
			})
		}
		/* if other database error */
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
	date, hours_worked, worked_today, food_costs
) VALUES (
	$1, $2, $3, $4
) RETURNING id, created_at, updated_at`,
		day.Date,
		day.HoursWorked,
		day.WorkedToday,
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
		`UPDATE days SET hours_worked = $1,
worked_today = $2, food_costs = $3,
updated_at = now()
WHERE id = $5
RETURNING date, created_at, updated_at`,
		day.HoursWorked,
		day.WorkedToday,
		day.FoodCosts,
		c.Params("id"),
	).Scan(
		&day.Date,
		&day.CreatedAt,
		&day.UpdatedAt,
	)
	if err != nil {
		log.Print("Error in UpdateDay: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while updating day"})
	}
	return c.Status(200).JSON(day)
}
