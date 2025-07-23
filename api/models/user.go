package models

import (
	"time"
)

type User struct {
	ID        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

type UserInput struct {
	Username    *string `db:"username" json:"username"`
	NewPassword *string `json:"newPassword"`
}
