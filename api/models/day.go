package models

import (
	"time"
)

type EmployeeHours struct {
	Employee Employee `json:"employee"`
	Hours    float64  `json:"hours"`
}

type EmployeeWorked struct {
	Employee    Employee `json:"employee"`
	WorkedToday bool     `json:"workedToday"`
}

type FoodCostsItem struct {
	Cost  float64 `json:"cost"`
	Notes string  `json:"notes"`
}

type Day struct {
	ID               int              `db:"id" json:"id"`
	Date             time.Time        `db:"date" json:"date"`
	HoursWorked      []EmployeeHours  `db:"hours_worked" json:"hoursWorked"`
	WorkedToday      []EmployeeWorked `db:"worked_today" json:"workedToday"`
	CurrentEmployees []Employee       `db:"current_employees" json:"currentEmployees"`
	FoodCosts        []FoodCostsItem  `db:"food_costs" json:"foodCosts"`
	CreatedAt        time.Time        `db:"created_at" json:"createdAt"`
	UpdatedAt        time.Time        `db:"updated_at" json:"updatedAt"`
}
