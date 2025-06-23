package models

import (
	"time"
)

type User struct {
    ID         int       `db:"id" json:"id"`
    Username   string    `db:"name" json:"name"`
    CreatedAt  time.Time `db:"created_at" json:"createdAt"`
    UpdatedAt  time.Time `db:"updated_at" json:"updatedAt"`
}

type NewUser struct {
    Username   	string `db:"name" json:"name"`
	NewPassword string `json:"newPassword"`
}
