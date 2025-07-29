package models

import (
	"time"
)

type EmployeeHours struct {
	Employee Employee `json:"employee"`
	Hours    *float64  `json:"hours"`
}

type EmployeeWorked struct {
	Employee    Employee `json:"employee"`
	WorkedToday bool     `json:"workedToday"`
}

type FoodCostsItem struct {
	Cost  *float64 `json:"cost"`
	Notes *string  `json:"notes"`
}

type Day struct {
	ID          int              `db:"id" json:"id"`
	Date        string           `db:"date" json:"date"`
	HoursWorked []EmployeeHours  `db:"hours_worked" json:"hoursWorked"`
	WorkedToday []EmployeeWorked `db:"worked_today" json:"workedToday"`
	FoodCosts   []FoodCostsItem  `db:"food_costs" json:"foodCosts"`
	FoodBudgetIncrease *float64 `db:"food_budget_increase" json:"foodBudgetIncrease"`
	KitchenBudgetIncrease *float64 `db:"kitchen_budget_increase" json:"kitchenBudgetIncrease"`
	FloorBudgetIncrease *float64 `db:"floor_budget_increase" json:"floorBudgetIncrease"`
	CreatedAt   time.Time        `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time        `db:"updated_at" json:"updatedAt"`
}
