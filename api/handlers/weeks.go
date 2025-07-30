package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/georgysavva/scany/v2/pgxscan"

	"la-cipollina-budgeter-api/db"
	"la-cipollina-budgeter-api/models"
)

func GetWeekByDate(c *fiber.Ctx) error {
	var week models.Week
	err := db.Pool.QueryRow(
		context.Background(),
		`SELECT id, start_date::text, end_date::text,
	start_floor_budget, start_kitchen_budget,
	start_food_budget, created_at, updated_at
FROM weeks
WHERE $1::date BETWEEN start_date AND end_date
LIMIT 1`,
		c.Params("Date"),
	).Scan(
		&week.ID,
		&week.StartDate,
		&week.EndDate,
		&week.StartFloorBudget,
		&week.StartKitchenBudget,
		&week.StartFoodBudget,
		&week.CreatedAt,
		&week.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			/* if not found */
			return c.Status(404).JSON(fiber.Map{
				"error": "Week not found",
				"statusCode": 404,
			})
		}
		/* if other database error */
		log.Print("Error in GetWeekByDate: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while getting week by date"})
	}
	return c.JSON(week)
}

func GetAllWeeks(c *fiber.Ctx) error {
	var weeks []models.Week
	err := pgxscan.Select(
		context.Background(),
		db.Pool,
		&weeks,
		`SELECT id, start_date::text, end_date::text,
start_floor_budget, start_kitchen_budget,
start_food_budget, created_at, updated_at
FROM weeks`,
	)
	if err != nil {
		log.Print("Error in GetWeekByDate: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while getting week by date"})
	}
	return c.JSON(weeks)
}

func AddWeek(c *fiber.Ctx) error {
	var week models.Week
	if err := c.BodyParser(&week); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := db.Pool.QueryRow(
		context.Background(),
`INSERT INTO weeks (
	start_date, end_date, start_floor_budget,
	start_kitchen_budget, start_food_budget
) VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at`,
		week.StartDate,
		week.EndDate,
		week.StartFloorBudget,
		week.StartKitchenBudget,
		week.StartFoodBudget,
	).Scan(
		&week.ID,
		&week.CreatedAt,
		&week.UpdatedAt,
	)
	if err != nil {
		log.Print("Error in AddWeek: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while adding week"})
	}
	return c.Status(201).JSON(week)
}

func DeleteWeek(c *fiber.Ctx) error {
	_, err := db.Pool.Exec(
		context.Background(),
		`DELETE FROM weeks WHERE id = $1`,
		c.Params("id"),
	)
	if err != nil {
		log.Print("Error in DeleteWeek: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while deleting week"})
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}
