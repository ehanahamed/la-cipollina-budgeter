package models

import (
	"time"
)

type Day struct {
	ID               int             `db:"id" json:"id"`
	Date             time.Time       `db:"date" json:"date"`
	HoursWorked      map[int]float64 `db:"hours_worked" json:"hoursWorked"`
	WorkedToday      map[int]bool    `db:"worked_today" json:"workedToday"`
	CurrentEmployees []Employee      `db:"current_employees" json:"currentEmployees"`
	FoodCosts        []float64       `db:"food_costs" json:foodCosts`
	CreatedAt        time.Time       `db:"created_at" json:"createdAt"`
	UpdatedAt        time.Time       `db:"updated_at" json:"updatedAt"`
}
