package models

import (
	"time"
)

type Week struct {
	ID          int              `db:"id" json:"id"`
	StartDate        string           `db:"start_date" json:"startDate"`
	EndDate        string           `db:"end_date" json:"endDate"`
	FirstDate        *string           `db:"first_date" json:"firstDate"`
	LastDate        *string           `db:"last_date" json:"lastDate"`
	StartFloorBudget *float64  `db:"start_floor_budget" json:"startFloorBudget"`
	StartKitchenBudget *float64 `db:"start_kitchen_budget" json:"startKitchenBudget"`
	StartFoodBudget   *float64  `db:"start_food_budget" json:"startFoodBudget"`
	WeeklyPay []Employee `db:"weekly_pay" json:"weeklyPay"`
	ShareLinkToken string `db:"share_link_token" json:"shareLinkToken"`
	CreatedAt   time.Time        `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time        `db:"updated_at" json:"updatedAt"`
}
