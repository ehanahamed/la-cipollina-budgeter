package auth

import (
	"time"
)

type Employee struct {
    ID         	int       `db:"id" json:"id"`
    Username   	string    `db:"name" json:"name"`
    NewPassword string    `json:"newPassword"`
    CreatedAt  	time.Time `db:"created_at" json:"createdAt"`
    UpdatedAt  	time.Time `db:"updated_at" json:"updatedAt"`
}
