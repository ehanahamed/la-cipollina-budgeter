package models

import (
	"time"
)

type Valentino struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	WeeklyPay *float64  `db:"weekly_pay" json:"weeklyPay"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
