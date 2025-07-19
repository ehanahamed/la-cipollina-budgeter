package models

type Valentino struct {
    ID         int     	`db:"id" json:"id"`
    Name       string  	`db:"name" json:"name"`
    WeeklyPay *float64 	`db:"wage" json:"wage"`
    CreatedAt time.Time `db:"created_at" json:"createdAt"`
    UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
