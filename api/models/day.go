package models

import (
	"time"
)

type SpecialPayRule struct {
	PerDay  *float64 `json:"perDay,omitempty"`
	PerHour *float64 `json:"perHour,omitempty"`
}

type SpecialPay map[string]*SpecialPayRule

type Day struct {
	ID         int        `db:"id" json:"id"`
	Day        time.Time  `db:"name" json:"name"`
	Type       string     `db:"type" json:"type"` // "FLOOR" or "KITCHEN"
	Wage       *float64   `db:"wage" json:"wage"` // can be nil if only special pay
	SpecialPay SpecialPay `db:"special_pay" json:"specialPay"`
	CreatedAt  time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updatedAt"`
}
