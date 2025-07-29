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

func GetDayByDate(c *fiber.Ctx) error {
	var day models.Day
	err := db.Pool.QueryRow(
		context.Background(),
		`SELECT id, date::text, hours_worked,
worked_today, food_costs, food_budget_increase,
kitchen_budget_increase, floor_budget_increase,
created_at, updated_at
FROM days WHERE date = $1`,
		c.Params("date"),
	).Scan(
		&day.ID,
		&day.Date,
		&day.HoursWorked,
		&day.WorkedToday,
		&day.FoodCosts,
		&day.FoodBudgetIncrease,
		&day.KitchenBudgetIncrease,
		&day.FloorBudgetIncrease,
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

func GetDaysInWeekByDate(c *fiber.Ctx) error {
	var days []models.Day
	err := pgxscan.Select(
		context.Background(),
		db.Pool,
		&days,
`SELECT d.id, d.date::text,
	d.hours_worked, d.worked_today, d.food_costs,
	d.food_budget_increase, d.kitchen_budget_increase,
	d.floor_budget_increase, d.created_at, d.updated_at
FROM days d
JOIN weeks w
	ON d.date BETWEEN w.start_date AND w.end_date
WHERE $1::date BETWEEN w.start_date and w.end_date
ORDER BY d.date`,
		c.Params("date"),
	)
	if err != nil {
		log.Print("Error in GetDaysInWeekByDate: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while getting days from week by date"})
	}
	return c.JSON(days)
}

func RecordDay(c *fiber.Ctx) error {
	var day models.Day
	if err := c.BodyParser(&day); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := db.Pool.QueryRow(
		context.Background(),
		`INSERT INTO days (
	date, hours_worked, worked_today, food_costs,
	food_budget_increase, kitchen_budget_increase,
	floor_budget_increase
) VALUES (
	$1, $2, $3, $4, $5, $6, $7
) RETURNING id, created_at, updated_at`,
		day.Date,
		day.HoursWorked,
		day.WorkedToday,
		day.FoodCosts,
		day.FoodBudgetIncrease,
		day.KitchenBudgetIncrease,
		day.FloorBudgetIncrease,
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
		`UPDATE days SET hours_worked = $2,
worked_today = $3, food_costs = $4,
food_budget_increase = $5,
kitchen_budget_increase = $6,
floor_budget_increase = $7
updated_at = now()
WHERE id = $1
RETURNING date::text, created_at, updated_at`,
		c.Params("id"),
		day.HoursWorked,
		day.WorkedToday,
		day.FoodCosts,
		day.FoodBudgetIncrease,
		day.KitchenBudgetIncrease,
		day.FloorBudgetIncrease,
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

func DeleteDay(c *fiber.Ctx) error {
	_, err := db.Pool.Exec(
		context.Background(),
		`DELETE FROM days WHERE id = $1`,
		c.Params("id"),
	)
	if err != nil {
		log.Print("Error in DeleteDay: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Database error while deleting day"})
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}
